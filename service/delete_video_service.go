package service

import (
	"giligili/model"
	"giligili/serializer"
)

type DeleteVideoService struct {}


func (s *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	model.DB.First(&video, id)
	if err := model.DB.Delete(&video).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg: "删除失败",
		}
	} else {
		return serializer.Response{
			Msg: "删除成功",
		}
	}

}
