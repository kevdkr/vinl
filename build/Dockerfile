FROM golang:1.20-alpine AS build

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . /app
RUN go build -o /vinl ./cmd/vinl

FROM scratch

WORKDIR /
COPY --from=build /vinl /vinl
EXPOSE 8080
ENTRYPOINT ["/vinl"]
