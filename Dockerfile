FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o gologin-edufund

EXPOSE 8080

CMD ./gologin-edufund