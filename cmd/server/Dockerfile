FROM golang:latest

RUN go version 
ENV GOPATH=/

COPY ./ ./ 

RUN go mod download
RUN go build -o linkshortener ./cmd/server/main.go

CMD ["./linkshortener"]