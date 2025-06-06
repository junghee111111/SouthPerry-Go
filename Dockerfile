FROM golang:1.24

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o server .

EXPOSE 8484

CMD ["./server"]
