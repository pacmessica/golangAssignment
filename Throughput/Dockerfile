FROM alpine:3.4

RUN apk update
RUN apk add go
RUN mkdir /app
ADD Main.go /app/
WORKDIR /app
RUN go run Main.go
CMD ["/app/main"]
