FROM golang:1.16

WORKDIR /app

RUN go get github.com/cespare/reflex

EXPOSE 4002

CMD [ "reflex", "-c", "config/reflex.conf" ]