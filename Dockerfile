FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY service ./service
COPY util ./util

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /yktool


FROM alpine

WORKDIR /

COPY --from=build /yktool /yktool

ENTRYPOINT ["/yktool"]