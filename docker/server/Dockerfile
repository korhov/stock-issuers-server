FROM golang:1.12-alpine

ENV PROJECT_NAME stock-issuers-server
ENV SOURCE_PATH /go/src/github.com/korhov/stock-issuers-server
ENV BUILD_PATH /etc


RUN apk update && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime && \
    echo "Europe/Moscow" >  /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*


WORKDIR /home/$PROJECT_NAME
RUN mkdir -pv /home/$PROJECT_NAME/data \
# && curl https://www.moex.com/ru/listing/securities-list-csv.aspx?type=1 --output ./data/securities-list.csv \
 && wget https://www.moex.com/ru/listing/securities-list-csv.aspx?type=1 -O ./data/securities-list.csv \
 && iconv -f CP1251 -t UTF-8 ./data/securities-list.csv > ./data/securities-list-utf-8.csv \
 && mv ./data/securities-list-utf-8.csv ./data/securities-list.csv


RUN mkdir -pv $BUILD_PATH && mkdir -pv $SOURCE_PATH
ADD source.tar.gz $SOURCE_PATH

RUN CGO_ENABLED=0 go build -v -ldflags="-s -w" -o $BUILD_PATH/$PROJECT_NAME $SOURCE_PATH/cmd/main.go


RUN chmod +x $BUILD_PATH/$PROJECT_NAME \
 && rm -rf $SOURCE_PATH

ENTRYPOINT $BUILD_PATH/$PROJECT_NAME
