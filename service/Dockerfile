FROM docker.io/golang:1.16-buster AS build
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o portfolio-service

FROM docker.io/alpine:3.14
COPY --from=build /build/portfolio-service portfolio-service
EXPOSE 8080
ENTRYPOINT ["./portfolio-service"]