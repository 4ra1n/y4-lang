FROM ubuntu:22.04

LABEL author="4ra1n"
LABEL github="https://github.com/4ra1n"

ENV Y4_VERSION v0.0.3

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates --reinstall

RUN echo "\
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy main restricted universe multiverse\n\
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-updates main restricted universe multiverse\n\
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse\n\
deb http://security.ubuntu.com/ubuntu/ jammy-security main restricted universe multiverse" | tee /etc/apt/sources.list \
    && apt-get update && apt-get install -y zip unzip upx wget

RUN rm -rf /usr/local/go \
	&& wget https://mirrors.aliyun.com/golang/go1.21.6.linux-amd64.tar.gz --no-check-certificate \
	&& tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz  \
	&& export PATH=$PATH:/usr/local/go/bin \
	&& . ~/.profile \
	&& go version

ENV PATH="${PATH}:/usr/local/go/bin"

COPY . .

RUN cd gox && \
    go build -o ../cmd/my-gox

RUN cd cmd && \
    mkdir y4 && \
    ./my-gox -parallel 5 -osarch="darwin/arm64 darwin/amd64 linux/386 linux/amd64 linux/arm linux/arm64 windows/arm windows/arm64 windows/386 windows/amd64" -ldflags="-extldflags=-static -s -w" -output="y4/y4lang_${Y4_VERSION}_{{.OS}}_{{.Arch}}" && \
    find y4 -type f -exec sh -c ' \
          upx "$1" && \
          sha256sum "$1" > "$1_sha256.txt" \
        ' sh {} \; && \
    zip -r ../build.zip y4/*

CMD ["echo", "build y4-lang ${Y4_VERSION} completed - /app/build.zip"]

