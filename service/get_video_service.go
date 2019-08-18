package service

import (
	"giligili/model"
	"giligili/serializer"
)

type GetVideoService struct {}


func (s *GetVideoService) Get(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Status: 5001,
			Data:   nil,
			Msg:    "记录不存在",
		}
	} else {
		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	}
}
