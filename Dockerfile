FROM golang:alpine AS build

RUN mkdir /go/src/mac2vendor

ADD go.mod main.go parse.go formatters.go /go/src/mac2vendor/

ADD radix /go/src/mac2vendor/radix/

RUN cd /go/src/mac2vendor && CGO_ENABLED=0 go build .

RUN apk --no-cache add curl binutils && \
    curl -sL 'https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf;hb=HEAD' | \
    gzip -c9 > /tmp/oui.txt.gz && \
    strip -s /go/src/mac2vendor/mac2vendor

FROM scratch

COPY --from=build /go/src/mac2vendor/mac2vendor /
COPY --from=build /tmp/oui.txt.gz /

USER 65534

EXPOSE 3000

ENTRYPOINT ["/mac2vendor"]

CMD ["/oui.txt.gz"]
