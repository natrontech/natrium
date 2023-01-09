# Pocketbase backend build
FROM golang:1.19 as build
WORKDIR /go/src/github.com/natrongmbh/natrium
# RUN go get -d -v golang.org/x/net/html
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o natrium ./cmd/natrium

# Production build
FROM alpine:3.16

WORKDIR /app

RUN addgroup -g 1001 -S natrium
RUN adduser -S natrium -u 1001

RUN apk --no-cache add ca-certificates

COPY --from=build --chown=natrium:natrium /go/src/github.com/natrongmbh/natrium/natrium ./
COPY --from=build --chown=natrium:natrium /go/src/github.com/natrongmbh/natrium/config.yaml ./config.yaml

EXPOSE 8090

CMD ["./natrium", "serve", "--http=0.0.0.0:8090"]
