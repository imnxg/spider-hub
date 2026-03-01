### 接入其他 Go 项目（SDK 集成）

如果你想在 **其他独立的 Go 业务项目** 中直接引入它的能力，只需拉取依赖并调用：

1. **安装依赖**
```bash
go get github.com/xieburoucoco/spider-hub
```
*(如果 `spider-hub` 尚未推送到远端，可在宿主项目的 `go.mod` 中使用 `replace` 指向本地目录进行调试)*

2. **在宿主项目中编写调用代码**
```go
package main

import (
	"context"
	"fmt"
	"log"
	
	// 直接引入 amazon 爬虫包
	"github.com/xieburoucoco/spider-hub/platforms/amazon"
)

func main() {
	// 创建爬虫实例
	spider := amazon.NewAmazonSpider()
	
	// 配置要爬取的链接
	productURL := "https://www.amazon.com/dp/B08N5WRWNW"
	
	// 执行抓取（支持传入外部 ctx 用于控制超时和链路追踪）
	result, err := spider.FetchProductDetail(context.Background(), productURL)
	if err != nil {
		log.Fatalf("抓取失败: %v", err)
	}
	
	fmt.Printf("抓取成功！获取到商品: %s\n", result.Title)
}
```

以图片搜索商品调用接口示例：
```go
package main

import (
	"context"
	"fmt"
	"log"
	
	// 直接引入 amazon 爬虫包
	"github.com/xieburoucoco/spider-hub/platforms/amazon"
)

func main() {
	// 创建亚马逊爬虫实例
	spider := NewAmazonSpider()

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Println("🖼️ 亚马逊图片搜索功能演示")
	fmt.Println(strings.Repeat("=", 80))

	// 示例1: 通过在线图片URL搜索商品
	fmt.Println("\n📡 示例1: 通过在线图片URL搜索商品")
	fmt.Println(strings.Repeat("-", 50))

	imageURL := "https://m.media-amazon.com/images/I/61RjUcDPH1L._AC_SX679_.jpg"
	fmt.Printf("🔍 搜索图片: %s\n\n", imageURL)

	products, err := spider.SearchProductsByImageURL(ctx, imageURL, nil)
	if err != nil {
		t.Logf("❌ 图片搜索失败: %v", err)
		return
	}

	fmt.Printf("✅ 找到 %d 个相关商品:\n\n", len(products))
	displayProducts(products, 5) // 显示前5个结果

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("🎉 URL图片搜索测试完成！")
	
	
}
```

以图片搜索商品调用接口示例：
```go
package main

import (
	"context"
	"fmt"
	"log"
	
	// 直接引入 amazon 爬虫包
	"github.com/xieburoucoco/spider-hub/platforms/amazon"
)

func main() {
	// 创建亚马逊爬虫实例
	spider := NewAmazonSpider()

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Println("🖼️ 亚马逊图片搜索功能演示")
	fmt.Println(strings.Repeat("=", 80))

	// 示例2: 通过本地图片文件搜索商品
	fmt.Println("\n📁 示例2: 通过本地图片文件搜索商品")
	fmt.Println(strings.Repeat("-", 50))

	// 首先下载一个测试图片到本地
	imageURL := "https://m.media-amazon.com/images/I/71c-jiE2IcL._AC_SX679_.jpg"
	testImagePath := "test_camera.jpg"
	fmt.Printf("⬇️ 下载测试图片到: %s\n", testImagePath)

	if err := downloadImageToFile(imageURL, testImagePath); err != nil {
		t.Fatalf("❌ 下载测试图片失败: %v", err)
	}
	defer os.Remove(testImagePath) // 清理临时文件

	// 读取本地图片文件
	imageData, err := readLocalImageFile(testImagePath)
	if err != nil {
		t.Fatalf("❌ 读取本地图片失败: %v", err)
	}

	fmt.Printf("📄 图片文件大小: %.2f KB\n", float64(len(imageData))/1024)
	fmt.Println("🔍 开始搜索...")

	// 搜索商品
	products, err := spider.SearchProductsByImageData(ctx, imageData, nil)
	if err != nil {
		t.Logf("❌ 图片搜索失败: %v", err)
		return
	}

	fmt.Printf("✅ 找到 %d 个相关商品:\n\n", len(products))
	displayProducts(products, 5) // 显示前5个结果

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("🎉 本地图片搜索测试完成！")
	fmt.Println("\n💡 使用提示:")
	fmt.Println("  • 建议在请求之间添加适当延时，避免被反爬虫机制阻止")
	fmt.Println("  • 使用高质量、清晰的图片能获得更好的搜索结果")
	fmt.Println("  • 如果访问受限，可以配置代理服务器")
	fmt.Println("  • 搜索结果可能随时间变化，建议定期更新数据")
	
	
}
```