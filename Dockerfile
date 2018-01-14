FROM golang:1.9-alpine as builder
RUN apk update && apk add git

ENV CGO_ENABLED 0
ENV LDFLAGS ""
WORKDIR "$GOPATH/src/github.com/srossross/k8s-test-controller"

COPY *.go ./
COPY pkg ./pkg
COPY ./vendor ./vendor
COPY ./Gopkg.lock ./
COPY ./Gopkg.toml ./

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go build -ldflags "${LDFLAGS}" -o /test-controller ./main.go

# ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
# ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

FROM alpine

COPY --from=builder /test-controller  /test-controller
CMD ["/test-controller"]
