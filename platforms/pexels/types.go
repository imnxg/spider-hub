package pexels

// SearchResponse 搜索响应结构
type SearchResponse struct {
    PageProps struct {
        InitialData struct {
            Data []MediaItem `json:"data"`
        } `json:"initialData"`
    } `json:"pageProps"`
}

// MediaItem 媒体项目（视频或图片）
type MediaItem struct {
    ID         string `json:"id"`
    Type       string `json:"type"` // "video" 或 "photo"
    Attributes struct {
        ID          int     `json:"id"`
        Slug        string  `json:"slug"`
        Description string  `json:"description"`
        Width       int     `json:"width"`
        Height      int     `json:"height"`
        Status      string  `json:"status"`
        CreatedAt   string  `json:"created_at"`
        UpdatedAt   string  `json:"updated_at"`
        PublishAt   string  `json:"publish_at"`
        FeedAt      string  `json:"feed_at"`
        Title       string  `json:"title"`
        AspectRatio float64 `json:"aspect_ratio"`
        License     string  `json:"license"`
        Published   bool    `json:"published"`
        Starred     bool    `json:"starred"`
        Pending     bool    `json:"pending"`
        User        User    `json:"user"`
        Tags        []Tag   `json:"tags"`
        
        // 视频特有字段
        Duration    *int           `json:"duration,omitempty"`
        VideoFiles  []VideoFile    `json:"video_files,omitempty"`
        VideoPicture *VideoPicture `json:"video_picture,omitempty"`
        
        // 图片特有字段
        Src *PhotoSrc `json:"src,omitempty"`
    } `json:"attributes"`
}

// User 用户信息
type User struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Slug      string `json:"slug"`
    Username  string `json:"username"`
    Location  string `json:"location"`
    Avatar    struct {
        Small  string `json:"small"`
        Medium string `json:"medium"`
    } `json:"avatar"`
    Hero      bool `json:"hero"`
    Following bool `json:"following"`
}

// Tag 标签信息
type Tag struct {
    Name       string `json:"name"`
    SearchTerm string `json:"search_term"`
}

// VideoFile 视频文件信息
type VideoFile struct {
    ID       int    `json:"id"`
    Quality  string `json:"quality"`
    FileType string `json:"file_type"`
    Width    int    `json:"width"`
    Height   int    `json:"height"`
    Link     string `json:"link"`
}

// VideoPicture 视频缩略图
type VideoPicture struct {
    ID      int    `json:"id"`
    Picture string `json:"picture"`
    Nr      int    `json:"nr"`
}

// PhotoSrc 图片源信息
type PhotoSrc struct {
    Original  string `json:"original"`
    Large2x   string `json:"large2x"`
    Large     string `json:"large"`
    Medium    string `json:"medium"`
    Small     string `json:"small"`
    Portrait  string `json:"portrait"`
    Landscape string `json:"landscape"`
    Tiny      string `json:"tiny"`
}

// VideoDetailResponse 视频详情响应
type VideoDetailResponse struct {
    PageProps struct {
        ID     string `json:"id"`
        Medium struct {
            ID         string      `json:"id"`
            Type       string      `json:"type"`
            Attributes VideoDetail `json:"attributes"`
        } `json:"medium"`
    } `json:"pageProps"`
}

// VideoDetail 视频详情
type VideoDetail struct {
    ID          int           `json:"id"`
    Width       int           `json:"width"`
    Height      int           `json:"height"`
    Duration    int           `json:"duration"`
    FullRes     interface{}   `json:"full_res"`
    Tags        []string      `json:"tags"`
    URL         string        `json:"url"`
    Image       string        `json:"image"`
    AvgColor    string        `json:"avg_color"`
    User        User          `json:"user"`
    VideoFiles  []VideoFile   `json:"video_files"`
    VideoPictures []VideoPicture `json:"video_pictures"`
}

// BrowserFingerprint 浏览器指纹信息
type BrowserFingerprint struct {
    UserAgent      string            `json:"user_agent"`
    Headers        map[string]string `json:"headers"`
    TLSFingerprint string            `json:"tls_fingerprint"`
    HTTP2Settings  map[string]int    `json:"http2_settings"`
}