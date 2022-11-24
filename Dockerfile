FROM golang:1.19.3-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o gorcon app/*.go
CMD ["/app/gorcon"]