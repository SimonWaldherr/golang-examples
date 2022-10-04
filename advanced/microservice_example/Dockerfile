FROM golang:alpine 

RUN mkdir /app 

ADD . /app 

WORKDIR /app 

RUN go build -o main . 

EXPOSE 8080 

CMD ["/app/main"]
