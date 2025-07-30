# 🔮 Foresight News爬虫

> Foresight News是一个专业的区块链资讯平台，提供最新的加密货币、DeFi、NFT等领域的新闻和深度分析文章。本爬虫实现了对其资讯API的数据获取和解密功能。

## 📋 功能特性

- ✅ 获取最新区块链资讯列表
- ✅ 支持分页查询
- ✅ Base64和Zlib数据解密
- ✅ 返回原始字节数据或结构化数据
- ✅ 支持代理配置
- ✅ 完善的错误处理机制
- ✅ 上下文超时控制

## 🚀 快速开始

### 安装依赖

```bash
go mod init your-project
go get -u github.com/your-username/spider-hub/platforms/foresightnews
```

### 基本使用

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/your-username/spider-hub/platforms/foresightnews"
)

func main() {
    // 创建爬虫实例
    spider := foresightnews.NewForesightNewsSpider()
    
    // 设置超时上下文
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // 获取第1页，20条资讯（返回字节数据）
    data, err := spider.FeedV2(ctx, 1, 20, "")
    if err != nil {
        log.Fatalf("获取资讯失败: %v", err)
    }
    
    fmt.Printf("获取到 %d 字节的数据\n", len(data))
    fmt.Printf("原始JSON数据: %s\n", string(data))
}
```

### 高级用法

#### 1. 获取结构化数据

```go
// 获取结构化数据
feedData, err := spider.FeedV2WithStruct(ctx, 1, 10, "")
if err != nil {
    log.Fatalf("获取资讯失败: %v", err)
}

fmt.Printf("获取到 %d 条资讯\n", len(feedData.List))

// 遍历资讯列表
for i, item := range feedData.List {
    fmt.Printf("资讯 %d:\n", i+1)
    fmt.Printf("  ID: %d\n", item.ID)
    fmt.Printf("  类型: %s\n", item.SourceType)
    
    if item.News != nil {
        fmt.Printf("  新闻标题: %s\n", item.News.Title)
        fmt.Printf("  新闻摘要: %s\n", item.News.Brief)
        fmt.Printf("  重要性: %v\n", item.News.IsImportant)
    } else if item.Article != nil {
        fmt.Printf("  文章标题: %s\n", item.Article.Title)
        fmt.Printf("  文章摘要: %s\n", item.Article.Brief)
        if item.Article.Author != nil {
            fmt.Printf("  作者: %s\n", item.Article.Author.Username)
        }
    }
    fmt.Println()
}
```

#### 2. 使用代理

```go
// 设置代理
proxyURL := "http://127.0.0.1:8080"
data, err := spider.FeedV2(ctx, 1, 20, proxyURL)
if err != nil {
    log.Fatalf("使用代理获取资讯失败: %v", err)
}
```

#### 3. 手动解密数据

```go
// 如果你有Base64+Zlib加密的数据，可以手动解密
encryptedData := "your-base64-zlib-data"
decryptedBytes, err := spider.DecryptBase64AndZlib(ctx, encryptedData)
if err != nil {
    log.Fatalf("解密失败: %v", err)
}

fmt.Printf("解密结果: %s\n", string(decryptedBytes))
```

## 📊 返回数据结构

### 主要数据结构

#### FeedV2Data
```go
type FeedV2Data struct {
    List []FeedItem `json:"list"`
}
```

#### FeedItem
```go
type FeedItem struct {
    ID          int64        `json:"id"`
    SourceID    int64        `json:"source_id"`
    SourceType  string       `json:"source_type"`  // "news" 或 "article"
    PublishedAt int64        `json:"published_at"` // Unix时间戳
    News        *NewsItem    `json:"news,omitempty"`
    Article     *ArticleItem `json:"article,omitempty"`
}
```

#### NewsItem（新闻）
```go
type NewsItem struct {
    ID           int64  `json:"id"`
    Title        string `json:"title"`        // 新闻标题
    Brief        string `json:"brief"`        // 新闻摘要
    Content      string `json:"content"`      // 新闻内容（HTML格式）
    Img          string `json:"img"`          // 新闻图片URL
    Tags         []Tag  `json:"tags"`         // 标签列表
    IsImportant  bool   `json:"is_important"` // 是否重要新闻
    ImportantTag *Tag   `json:"important_tag,omitempty"` // 重要标签
    Label        string `json:"label"`        // 标签文本
    SourceLink   string `json:"source_link"`  // 原文链接
    PublishedAt  int64  `json:"published_at"` // 发布时间
    Favorited    bool   `json:"favorited"`    // 是否收藏
}
```

#### ArticleItem（文章）
```go
type ArticleItem struct {
    ID          int64   `json:"id"`
    Title       string  `json:"title"`        // 文章标题
    Brief       string  `json:"brief"`        // 文章摘要
    Content     string  `json:"content"`      // 文章内容
    Img         string  `json:"img"`          // 文章封面图
    SourceLink  string  `json:"source_link"`  // 原文链接
    Column      *Column `json:"column,omitempty"` // 专栏信息
    Author      *Author `json:"author,omitempty"` // 作者信息
    IsTop       bool    `json:"is_top"`       // 是否置顶
    IsImportant bool    `json:"is_important"` // 是否重要
    PublishedAt int64   `json:"published_at"` // 发布时间
    Status      string  `json:"status"`       // 状态
}
```

## 🧪 测试和示例

```bash
# 进入foresightnews目录
cd platforms/foresightnews

# 运行所有测试
go test -v

# 运行特定测试
go test -v -run TestFeedV2
go test -v -run TestDecryptBase64AndZlib
```

### 测试用例说明

- `TestFeedV2`: 测试获取资讯列表（字节数据）
- `TestFeedV2WithStruct`: 测试获取资讯列表（结构化数据）
- `TestDecryptBase64AndZlib`: 测试Base64和Zlib解密功能
- `TestInvalidParameters`: 测试无效参数处理

## ⚠️ 注意事项

### 使用限制
- 请合理控制请求频率，避免对服务器造成压力
- 建议在请求间添加适当的延时
- 单次请求的最大分页大小为100条

### 数据说明
- API返回的数据经过Base64编码和Zlib压缩
- 时间戳为Unix格式，需要转换为可读时间
- 内容字段可能包含HTML标签
- 部分字段可能为空，使用时需要检查

### 错误处理
- 网络请求失败会返回相应错误
- Base64解码失败会返回解码错误
- Zlib解压失败会返回解压错误
- API返回错误码时会返回相应错误信息

### 免责声明
- 本工具仅供学习和研究使用
- 请遵守目标网站的使用条款和robots.txt协议
- 不得用于商业用途或恶意爬取
- 使用者需自行承担使用风险

## 📝 更新日志

### v1.0.0 (2024-07-30)
- ✅ 初始版本发布
- ✅ 实现FeedV2接口数据获取
- ✅ 实现Base64和Zlib数据解密
- ✅ 支持字节数据和结构化数据返回
- ✅ 完善的测试用例和文档
