package controller

import "github.com/RaymondCode/simple-demo/model"

var DemoVideos = []model.Video{
	{
		Id:     1,
		Author: DemoUser,
		//PlayUrl: "https://www.w3schools.com/html/movie.mp4",
		PlayUrl: "http://39.98.41.126:31101/static/test.mp4",
		//CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CoverUrl: "http://39.98.41.126:31101/static/map1.jpg", FavoriteCount: 0,
		CommentCount: 0,
		IsFavorite:   false,
	},
	{
		Id:     2,
		Author: DemoUser,
		//PlayUrl: "https://www.w3schools.com/html/movie.mp4",
		PlayUrl: "http://39.98.41.126:31101/static/1_bear.mp4",
		//CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CoverUrl: "http://39.98.41.126:31101/static/map1.jpg", FavoriteCount: 0,
		CommentCount: 0,
		IsFavorite:   false,
	},
}

var DemoComments = []model.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = model.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
