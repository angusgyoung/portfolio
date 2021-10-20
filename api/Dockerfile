FROM docker.io/golang:1.16-buster AS build
WORKDIR /build
COPY . .
RUN go mod download

ARG version

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s -X main.Version=$version" -o portfolio-api

FROM docker.io/alpine:3.14
COPY --from=build /build/portfolio-api portfolio-api
EXPOSE 8080
ENTRYPOINT ["./portfolio-api"]