FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o breadcrumbs ./cmd/breadcrumbs/
EXPOSE 80
CMD ["./breadcrumbs"]