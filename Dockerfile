FROM golang:1.23

LABEL maintainer="By Mohsen Taheri"
LABEL Email="m.rozbehano@outlook.com"

RUN apt update && apt install vim -y

ENV APP_ENV=test
ENV PORT=8585
ENV IP="0.0.0.0"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

EXPOSE 6985

WORKDIR $GOPATH/src/gateway

RUN mkdir -p ./config/file

COPY ./cmd/gateway .

COPY ./config/file/* ./config/file

CMD ["./gateway"]