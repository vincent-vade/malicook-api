FROM golang:1.20-alpine as builder

RUN apk --no-cache add curl

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
RUN chmod +x /usr/local/bin/dbmate

RUN dbmate migrate

RUN GCO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./http/main ./http/

FROM scratch

COPY --from=builder /app/http .

EXPOSE 5003
#de
CMD [ "./main" ]

