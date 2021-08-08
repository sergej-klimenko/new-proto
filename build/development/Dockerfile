#build
FROM golang:latest as build
WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o /build/main

# final
FROM alpine:3.9.4
WORKDIR /app
COPY --from=build build/main /app/
ENV ENVIRONMENT=development-docker
EXPOSE 8888
CMD ["./main"]