# alpine linux
FROM alpine:3.19

ENV APP_BIN=lrbooks_web
ARG SERVER_DIR=/home/.server
WORKDIR $SERVER_DIR
COPY ./bin/${APP_BIN} .
COPY ./service/web/adapter/templates/ ./templates

ENV GIN_MODE=release

CMD ./${APP_BIN}