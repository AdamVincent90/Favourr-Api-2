FROM golang:1.16

WORKDIR /app

COPY . /app/

RUN go get github.com/cespare/reflex 
RUN go mod tidy
RUN go build main.go

EXPOSE 4002
EXPOSE 4000

CMD [ "reflex", "-c", "reflex.conf" ]


