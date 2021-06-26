FROM golang:1.16 as build-env

ENV GO111MODULE on

WORKDIR /go/src/cdn
ADD . /go/src/cdn

RUN go get -d -v ./...

RUN make
RUN mv ./cdn /go/bin/cdn
FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/cdn /
CMD ["/cdn"]