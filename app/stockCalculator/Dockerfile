FROM golang:1.24-bookworm AS builder

WORKDIR /src

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    gcc \
    libc6-dev \
    libgl1-mesa-dev \
    libx11-dev \
    libxcursor-dev \
    libxinerama-dev \
    libxrandr-dev \
    libxxf86vm-dev \
    pkg-config \
  && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
COPY vendor ./vendor
COPY . .

ENV CGO_CFLAGS="-Wno-deprecated-declarations"
RUN go build -mod=vendor -o /out/stockcalculator-webapp ./webapp

FROM debian:bookworm-slim

RUN apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates \
  && update-ca-certificates \
  && rm -rf /var/lib/apt/lists/* \
  && useradd -r -u 10001 -g users -d /app stockcalc \
  && mkdir -p /app /data \
  && chown -R stockcalc:users /app /data

WORKDIR /app
COPY --from=builder /out/stockcalculator-webapp /app/stockcalculator-webapp

ENV STOCKCALC_DATA_DIR=/data
ENV PORT=8080
ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

EXPOSE 8080
VOLUME ["/data"]

USER stockcalc
ENTRYPOINT ["/app/stockcalculator-webapp"]
