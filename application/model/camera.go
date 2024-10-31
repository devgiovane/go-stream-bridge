package model

type Camera struct {
	Id     int    `json:"id"`
	Rtsp   string `json:"rtsp"`
	Rtmp   string `json:"rtmp"`
	Action string `json:"action"`
}
