# FROM golang:1.12
# ENV GO111MODULE=on
# WORKDIR /app
# COPY . .
# # COPY go.mod .
# # COPY go.sum .
# # RUN go mod download
# # RUN GOBIN=/ go install ./...
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOBIN=/ go build ./...
# EXPOSE 8080
# CMD ["./breadcrumbs"]

FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o breadcrumbs ./cmd/breadcrumbs/

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=builder /app/breadcrumbs .
# COPY --from=builder /app/.env .  
# COPY --from=builder /app/web .  
# COPY --from=builder /app/templates .  
EXPOSE 8080
CMD ["./breadcrumbs"]