FROM golang:1.12.2
COPY . .
RUN go build hello.go
CMD [ "hello" ]
