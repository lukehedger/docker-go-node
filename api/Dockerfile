FROM golang:1.12-alpine as builder
COPY . /api
WORKDIR /api
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /api .

CMD ["./server"]
