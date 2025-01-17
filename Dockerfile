# Start from the latest golang base image
FROM golang:latest as builder
MAINTAINER Valentin Kuznetsov vkuznet@gmail.com
ENV WDIR=/data
WORKDIR $WDIR

# RUN go get github.com/vkuznet/gpm
ARG CGO_ENABLED=0
RUN git clone https://github.com/vkuznet/gpm.git && cd gpm && make

# FROM alpine
# RUN mkdir -p /data
# https://blog.baeke.info/2021/03/28/distroless-or-scratch-for-go-apps/
FROM gcr.io/distroless/static AS final
# COPY --from=builder /data/gpm/gpm /data/
