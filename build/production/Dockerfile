#build
FROM golang:latest as build
WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /build/main

# final
FROM alpine:3.9.4
WORKDIR /app
COPY --from=build build/main /app/
EXPOSE 8888
ENTRYPOINT ["./main"]