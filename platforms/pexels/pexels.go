package pexels

import (
    "context"
    "encoding/json"
    "fmt"
    "net/url"
    "time"

    "github.com/imroc/req/v3"
)

// PexelsSpider Pexels爬虫客户端
type PexelsSpider struct {
    client *req.Client
}

// NewPexelsSpider 创建新的Pexels爬虫实例
func NewPexelsSpider() *PexelsSpider {
    // 创建客户端并启用开发模式
    client := req.DevMode()
    
    // 模拟Safari浏览器指纹
    client.ImpersonateSafari()
    
    // 设置通用请求头
    client.SetCommonHeaders(map[string]string{
        "Accept":          "application/json, text/plain, */*",
        "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
        "Cache-Control":   "no-cache",
        "Pragma":          "no-cache",
        "Sec-Fetch-Dest":  "empty",
        "Sec-Fetch-Mode":  "cors",
        "Sec-Fetch-Site":  "same-origin",
    })
    
    // 设置超时
    client.SetTimeout(30 * time.Second)
    
    return &PexelsSpider{
        client: client,
    }
}

// SearchVideos 搜索视频
func (s *PexelsSpider) SearchVideos(ctx context.Context, query string, page int) (*SearchResponse, error) {
    if page < 1 {
        page = 1
    }
    
    // 构建API URL
    apiURL := fmt.Sprintf("%s/zh-CN/search/videos/%s.json", NextDataBaseURL, url.QueryEscape(query))
    
    // 发送请求
    resp, err := s.client.R().
        SetContext(ctx).
        SetQueryParam("query", query).
        Get(apiURL)
    
    if err != nil {
        return nil, fmt.Errorf("请求失败: %w", err)
    }
    
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
    }
    
    // 解析响应
    var result SearchResponse
    if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
        return nil, fmt.Errorf("解析响应失败: %w", err)
    }
    
    return &result, nil
}

// SearchPhotos 搜索图片
func (s *PexelsSpider) SearchPhotos(ctx context.Context, query string, page int) (*SearchResponse, error) {
    if page < 1 {
        page = 1
    }
    
    // 构建API URL
    apiURL := fmt.Sprintf("%s/zh-CN/search/%s.json", NextDataBaseURL, url.QueryEscape(query))
    
    // 发送请求
    resp, err := s.client.R().
        SetContext(ctx).
        SetQueryParam("query", query).
        Get(apiURL)
    
    if err != nil {
        return nil, fmt.Errorf("请求失败: %w", err)
    }
    
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
    }
    
    // 解析响应
    var result SearchResponse
    if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
        return nil, fmt.Errorf("解析响应失败: %w", err)
    }
    
    return &result, nil
}

// GetVideoDetail 获取视频详情
func (s *PexelsSpider) GetVideoDetail(ctx context.Context, videoID int) (*VideoDetail, error) {
    // 构建API URL
    apiURL := fmt.Sprintf("%s/zh-CN/video/%d.json", NextDataBaseURL, videoID)
    
    // 发送请求
    resp, err := s.client.R().
        SetContext(ctx).
        Get(apiURL)
    
    if err != nil {
        return nil, fmt.Errorf("请求失败: %w", err)
    }
    
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
    }
    
    // 解析响应
    var result VideoDetailResponse
    if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
        return nil, fmt.Errorf("解析响应失败: %w", err)
    }
    
    if result.PageProps.Medium.Type != "video" {
        return nil, fmt.Errorf("返回的不是视频类型")
    }

    return &result.PageProps.Medium.Attributes, nil
}

// SetProxy 设置代理
func (s *PexelsSpider) SetProxy(proxyURL string) error {
    s.client.SetProxyURL(proxyURL)
    return nil
}

// DumpRequest 调试请求信息
func (s *PexelsSpider) DumpRequest(ctx context.Context, url string) (string, error) {
    resp, err := s.client.R().
        SetContext(ctx).
        Get(url)
    
    if err != nil {
        return "", fmt.Errorf("请求失败: %w", err)
    }
    
    return resp.Dump(), nil
}