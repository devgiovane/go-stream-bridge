FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 go build main.go

FROM scratch
WORKDIR /go/src
COPY --from=builder /go/src/main /go/src
ENTRYPOINT  [ "./main" ]