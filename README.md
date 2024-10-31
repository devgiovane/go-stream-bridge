<img src="./go.png" width="80" height="80" alt="logo" />

# GO Push RTSP to RTMP

### About

> My study of the wait groups and goroutine applied in push service RTSP to RTMP

### Concepts

- Wait groups
- Goroutines
- Ffmpeg
- RabbitMQ

### Commands

create queue `sync_camera`

```bash
sudo docker compose up -d
```

```bash
go run main.go
```

### Create by
© [Giovane Santos](https://giovanesantossilva.github.io/)