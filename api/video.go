package api

import (
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
	s := service.GetVideoService{}
	data := s.Get(c.Param("id"))
	c.JSON(http.StatusOK, data)
}


func UpdateVideo(c *gin.Context) {
	s := service.UpLoadVideoService{}
	if err := c.Should(&s); err == nil {
		res := s.Update(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteVideo(c *gin.Context) {
	s := service.DeleteVideoService{}
	res := s.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListVideo(c *gin.Context) {
	s := service.ListVIdeoService{}
	data := s.List()
	c.JSON(http.StatusOK, data)
}