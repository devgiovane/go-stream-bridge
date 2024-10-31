package service

import (
	"study/application/model"
	"study/application/service/stream"
)

type Sync struct {
	StreamWorker *stream.Worker
}

func NewSync() *Sync {
	return &Sync{
		StreamWorker: stream.NewWorker(),
	}
}

func (s *Sync) Execute(cameras []*model.Camera) {
	for _, camera := range cameras {
		if camera.Action == "start" {
			s.StreamWorker.Add(camera)
		}
		if camera.Action == "stop" {
			s.StreamWorker.Remove(camera)
		}
	}
}
