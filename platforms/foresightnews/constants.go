package foresightnews

const (
	// BaseURL API基础URL
	BaseURL = "https://api.foresightnews.pro"

	// FeedV2Endpoint 资讯列表接口端点
	FeedV2Endpoint = "/v2/feed"

	// DefaultUserAgent 默认User-Agent
	DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

	// DefaultReferer 默认Referer
	DefaultReferer = "https://s.foresightnews.pro/"

	// DefaultOrigin 默认Origin
	DefaultOrigin = "https://s.foresightnews.pro"

	// DefaultPageSize 默认分页大小
	DefaultPageSize = 20

	// MaxPageSize 最大分页大小
	MaxPageSize = 100
)

// HTTPHeaders 默认请求头配置
var HTTPHeaders = map[string]string{
	"accept":             "application/json, text/plain, */*",
	"accept-language":    "zh-CN,zh;q=0.9",
	"origin":             DefaultOrigin,
	"priority":           "u=1, i",
	"referer":            DefaultReferer,
	"sec-ch-ua":          "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\"",
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": "\"Windows\"",
	"sec-fetch-dest":     "empty",
	"sec-fetch-mode":     "cors",
	"sec-fetch-site":     "same-site",
	"user-agent":         DefaultUserAgent,
	"x-requested-with":   "XMLHttpRequest",
}
