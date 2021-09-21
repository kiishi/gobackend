FROM golang:latest as builder

WORKDIR /getir

ENV GO111MODULE=on

COPY . .

RUN go get

RUN CGO_ENABLED=0 go build -o backend .

FROM alpine:latest

WORKDIR /getir

COPY --from=builder /getir/backend .

CMD ["./backend"]