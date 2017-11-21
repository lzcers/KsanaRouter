package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// todo 应该再构建一层抽象屏障，屏蔽数据库

// Post 文章
type Post struct {
	ID          bson.ObjectId `json:"-" bson:"_id"`
	PostName    string        `json:"PostName" bson:"PostName"`
	Tags        []string      `json:"Tags" bson:"Tags"`
	Content     string        `json:"Content" bson:"Content"`
	PublishDate string        `json:"PublishDate" bson:"PublishDate"`
	LastUpdate  string        `json:"LastUpdate" bson:"LastUpdate"`
}

// AddPost 写入一篇文章至数据库
func AddPost(p Post) {
	p.ID = bson.NewObjectId()
	DB.C("posts").Insert(&p)
}

// GetPost 根据ID获取文章，ID为空则获取所有文章
func GetPost(pID string) []Post {
	var (
		result []Post
		err    error
	)
	if pID != "" {
		fmt.Println(pID)
		err = DB.C("posts").FindId(bson.ObjectIdHex(pID)).All(&result)
	} else {
		err = DB.C("posts").Find(nil).All(&result)
	}
	if err != nil {
		// 我他妈也不知道该做啥
	}
	return result
}

// GetTags 获取所有标签
func GetTags() struct {
	Tags []string `bson:"Tags"`
} {
	var (
		mapTags map[string]bool
		tags    struct {
			Tags []string `bson:"Tags"`
		}
		result []struct {
			Tags []string `bson:"Tags"`
		}
	)
	mapTags = make(map[string]bool)
	err := DB.C("posts").Find(nil).Select(bson.M{"Tags": 1}).All(&result)
	if err != nil {
		// 我他妈也不知道该做啥
	}
	// 去重
	for _, i := range result {
		for _, t := range i.Tags {
			mapTags[t] = true
		}
	}
	for tagName := range mapTags {
		tags.Tags = append(tags.Tags, tagName)
	}
	return tags
}
