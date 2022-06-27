FROM golang:1.17 as builder
WORKDIR /app
COPY . /app

COPY go.mod ./
RUN go mod download
COPY . .

# Statically compile our app for use in a distroless container
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -v -o app .

# A distroless container image with some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static

COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]