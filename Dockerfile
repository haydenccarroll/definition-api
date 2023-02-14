FROM golang:1.20.0-alpine3.17

LABEL maintainer="Hayden Carroll"

RUN mkdir /app
WORKDIR /app

COPY . .

RUN apk add --update make
RUN make definition_api

CMD ["definition_api", "serve"]
