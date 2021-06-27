FROM golang:1.16

WORKDIR /app

RUN go get github.com/cespare/reflex 
RUN go mod tidy
RUN go build main.go

EXPOSE 4000
EXPOSE 8080

CMD [ "reflex", "-c", "reflex.conf" ]


