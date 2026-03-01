package amazon

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

// TestAmazonProductDetail 测试获取亚马逊商品详情
func TestAmazonProductDetail(t *testing.T) {
	// 创建亚马逊爬虫实例
	spider := NewAmazonSpider()

	// 设置测试URL (Sony WH-1000XM5 耳机)
	//testURL := "https://www.amazon.com/Sony-WH-1000XM5-Canceling-Headphones-Hands-Free/dp/B09XS7JWHH"
	testURL := "https://www.amazon.com/dp/B0G2LT19TC/ref=mweb_up_am_fl_st_na_un_up_sm_web?th=1&psc=1"

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 获取商品详情
	result, err := spider.FetchProductDetail(ctx, testURL)
	if err != nil {
		t.Fatalf("❌ 获取商品详情失败: %v", err)
	}

	// 验证结果
	if result.Title == "" {
		t.Error("❌ 商品标题为空")
	} else {
		t.Logf("✅ 商品标题: %s", result.Title)
	}

	if len(result.Images) == 0 {
		t.Error("❌ 没有找到商品图片")
	} else {
		t.Logf("✅ 找到 %d 张商品图片", len(result.Images))
		t.Logf("  第一张图片: %s", result.Images[0])
	}

	if result.Price == nil {
		t.Log("⚠️ 没有找到商品价格")
	} else {
		t.Logf("✅ 商品价格: %s", *result.Price)
	}

	t.Logf("✅ 商品描述: %s", result.Desc)
	t.Logf("✅ 商品语言: %s", result.Language)

	// 输出完整结果
	priceStr := "无"
	if result.Price != nil {
		priceStr = *result.Price
	}
	discountStr := "无"
	if result.Discount != nil {
		discountStr = *result.Discount + "%"
	}

	t.Logf("\n🛍️ 商品详情结果:\n"+
		"  标题: %s\n"+
		"  描述: %s\n"+
		"  价格: %s\n"+
		"  折扣: %s\n"+
		"  图片数量: %d\n"+
		"  视频数量: %d\n",
		result.Title,
		truncateString(result.Desc, 100),
		priceStr,
		discountStr,
		len(result.Images),
		len(result.Videos))
}

// TestSearchProductsByImageURL 测试通过在线图片URL搜索商品
func TestSearchProductsByImageURL(t *testing.T) {
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

	imageURL := "https://m.media-amazon.com/images/I/71c-jiE2IcL._AC_SX679_.jpg"
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

// TestSearchProductsByImageData 测试通过本地图片文件搜索商品
func TestSearchProductsByImageData(t *testing.T) {
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

// displayProducts 格式化显示商品信息
func displayProducts(products []ImageSearchProduct, maxCount int) {
	count := len(products)
	if maxCount > 0 && maxCount < count {
		count = maxCount
	}

	for i := 0; i < count; i++ {
		product := products[i]
		fmt.Printf("📦 商品 %d:\n", i+1)
		fmt.Printf("   🏷️  标题: %s\n", truncateString(product.Title, 60))
		fmt.Printf("   🏭 品牌: %s\n", product.ByLine)
		fmt.Printf("   💰 价格: %s", product.Price)
		if product.ListPrice != nil && *product.ListPrice != "" {
			fmt.Printf(" (原价: %s)", *product.ListPrice)
		}
		fmt.Println()
		fmt.Printf("   ⭐ 评分: %.1f (%s条评论)\n", product.AverageOverallRating, product.TotalReviewCount)
		fmt.Printf("   📦 库存: %s\n", getAvailabilityText(product.Availability))
		fmt.Printf("   🔗 链接: %s\n", product.LinkURL)

		if len(product.TwisterVariations) > 0 {
			fmt.Printf("   🎨 变体: %d个选项\n", len(product.TwisterVariations))
		}
		fmt.Println()
	}

	if len(products) > count {
		fmt.Printf("... 还有 %d 个商品未显示\n", len(products)-count)
	}
}

// truncateString 截断字符串
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// getAvailabilityText 获取库存状态的中文描述
func getAvailabilityText(availability string) string {
	switch availability {
	case "IN_STOCK":
		return "有库存 ✅"
	case "IN_STOCK_SCARCE":
		return "库存紧张 ⚠️"
	case "OUT_OF_STOCK":
		return "缺货 ❌"
	default:
		return availability
	}
}

// downloadImageToFile 下载图片到本地文件
func downloadImageToFile(imageURL, filePath string) error {
	// 直接下载图片
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("下载失败，状态码: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// readLocalImageFile 读取本地图片文件
func readLocalImageFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
