FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary


# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
ENTRYPOINT ["/app/binary"]