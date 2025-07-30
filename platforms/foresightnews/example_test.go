package foresightnews

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// TestFeedV2 测试获取资讯列表（返回字节数据）
func TestFeedV2(t *testing.T) {
	spider := NewForesightNewsSpider()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 测试获取第一页，20条数据
	data, err := spider.FeedV2(ctx, 1, 20, "")
	if err != nil {
		t.Fatalf("❌ FeedV2 测试失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("❌ 返回的数据为空")
	}

	fmt.Printf("✅ FeedV2 测试成功，获取到 %d 字节的数据\n", len(data))

	// 验证数据是否为有效的JSON
	var feedData FeedV2Data
	if err := json.Unmarshal(data, &feedData); err != nil {
		t.Fatalf("❌ 返回的数据不是有效的JSON: %v", err)
	}

	fmt.Printf("✅ 解析成功，获取到 %d 条资讯\n", len(feedData.List))

	// 打印前几条资讯的标题
	for i, item := range feedData.List {
		if i >= 3 { // 只打印前3条
			break
		}
		if item.News != nil {
			fmt.Printf("  📰 新闻 %d: %s\n", i+1, item.News.Title)
		} else if item.Article != nil {
			fmt.Printf("  📄 文章 %d: %s\n", i+1, item.Article.Title)
		}
	}
}

// limitString 限制字符串长度的辅助函数
func limitString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
