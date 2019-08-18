package service

import (
	"giligili/model"
	"giligili/serializer"
)

type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=20"`
	Info string `form:"info" json:"info" binding:"required,min=0,max=200"`
}

// 视频投稿
func (s *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: s.Title,
		Info: s.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 5001,
			Msg: "数据库连接出错",
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
