package pexels

// API相关常量
const (
    // NextData API基础URL
    NextDataBaseURL = "https://www.pexels.com/_next/data/0IujY-GsVsfR750ZpT6Yt"
    
    // 默认分页大小
    DefaultPageSize = 15
    MaxPageSize     = 80
)

// Safari浏览器指纹相关常量
const (
    SafariUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15"
)

// 请求头配置
var DefaultHeaders = map[string]string{
    "Accept":             "application/json, text/plain, */*",
    "Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8",
    "Accept-Encoding":    "gzip, deflate, br",
    "Cache-Control":      "no-cache",
    "Pragma":             "no-cache",
    "Sec-Fetch-Dest":     "empty",
    "Sec-Fetch-Mode":     "cors",
    "Sec-Fetch-Site":     "same-origin",
    "X-Requested-With":   "XMLHttpRequest",
}

// 媒体类型
const (
    MediaTypeVideo = "video"
    MediaTypePhoto = "photo"
)

// 视频质量选项
const (
    QualityHD     = "hd"
    QualitySD     = "sd"
    QualityMobile = "mobile"
)

// 错误消息
const (
    ErrInvalidQuery    = "搜索关键词不能为空"
    ErrInvalidVideoID  = "视频ID无效"
    ErrRequestFailed   = "请求失败"
    ErrParseResponse   = "解析响应失败"
    ErrFingerprintFail = "浏览器指纹识别失败"
)