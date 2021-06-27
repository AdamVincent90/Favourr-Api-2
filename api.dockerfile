FROM golang:1.16

WORKDIR /app

RUN go get github.com/cespare/reflex

EXPOSE 4002
EXPOSE 4000

CMD [ "reflex", "-c", "reflex.conf" ]