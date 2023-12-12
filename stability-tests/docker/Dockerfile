ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as frind
FROM ${KASPAMINER_IMAGE} as frinminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/frin-network/frind/

WORKDIR /go/src/github.com/frin-network/frind/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=frind /app/ /app/
COPY --from=frinminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/frin-network/frind/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]