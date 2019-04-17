FROM golang:1.12-alpine AS build-base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/liam-j-bennett/mather
ENV GO111MODULE=on

ENV PORT=5000
ENV API_KEY=some-api-key

COPY go.mod .
COPY go.sum .
RUN go mod download

FROM build-base AS build-env

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o matherapi ./cmd/matherapi

FROM alpine:latest AS runner

COPY --from=build-env /go/src/github.com/liam-j-bennett/mather/matherapi .
CMD ["/matherapi"]
