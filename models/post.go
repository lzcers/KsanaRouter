package models

// Post 文章
type Post struct {
	PostNane    string `json:"PostNane"`
	Tags        string `json:"Tags"`
	Content     string `json:"Content"`
	PublishDate string `json:"PublishDate"`
	LastUpdate  string `json:"LastUpdate"`
}

// AddPost 写入一篇文章至数据库
// todo 应该再构建一层抽象屏障，屏蔽数据库
func AddPost(p Post) {
	DB.C("posts").Insert(&p)
}
