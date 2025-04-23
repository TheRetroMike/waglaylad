# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/waglayla/waglaylad

WORKDIR /go/src/github.com/waglayla/waglaylad

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache astrixd dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o waglaylawallet ./cmd/waglaylawallet

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

COPY --from=build /go/src/github.com/waglayla/waglaylad/waglaylawallet /app/
#COPY --from=build /go/src/github.com/astrix-network/astrixd/infrastructure/config/sample-astrixd.conf /app/

WORKDIR /
CMD ["/app/waglaylawallet", "start-daemon"]
