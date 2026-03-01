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

