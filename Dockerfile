FROM golang:alpine as builder
COPY main.go /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redirect

FROM scratch
COPY --from=builder /app/redirect /app/redirect
ENTRYPOINT ["/app/redirect"]
EXPOSE 8000
