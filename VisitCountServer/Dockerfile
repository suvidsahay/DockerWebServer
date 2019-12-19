FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -u github.com/go-sql-driver/mysql
EXPOSE 8080
CMD ["go","run","server.go"]
