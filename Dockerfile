FROM golang:1.22.1-alpine

COPY . .


RUN go get .
RUN go build

EXPOSE 8080

CMD ["./airport-app-backend"]