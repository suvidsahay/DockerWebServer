FROM golang:onbuild

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]