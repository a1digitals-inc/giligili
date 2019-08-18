package serializer

import "giligili/model"

type Video struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Info string `json:"info"`
	CreatedAt int64 `json:"created_at"`
}


func BuildVideo(item model.Video) Video {
	return Video{
		Title: item.Title,
		Info: item.Info,
		ID: item.ID,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildVideos(item []model.Video) []Video {
	var res []Video
	for _, video := range item {
		res = append(res, BuildVideo(video))
	}
	return res
}