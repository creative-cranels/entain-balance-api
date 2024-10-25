FROM golang:1.23 as base

WORKDIR /app

RUN apt-get update -y && apt-get install openssh-client zstd rsync make -y

COPY . .

ENV GO111MODULE=on

