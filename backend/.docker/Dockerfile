### STAGE 1: Build ###
FROM golang:1.19.4-alpine3.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /main

### STAGE 2: Run ###
FROM golang:1.19.4-alpine3.17

COPY --from=build /main ./main

EXPOSE 80

CMD [ "./main" ]