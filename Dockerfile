FROM golang as builder
RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
RUN dep ensure -vendor-only
COPY ./*.go /go/src/project/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo  -o /go/bin/app


FROM alpine
RUN  apk add --no-cache --virtual ca-certificates
COPY --from=builder go/bin/app /app
ADD  conf.json /
CMD  ["/app"]
