FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go mod tidy
RUN go build -o main .

ENV PORT=8000

EXPOSE 8000

CMD [ "./main" ]