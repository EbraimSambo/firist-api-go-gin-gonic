FROM golang:1:18

WORKDIR /o/src/app

COPY . .

EXPOSE 8000

RUN go build -o main cmd/main.go

CMD [ "./main" ]