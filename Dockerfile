ARG buildImage=golang
ARG targetImage=ubuntu
FROM $buildImage AS build
COPY . .
RUN go build hello.go

FROM $targetImage
COPY --from=build /go/hello /hello
CMD [ "/hello" ]
