# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

FROM alpine as alpine
RUN apk --no-cache add ca-certificates

FROM  golang:1.19 AS builder
ADD . /src
WORKDIR /src
RUN cd /src && echo $(ls -1 /src)
RUN go mod download
RUN go build -o /inclusion-prover .

# final stage
FROM debian:stable-slim
COPY --from=builder /inclusion-prover ./
RUN chmod +x ./inclusion-prover
RUN mkdir -p /mount
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
LABEL org.opencontainers.image.source https://github.com/sygmaprotocol/sygma-inclusion-prover
ENTRYPOINT ["./inclusion-prover"]
