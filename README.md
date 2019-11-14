# mac2vendor

[![Build Status](https://drone.srv.kojedz.in/api/badges/krichy/mac2vendor/status.svg)](https://drone.srv.kojedz.in/krichy/mac2vendor)

Is a small service offered via http to look up mac addresses to vendors. It uses Wireshark's oui database,
reading the whole content into memory and serving requests over http, replying with a single text content.

## Usage

```bash
$ docker run -it --rm -p 3000:3000 rkojedzinszky/mac2vendor
```

To access:

```bash
$ curl http://127.0.0.1:3000/F0:41:C8:A0:00:00
IeeeRegi
$ curl http://127.0.0.1:3000/00:02:4B
Cisco
```

To request json output, you can specify it in the request:
```bash
$ curl http://127.0.0.1:3000/F0:41:C8:A0:00:00?format=json
{"prefix":"F0:41:C8","vendor":"IeeeRegi","comments":"IEEE Registration Authority"}
```

# Service

A free live service is available [here](https://mac2vendor.srv.kojedz.in).