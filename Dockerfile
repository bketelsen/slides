FROM alpine:3.7

WORKDIR /app

COPY . /app
ENV GIN_MODE=release

CMD ./hacker-slides $PORT
