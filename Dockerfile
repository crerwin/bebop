FROM golang:1.18.6-alpine AS build

WORKDIR /bebop

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go install ./...

# Final image
FROM alpine:3.16.2

# copy binary
COPY --from=build /go/bin/bebop /usr/local/bin/bebop

ENTRYPOINT [ "bebop" ]