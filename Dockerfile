FROM golang:alpine AS build

RUN mkdir /go/src/mac2vendor

ADD go.mod main.go parse.go formatters.go /go/src/mac2vendor/

ADD radix /go/src/mac2vendor/radix/

RUN cd /go/src/mac2vendor && CGO_ENABLED=0 go build .

RUN apk --no-cache add curl binutils && \
    strip -s /go/src/mac2vendor/mac2vendor && \
    curl -sL 'https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf;hb=HEAD' > /oui.txt && \
    grep "^00:00:00" /oui.txt && \
    grep -i "^00:00:0C" /oui.txt && \
    grep -i "^00:01:E6" /oui.txt && \
    gzip -9 /oui.txt

FROM scratch

COPY --from=build /go/src/mac2vendor/mac2vendor /
COPY --from=build /oui.txt.gz /

USER 65534

EXPOSE 3000

ENTRYPOINT ["/mac2vendor"]

CMD ["/oui.txt.gz"]
