package foresightnews

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// ForesightNewsSpider Foresight News爬虫结构体
type ForesightNewsSpider struct {
	client   *http.Client
	proxyURL string
}

// NewForesightNewsSpider 创建新的Foresight News爬虫实例
func NewForesightNewsSpider() *ForesightNewsSpider {
	return &ForesightNewsSpider{
		client: &http.Client{},
	}
}

// SetProxy 设置代理
func (s *ForesightNewsSpider) SetProxy(proxyURL string) {
	s.proxyURL = proxyURL
	// 这里可以根据需要配置代理客户端
}

// DecryptBase64AndZlib 解密Base64和Zlib压缩的数据
func (s *ForesightNewsSpider) DecryptBase64AndZlib(ctx context.Context, base64String string) ([]byte, error) {
	// 1. Base64 解码
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, fmt.Errorf("❌ Base64解码失败: %w", err)
	}

	// 2. 创建 zlib 解压读取器
	buf := bytes.NewReader(decodedBytes)
	zlibReader, err := zlib.NewReader(buf)
	if err != nil {
		return nil, fmt.Errorf("❌ 创建zlib读取器失败: %w", err)
	}
	defer zlibReader.Close()

	// 3. 解压缩数据
	result, err := io.ReadAll(zlibReader)
	if err != nil {
		return nil, fmt.Errorf("❌ Zlib解压缩失败: %w", err)
	}

	return result, nil
}

// FeedV2 获取资讯列表数据，直接返回解密后的字节数据
func (s *ForesightNewsSpider) FeedV2(ctx context.Context, page, size int64, proxyURL string) ([]byte, error) {
	// 参数验证
	if page < 1 {
		page = 1
	}
	if size < 1 || size > MaxPageSize {
		size = DefaultPageSize
	}

	// 构建请求URL和参数
	url := BaseURL + FeedV2Endpoint + "?page=" + strconv.FormatInt(page, 10) + "&size=" + strconv.FormatInt(size, 10)

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("❌ 创建请求失败: %w", err)
	}

	// 设置请求头
	for key, value := range HTTPHeaders {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("❌ 请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("❌ 读取响应体失败: %w", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("❌ 请求失败，状态码: %d，响应: %s", resp.StatusCode, string(responseBody))
	}

	// 响应获取成功

	// 解析响应结构
	var apiResponse FeedV2Res
	if err = json.Unmarshal(responseBody, &apiResponse); err != nil {
		return nil, fmt.Errorf("❌ 解析响应JSON失败: %w，响应内容: %s", err, string(responseBody))
	}

	// 检查响应码 - 看起来code为1是正常的，不是错误
	// if apiResponse.Code != 0 {
	//     return nil, fmt.Errorf("❌ API返回错误，代码: %d，消息: %s，完整响应: %s", apiResponse.Code, apiResponse.Message, string(responseBody))
	// }

	// 检查数据是否为空
	if apiResponse.Data == "" {
		return nil, fmt.Errorf("❌ API返回的数据为空，完整响应: %s", string(responseBody))
	}

	// 解密数据
	decryptedBytes, err := s.DecryptBase64AndZlib(ctx, apiResponse.Data)
	if err != nil {
		return nil, fmt.Errorf("❌ 解密数据失败: %w", err)
	}

	return decryptedBytes, nil
}

// FeedV2WithStruct 获取资讯列表数据，返回结构化数据（可选方法）
func (s *ForesightNewsSpider) FeedV2WithStruct(ctx context.Context, page, size int64, proxyURL string) (*FeedV2Data, error) {
	// 获取原始字节数据
	decryptedBytes, err := s.FeedV2(ctx, page, size, proxyURL)
	if err != nil {
		return nil, err
	}

	// 解析为结构化数据
	var feedData FeedV2Data
	if err = json.Unmarshal(decryptedBytes, &feedData); err != nil {
		return nil, fmt.Errorf("❌ 解析解密数据失败: %w", err)
	}

	return &feedData, nil
}
