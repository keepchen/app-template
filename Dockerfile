FROM golang:latest AS builder

ARG COMMIT_ID
ARG VERSION=""
ARG VCS_BRANCH=""
ARG GRPC_STUB_REVISION=""
ARG PROJECT_NAME=app-template
ARG DOCKER_PROJECT_DIR=/build
ARG EXTRA_BUILD_ARGS=""
ARG GOCACHE=""
ARG GOPROXY=https://goproxy.cn,direct
ARG GOSUMDB=off

WORKDIR $DOCKER_PROJECT_DIR
COPY . $DOCKER_PROJECT_DIR

ENV GOPROXY=$GOPROXY
ENV GOSUMDB=$GOSUMDB

RUN mkdir -p /output \
    && make build -e GOCACHE=$GOCACHE \
    -e COMMIT_ID=$COMMIT_ID -e OUTPUT_FILE=/output/app-template \
    -e VERSION=$VERSION -e VCS_BRANCH=$VCS_BRANCH -e EXTRA_BUILD_ARGS=$EXTRA_BUILD_ARGS

FROM debian:buster-slim
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Shanghai
ARG SERVE_MODE
ENV SUB_CMD=SERVE_MODE

COPY --from=builder /output/app-template /usr/bin/app-template
# COPY pkg/app/[x]/config/config.yaml /data

CMD app-template ${SUB_CMD}