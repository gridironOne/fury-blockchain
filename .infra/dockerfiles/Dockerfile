## Starport

FROM --platform=linux/amd64 golang:latest AS starport
RUN curl https://get.starport.network/starport! | bash

##
## Fetch Golang Dependencies
##

FROM starport AS build-dependencies
WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod vendor

ONBUILD WORKDIR /app

##
## Fetch Golang Dependencies
##

FROM --platform=linux/amd64 build-dependencies AS build-tests
RUN go test -v ./...

FROM build-dependencies AS build-devel

COPY ./ .

ENTRYPOINT [ "starport", "chain", "serve" ]

# Build Golang files
FROM --platform=linux/amd64 build-dependencies AS build-binary
COPY ./ ./

ENV FURYD_HOME="/home/fury/"
RUN mkdir -p ${FURYD_HOME}/iid

RUN go mod tidy && go mod vendor
RUN make install

# RUN tar -xzf /dist/app/fury-blockchain_linux_amd64.tar.gz -C /dist/app/

# FROM ubuntu:latest AS setup-genisis
# COPY ./.infra/dockerfiles/setups/setup-genisis.sh /setup/setup-genisis.sh
# RUN /setup/setup-genisis.sh

# RUN curl https://raw.githubusercontent.com/furyfoundation/genesis/master/pandora-4/genesis.json

ENTRYPOINT ["fury", "start"]

##
## Build Distrubutiable Blockchain
##
FROM --platform=linux/amd64 ubuntu:latest as distributable

COPY --from=build-binary /go/bin/fury /app/fury
COPY --from=build-binary /go/pkg/mod/github.com/!cosm!wasm/wasmvm@v1.1.1/internal/api/ /go/pkg/mod/github.com/!cosm!wasm/wasmvm@v1.1.1/internal/api/
COPY ./.infra/dockerfiles/setups/setup-devnet.sh /root/setup-devnet.sh

ENV HOME=/root/
RUN rm -rf /root/.fury/ && chmod +x /root/setup-devnet.sh

ONBUILD LABEL world.fury.blockchain.image.authors="fury"

ONBUILD WORKDIR /app

ONBUILD EXPOSE 1317
ONBUILD EXPOSE 26657

# ONBUILD ENTRYPOINT ["/app/fury", "start", "--rpc.laddr", "tcp://0.0.0.0:26657"]

FROM distributable AS dist-devel
# RUN echo "Fetch genisis file"

FROM distributable AS dist-beta
# RUN echo "Fetch genisis file"

FROM distributable AS dist-rc
# RUN echo "Fetch genisis file"

FROM distributable AS dist-stable
# RUN echo "Fetch genisis file"
