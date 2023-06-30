FROM golang:1.20
WORKDIR /app

# COPYを用いて資材を取得
COPY go.mod ./
COPY *.go ./

RUN go build -o /app/main
EXPOSE 8080
CMD ["/app/main"]
