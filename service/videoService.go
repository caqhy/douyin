package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dal/db"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
)

type VideoService struct {
}

// NewVideoService 创建服务
func NewVideoService() *VideoService {
	return &VideoService{}
}

// PublishVideo 发布视频
// 发布成功则返回 fileName，否则抛出错误 err
func (v *VideoService) PublishVideo(user model.User, file *multipart.FileHeader, c *gin.Context) (finalName string, err error) {
	// 获取文件名以及生成最终的文件名
	filename := filepath.Base(file.Filename)
	finalName = fmt.Sprintf("%d_%s", user.Id, filename)

	// 设置存放路径，并保存到服务器
	saveFile := filepath.Join("./public/", finalName)
	if err = c.SaveUploadedFile(file, saveFile); err != nil {
		return
	}

	// TODO 获得各种视频信息 ...
	playUrl := "test"
	coverUrl := "test"

	// 将视频信息保存到数据库
	db.CreateVideo(user, playUrl, coverUrl)
	return
}
