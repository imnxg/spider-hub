package pexels

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

// TestSearchVideos 测试搜索视频功能
func TestSearchVideos(t *testing.T) {
	spider := NewPexelsSpider()

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
		t.Fatalf("❌ 搜索视频失败: %v", err)
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
