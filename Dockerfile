FROM golang:1.12.2 AS build
COPY . .
RUN go build hello.go

FROM ubuntu
COPY --from=build /go/hello /hello
CMD [ "/hello" ]
