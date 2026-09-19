# Sample: replace dockops-instance's nginx stub with a Go app.
# Build context is the repo root so this file can COPY ./app.
FROM golang:1.25-alpine AS builder

WORKDIR /go/src/app
COPY ./app ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /go/bin/
COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /go/src/app/config/parameters.yml.dist /go/bin/parameters.yml
EXPOSE 8080
CMD ["/go/bin/app", "-env", "prod"]
