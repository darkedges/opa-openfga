FROM golang:1.24.3 AS build

WORKDIR /build
COPY go/go.mod go/go.sum ./
RUN go mod download
COPY go/ ./
RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o opa

FROM alpine:latest
ARG USER=1000:1000
USER ${USER}
COPY --from=build /build/opa /opa
ENV PATH=${PATH}:/
ENTRYPOINT ["/opa"]
CMD ["run"]