package service

import (
	"giligili/model"
	"giligili/serializer"
)

type UpLoadVideoService struct {
	Title string `form:"title" json:"title" binding:"min=0,max=200"`
	Info string `form:"info" json:"info" binding:"min=0,max=200"`
}


func (s *UpLoadVideoService) Update(id string) serializer.Response {
	var video model.Video
	if err := model.DB.Where("id = ?", id).First(&video).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg: "参数有误",
		}
	} else {
		video.Title = s.Title
		video.Info = s.Info
		if err := model.DB.Save(&video).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg: "更新失败",
			}
		} else {
			return serializer.Response{
				Data: serializer.BuildVideo(video),
			}
		}
	}
}