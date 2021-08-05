#build
FROM golang:latest as build
WORKDIR /build
COPY ../api /build/
RUN go build -o /build/main

# final
FROM alpine:3.9.4
WORKDIR /app
COPY --from=build /build/main /app/
EXPOSE 8888
ENTRYPOINT ["./main"]
