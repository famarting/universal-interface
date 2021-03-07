FROM golang:latest

WORKDIR /go/src/github.com/famartinrh/uninterface/

COPY cmd cmd
COPY pkg pkg
COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/uninterface

FROM scratch
COPY --from=0 /go/src/github.com/famartinrh/uninterface/uninterface .
EXPOSE 7878
CMD ["/uninterface"]