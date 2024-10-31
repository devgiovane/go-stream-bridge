package stream

import (
	"study/application/model"
	"study/infrastructure/command"
	"sync"
)

type Worker struct {
	wg      sync.WaitGroup
	streams map[int]*Stream
}

func NewWorker() *Worker {
	return &Worker{
		wg:      sync.WaitGroup{},
		streams: make(map[int]*Stream),
	}
}

func (w *Worker) Wait() {
	w.wg.Wait()
}

func (w *Worker) exists(c *model.Camera) bool {
	return w.streams[c.Id] != nil
}

func (w *Worker) Add(c *model.Camera) {
	if w.exists(c) {
		return
	}
	w.wg.Add(1)
	stream := NewStream(c.Id, command.NewStream().Make(c.Rtsp, c.Rtmp))
	w.streams[c.Id] = stream
	go stream.Start(&w.wg)
}

func (w *Worker) Remove(c *model.Camera) {
	if !w.exists(c) {
		return
	}
	w.streams[c.Id].stop <- true
	delete(w.streams, c.Id)
}
