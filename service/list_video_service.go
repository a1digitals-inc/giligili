package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ListVIdeoService struct {
	Title string `form:"title" json:"title"`
	Info string `form:"info" json:"info"`
}

func (s *ListVIdeoService) List() serializer.Response {
	var videos []model.Video
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50001,
			Data:   nil,
			Msg:    "连接数据库出错",
			Error:  err.Error(),
		}
	} else {
		return serializer.Response{
			Data: serializer.BuildVideos(videos),
		}
	}
}
