FROM ubuntu:16.04

RUN \
  apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y \
      curl \
      git \
      make

RUN \
  curl -O https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz \
  && tar xf go1.8.1.linux-amd64.tar.gz -C /usr/local

RUN mkdir /var/go

ENV GOROOT /usr/local/go
ENV GOPATH /var/go
ENV GOBIN ${GOPATH}/bin
ENV PATH ${PATH}:${GOROOT}/bin:${GOBIN}
ENV WORKDIR ${GOPATH}/src/github.com/eggsbenjamin/image-service

RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}

