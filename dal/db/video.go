package db

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

// Video DO 对象
type Video struct {
	gorm.Model
	Id            int64  // 视频 id
	UserId        int64  // 发布者 id
	PlayUrl       string // 视频 url
	CoverUrl      string // 封面 url
	Tag           string // 视频标签
	FavoriteCount int64  // 点赞数量
	CommentCount  int64  // 评论数量
}

// CreateVideo 创建视频
func CreateVideo(user model.User, playUrl, coverUrl string) bool {
	// TODO 持久化逻辑
	return true
}
