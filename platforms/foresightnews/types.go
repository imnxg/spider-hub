package foresightnews

// FeedV2Res 表示API响应的顶层结构
type FeedV2Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"` // Base64+Zlib压缩的数据
}

// FeedV2Data 表示解密后的实际数据结构
type FeedV2Data struct {
	List []FeedItem `json:"list"`
}

// FeedItem 表示单个资讯项目
type FeedItem struct {
	ID          int64        `json:"id"`
	SourceID    int64        `json:"source_id"`
	SourceType  string       `json:"source_type"`
	PublishedAt int64        `json:"published_at"`
	News        *NewsItem    `json:"news,omitempty"`
	Article     *ArticleItem `json:"article,omitempty"`
}

// NewsItem 表示新闻类型的资讯
type NewsItem struct {
	ID           int64       `json:"id"`
	Title        string      `json:"title"`
	Brief        string      `json:"brief"`
	Content      string      `json:"content"`
	Img          string      `json:"img"`
	Tags         []Tag       `json:"tags"`
	IsImportant  bool        `json:"is_important"`
	ImportantTag *Tag        `json:"important_tag,omitempty"`
	Label        string      `json:"label"`
	SourceLink   string      `json:"source_link"`
	PublishedAt  int64       `json:"published_at"`
	Favorited    bool        `json:"favorited"`
	Wikis        interface{} `json:"wikis"`
	NewWikis     interface{} `json:"new_wikis"`
}

// ArticleItem 表示文章类型的资讯
type ArticleItem struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Brief       string      `json:"brief"`
	Content     string      `json:"content"`
	Img         string      `json:"img"`
	SourceLink  string      `json:"source_link"`
	Column      *Column     `json:"column,omitempty"`
	Author      *Author     `json:"author,omitempty"`
	Tags        interface{} `json:"tags"`
	IsTop       bool        `json:"is_top"`
	IsImportant bool        `json:"is_important"`
	PublishedAt int64       `json:"published_at"`
	Favorited   bool        `json:"favorited"`
	Status      string      `json:"status"`
	CanEdit     bool        `json:"can_edit"`
	Wikis       interface{} `json:"wikis"`
	NewWikis    interface{} `json:"new_wikis"`
}

// Tag 表示标签信息
type Tag struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Subscribed bool   `json:"subscribed"`
}

// Column 表示专栏信息
type Column struct {
	ID               int64       `json:"id"`
	Title            string      `json:"title"`
	Brief            string      `json:"brief"`
	Content          string      `json:"content"`
	Img              string      `json:"img"`
	DefaultCover     string      `json:"default_cover"`
	IsRecommend      bool        `json:"is_recommend"`
	Weight           int         `json:"weight"`
	Subscribed       bool        `json:"subscribed"`
	EnablePush       bool        `json:"enable_push"`
	CreatedAt        int64       `json:"created_at"`
	LastUpdateAt     int64       `json:"last_update_at"`
	LastArticleTitle string      `json:"last_article_title"`
	LastArticleID    string      `json:"last_article_id"`
	Total            int         `json:"total"`
	Articles         interface{} `json:"articles"`
}

// Author 表示作者信息
type Author struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Brief    string `json:"brief"`
}
