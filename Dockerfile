# builder
FROM golang:latest AS builder
RUN apt-get update 
WORKDIR /app
COPY . .
RUN go build -o main 

EXPOSE 6883

CMD [ "./main", "s" ]
# start PocketBase