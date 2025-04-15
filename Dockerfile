FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY db db
COPY handler handler
COPY l l
COPY middleware middleware
COPY util util
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM scratch
WORKDIR /app
COPY .env.template ./.env
COPY --from=build /server ./server
ENTRYPOINT [ "./server" ]
