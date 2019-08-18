package api

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


func CreateVideo(c *gin.Context) {
	s := service.CreateVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}


func GetVideo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if ok {
		var video model.Video
		if err := model.DB.Where("id = ?", id).First(&video).Error; err != nil {
			c.JSON(http.StatusOK, &serializer.Response{
				Status: 400001,
				Msg: "记录不存在"})
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Data: serializer.BuildVideo(video),
			})
		}

	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: 400001,
			Msg: "获取记录ID出错",
		})
	}
}


func UpdateVideo(c *gin.Context) {
	s := service.UpLoadVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Update(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteVideo(c *gin.Context) {
	s := service.DeleteVideoService()
	res := s.Delete(c)
	c.JSON(http.StatusOK, res)
}