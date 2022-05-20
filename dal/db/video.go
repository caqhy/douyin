package db

import (
	"gorm.io/gorm"
)

// Video DO 对象
type Video struct {
	gorm.Model
	UserId        int64  // 发布者 id
	PlayUrl       string // 视频 url
	CoverUrl      string // 封面 url
	Tag           string // 视频标签
	FavoriteCount int64  // 点赞数量
	CommentCount  int64  // 评论数量
}

// CreateVideo 创建视频
func CreateVideo(video *Video) bool {
	return DB.Create(video).Error != nil
}

// GetPublishByUserId 通过用户 ID 获取该用户发布的视频列表
func GetPublishByUserId(userId int64) []Video {
	var videoList []Video
	DB.Where("user_id = ?", userId).Find(&videoList)
	return videoList
}

func GetVideo(videoId int64) Video {
	var video Video
	DB.Where("id = ?", videoId).First(&video)
	return video
}
