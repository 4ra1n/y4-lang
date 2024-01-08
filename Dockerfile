FROM golang:1.21.5

WORKDIR /app

RUN apt-get update && apt-get install -y zip unzip

COPY . .

RUN cd gox && \
    go build -o ../cmd/my-gox

RUN cd cmd && \
    mkdir y4 && \
    ./my-gox -parallel 5 -osarch="darwin/arm64 darwin/amd64 linux/386 linux/amd64 linux/arm linux/arm64 windows/arm windows/arm64 windows/386 windows/amd64" -ldflags="-extldflags=-static -s -w" -output="y4/y4lang_{{.OS}}_{{.Arch}}" && \
    zip -r ../build.zip y4/*

CMD ["echo", "build completed - /app/build.zip"]

