FROM golang:1.13-alpine

RUN apk update && \
                apk add tzdata && \
                cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime && \
                echo "Europe/Moscow" >  /etc/timezone && \
                apk del tzdata && \
                rm -rf /var/cache/apk/*

ENV PROJECT_NAME stock-issuers-server
ENV SOURCE_PATH /go/src/github.com/korhov/stock-issuers-server
ENV BUILD_PATH /etc

RUN mkdir -pv $BUILD_PATH && mkdir -pv $SOURCE_PATH

ADD source.tar.gz $SOURCE_PATH

RUN ls -la $SOURCE_PATH
RUN ls -la $SOURCE_PATH/lint.sh
RUN $SOURCE_PATH/lint.sh

# вендоры с нами. Билдим
RUN CGO_ENABLED=0 go build -v -ldflags="-s -w" -o $BUILD_PATH/$PROJECT_NAME $SOURCE_PATH/cmd/main.go

RUN rm -rf $SOURCE_PATH
RUN chmod +x $BUILD_PATH/$PROJECT_NAME

ADD start.sh /etc/start.sh
RUN chmod +x /etc/start.sh

ENTRYPOINT /etc/start.sh