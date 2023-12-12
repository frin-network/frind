# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/frin-network/frind/

WORKDIR /go/src/github.com/frin-network/frind/

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache frind dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o frind .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/frin-network/frind/frind /app/
COPY --from=build /go/src/github.com/frin-network/frind/infrastructure/config/sample-frind.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
