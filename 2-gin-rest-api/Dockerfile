# The build stage
FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api main.go

# The run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/go-api .
EXPOSE 8080
CMD ["./go-api"]