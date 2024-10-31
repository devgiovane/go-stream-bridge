package command

import "os/exec"

type Stream struct {
}

func NewStream() *Stream {
	return &Stream{}
}

func (s *Stream) Make(rtsp string, rtmp string) *exec.Cmd {
	return exec.Command(
		"ffmpeg",
		"-loglevel", "panic", "-rtsp_transport", "tcp", "-i", rtsp, "-vcodec", "copy", "-map", "0:v", "-f", "flv", rtmp,
	)
}
