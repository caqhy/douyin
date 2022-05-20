package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dal/db"
	"time"
)

type FavoriteService struct {
}

// NewFavoriteService 创建服务
func NewFavoriteService() *FavoriteService {
	return &FavoriteService{}
}

func (f *FavoriteService) DoLike(userId int64, videoId int64) {
	key := fmt.Sprintf("douyin:favorite:video%d:user%d", videoId, userId)
	fmt.Println(key)
	_, err := db.Redis.Set(key, 1, time.Minute*5).Result()
	if err != nil {
		panic(err)
	}
	return
}

func (f *FavoriteService) CancelLike(userId int64, videoId int64) {
	key := fmt.Sprintf("douyin:favorite:video%d:user%d", videoId, userId)
	fmt.Println(key)
	_, err := db.Redis.Set(key, 0, time.Minute*5).Result()
	if err != nil {
		panic(err)
	}
	return
}
