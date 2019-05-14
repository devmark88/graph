FROM golang:alpine 
RUN mkdir /app 
RUN apk add git alpine-sdk openssh bash \
            libxml2-dev libxslt-dev \
            linux-headers && \
    rm -rf /var/cache/apk/*

ADD . /app/ 
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]