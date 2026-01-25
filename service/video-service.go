package service

import "testpkg/ginserver/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() *videoService {
	return &videoService{}
}

func (service *videoService) Save(e entity.Video) entity.Video {
	service.videos = append(service.videos, e)
	return e
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
