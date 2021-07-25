# builder image
FROM golang:latest AS builder

# copy source code
WORKDIR /src/
COPY . .

# fetch dependencies
RUN go mod download

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARC=amd64 go build -a -installsuffix cgo -o app


# base image
FROM alpine:latest

# Security related package
RUN apk --no-cache add ca-certificates

# copy the binary
COPY --from=builder /src/app .

# run the binary
ENTRYPOINT [ "./app" ]