FROM golang:1.18-alpine as builder

RUN adduser -D -g '' elf

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && \
    go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service .

FROM alpine:latest as runner

LABEL language="golang"
LABEL org.opencontainers.image.source https://github.com/guionardo/auth-service
# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
# copy the static executable
COPY --from=builder --chown=elf:1000 /auth-service /auth-service
# use a non-root user
USER elf

EXPOSE 3000

# run app
ENTRYPOINT ["/auth-service"]
