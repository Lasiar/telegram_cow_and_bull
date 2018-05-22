FROM amd64/golang as builder
RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
RUN dep ensure -vendor-only
COPY ./*.go /go/src/project/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main


FROM amd64/alpine
RUN mkdir /app
COPY --from=builder /go/src/project/main /app
ADD conf.json /app/conf.json
WORKDIR /app
CMD ["/app/main"]
