# 🖼️ Pexels爬虫

> Pexels免费视频素材爬虫，专门解决浏览器指纹识别问题，使用req/v3库模拟Safari指纹绕过检测

## 📋 功能特性

- ✅ **Safari指纹模拟**：使用req/v3的ImpersonateSafari功能绕过指纹检测
- 🎬 **视频搜索**：搜索高质量免费视频素材
- 🌐 **代理支持**：支持HTTP/HTTPS代理配置
- 📊 **完整数据**：返回完整的视频信息、作者信息、标签等

## 🚀 快速开始

### 安装依赖

```bash
go get github.com/imroc/req/v3
```

### 基本使用

#### 搜索视频

```go
package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "time"

    "github.com/xieburoucoco/spider-hub/platforms/pexels"
)

func main() {
    // 创建爬虫实例
    spider := pexels.NewPexelsSpider()

    // 设置超时上下文
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    fmt.Println("🎬 Pexels视频搜索功能演示")
    fmt.Println(strings.Repeat("=", 80))

    // 搜索关键词
    query := "tree"
    fmt.Printf("🔍 搜索关键词: %s\n\n", query)

    // 执行搜索
    result, err := spider.SearchVideos(ctx, query, 1)
    if err != nil {
        log.Fatalf("❌ 搜索视频失败: %v", err)
    }

    videos := result.PageProps.InitialData.Data
    fmt.Printf("✅ 找到 %d 个视频:\n\n", len(videos))

    // 显示前5个结果
    for i, video := range videos {
        if i >= 5 {
            break
        }

        attrs := video.Attributes
        fmt.Printf("视频 %d:\n", i+1)
        fmt.Printf("  ID: %s\n", video.ID)
        fmt.Printf("  标题: %s\n", attrs.Title)
        fmt.Printf("  描述: %s\n", attrs.Description)
        fmt.Printf("  尺寸: %dx%d\n", attrs.Width, attrs.Height)
        if attrs.Duration != nil {
            fmt.Printf("  时长: %d秒\n", *attrs.Duration)
        } else {
            fmt.Printf("  时长: 未知\n")
        }
        fmt.Printf("  作者: %s %s (@%s)\n", attrs.User.FirstName, attrs.User.LastName, attrs.User.Username)

        if len(attrs.VideoFiles) > 0 {
            fmt.Printf("  视频文件:\n")
            for j, file := range attrs.VideoFiles {
                if j >= 3 { // 只显示前3个文件
                    break
                }
                fmt.Printf("    - %s (%dx%d): %s\n", file.Quality, file.Width, file.Height, file.Link)
            }
        }
        fmt.Println()
    }

    fmt.Println(strings.Repeat("=", 80))
    fmt.Println("🎉 视频搜索测试完成！")
}
```

### 高级用法

#### 使用代理

```go
// 设置代理
spider := pexels.NewPexelsSpider()
err := spider.SetProxy("http://127.0.0.1:8080")
if err != nil {
    log.Fatalf("设置代理失败: %v", err)
}

// 使用代理进行搜索
result, err := spider.SearchVideos(ctx, "ocean", 1)
```

## 🔧 浏览器指纹绕过技术

### 核心技术

本爬虫使用 `github.com/imroc/req/v3` 库的 `ImpersonateSafari()` 功能来模拟真实的Safari浏览器指纹：

```go
// 创建客户端并启用开发模式
client := req.DevMode()

// 模拟Safari浏览器指纹
client.ImpersonateSafari()
```

### 指纹特征

模拟的Safari指纹包括：

- **TLS指纹**：模拟Safari的TLS握手特征
- **HTTP/2设置**：匹配Safari的HTTP/2参数
- **请求头顺序**：按照Safari的请求头顺序发送
- **User-Agent**：使用真实的Safari User-Agent
- **Accept头**：匹配Safari的Accept头格式

### 指纹校验

通过运行测试用例可以验证指纹是否成功绕过检测：

```bash
# 运行视频搜索测试，验证指纹绕过效果
go test -v -run TestSearchVideos
```

如果能成功获取到视频数据并显示详细信息，说明Safari指纹模拟成功绕过了Pexels的检测机制。

## 📊 返回数据结构

### SearchResponse - 搜索响应

```go
type SearchResponse struct {
    PageProps struct {
        InitialData struct {
            Data []MediaItem `json:"data"`
        } `json:"initialData"`
    } `json:"pageProps"`
}
```

### MediaItem - 视频项目

```go
type MediaItem struct {
    ID         string `json:"id"`
    Type       string `json:"type"` // "video"
    Attributes struct {
        ID          int         `json:"id"`
        Title       string      `json:"title"`
        Description string      `json:"description"`
        Width       int         `json:"width"`
        Height      int         `json:"height"`
        Duration    *int        `json:"duration,omitempty"`    // 视频时长
        VideoFiles  []VideoFile `json:"video_files,omitempty"` // 视频文件
        User        User        `json:"user"`
        Tags        []Tag       `json:"tags"`
    } `json:"attributes"`
}
```

### VideoFile - 视频文件

```go
type VideoFile struct {
    Quality  string `json:"quality"`  // "hd", "sd", "uhd"
    FileType string `json:"file_type"` // "video/mp4"
    Width    int    `json:"width"`
    Height   int    `json:"height"`
    FPS      float64 `json:"fps"`      // 帧率
    Link     string `json:"link"`     // 下载链接
}
```

## 🧪 测试

运行测试用例验证指纹绕过效果：

```bash
# 进入pexels目录
cd platforms/pexels

# 运行视频搜索测试
go test -v -run TestSearchVideos
```

### 测试输出示例

```
=== RUN   TestSearchVideos
🎬 Pexels视频搜索功能演示
================================================================================
🔍 搜索关键词: tree

✅ 找到 24 个视频:

视频 1:
  ID: 1448735
  标题: Tree Branches Swaying
  描述: Beautiful tree branches moving in the wind
  尺寸: 1920x1080
  时长: 15秒
  作者: John Doe (@johndoe)
  视频文件:
    - uhd (3840x2160): https://player.vimeo.com/external/...
    - hd (1920x1080): https://player.vimeo.com/external/...
    - sd (1280x720): https://player.vimeo.com/external/...

...

================================================================================
🎉 视频搜索测试完成！
--- PASS: TestSearchVideos (3.21s)
```

如果测试通过并显示视频数据，说明Safari指纹模拟成功绕过了Pexels的检测机制。

## ⚠️ 注意事项

### 使用限制
- 请遵守Pexels的使用条款和API限制
- 建议在请求间添加适当延时，避免被限制
- 下载的媒体文件请遵守相应的许可协议

### 指纹检测
- Pexels使用先进的浏览器指纹识别技术
- 本爬虫通过模拟Safari指纹来绕过检测
- 如果检测机制更新，可能需要调整指纹参数

### 技术说明
- 使用req/v3库的ImpersonateSafari功能
- 自动处理TLS指纹、HTTP/2设置等
- 支持完整的请求调试和分析

### 免责声明
- 本工具仅供学习和研究使用
- 请遵守目标网站的使用条款
- 不得用于商业用途或恶意爬取
- 使用者需自行承担使用风险

## 🔗 相关链接

- [Pexels官网](https://www.pexels.com/)
- [req/v3库文档](https://github.com/imroc/req)
- [浏览器指纹识别技术](https://github.com/imroc/req/blob/master/docs/tutorial/impersonate.md)

## 📄 许可证

MIT License