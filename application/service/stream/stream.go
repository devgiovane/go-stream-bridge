package stream

import (
	"fmt"
	"os/exec"
	"sync"
)

type Stream struct {
	id   int
	cmd  *exec.Cmd
	stop chan bool
}

func NewStream(id int, cmd *exec.Cmd) *Stream {
	return &Stream{
		id:   id,
		cmd:  cmd,
		stop: make(chan bool),
	}
}

func (s *Stream) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("~[Stream] goroutine %d start.\n", s.id)
	if err := s.cmd.Start(); err != nil {
		fmt.Printf("~[Stream] goroutine %d error in start ffmpeg: %v\n", s.id, err)
	}
	go func() {
		s.Manage()
	}()
	if err := s.cmd.Wait(); err != nil {
		fmt.Printf("~[Stream] goroutine %d error in wait ffmpeg: %v\n", s.id, err)
	}
}

func (s *Stream) Manage() {
	for {
		select {
		case <-s.stop:
			fmt.Printf("~[Stream] goroutine %d stopped.\n", s.id)
			if err := s.cmd.Process.Kill(); err != nil {
				fmt.Printf("~[Stream] goroutine %d error in kill ffmpeg: %v\n", s.id, err)
			}
			return
		}
	}
}
