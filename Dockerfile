FROM golang:1.21 AS build
WORKDIR /api-go-rest
COPY * ./
RUN go build -o /api-go-rest
EXPOSE 8000

CMD ["/api-go-rest"]ping www.google.