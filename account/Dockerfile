FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

#TODO: not change sensitively
RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

# Rerun app on changes to any go file
RUN reflex -r "\.go$$" -s -- sh -c "go run ./"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
# More valuable command can be added here
CMD ["./run"]