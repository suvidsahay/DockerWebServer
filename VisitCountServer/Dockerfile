FROM golang:latest

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get -u github.com/go-sql-driver/mysql
EXPOSE 8080
CMD ["go","run","server.go"]
