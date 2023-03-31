FROM golang:1.19.2-alpine

#RUN apk update && apk install wget curl gcc make rsync libreadline-dev
RUN apk update && apk add gcc
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update && apk add --no-cache curl bash tree tzdata binutils bash-completion curl openssl zstd\
    && rm -rf /var/cache/apk/* \
    && sed -i '1s?/bin/ash?/bin/bash?g' /etc/passwd\
    && sed -i 4d /etc/shells \
    && sed -i '2i/bin/bash' /etc/shells\
    && wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub \
    && wget -O /tmp/glibc-2.31-r0.apk https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.31-r0/glibc-2.31-r0.apk \
    && wget -O /tmp/glibc-bin-2.31-r0.apk https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.31-r0/glibc-bin-2.31-r0.apk \
    && wget -O /tmp/glibc-i18n-2.31-r0.apk https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.31-r0/glibc-i18n-2.31-r0.apk \
    && cd /tmp && apk add glibc-2.31-r0.apk  glibc-i18n-2.31-r0.apk glibc-bin-2.31-r0.apk \
    && wget -O /tmp/gcc-libs.tar.zst http://mirrors.ustc.edu.cn/archlinux/pool/packages/gcc-libs-11.1.0-3-x86_64.pkg.tar.zst \
    && /usr/bin/unzstd /tmp/gcc-libs.tar.zst \
    && mkdir /tmp/gcc \
    && tar -xf /tmp/gcc-libs.tar -C /tmp/gcc \
    && mv /tmp/gcc/usr/lib/libgcc* /tmp/gcc/usr/lib/libstdc++* /usr/glibc-compat/lib \
    && strip /usr/glibc-compat/lib/libgcc_s.so.* /usr/glibc-compat/lib/libstdc++.so* \
    && wget http://mirrors.ustc.edu.cn/archlinux/pool/packages/zlib-1%3A1.2.11-4-x86_64.pkg.tar.xz -O /tmp/libz.tar.xz \
    && mkdir /tmp/libz \
    && tar -xf /tmp/libz.tar.xz -C /tmp/libz \
    && mv /tmp/libz/usr/lib/libz.so* /usr/glibc-compat/lib \
    && apk del binutils \
    && rm -rf /tmp/* /var/cache/apk/* \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo 'Asia/Shanghai' > /etc/timezone
#RUN echo "short_open_tag = On" >> /etc/php/7.0/cli/php.ini

# docker
#RUN curl -O https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz && tar zxvf docker-latest.tgz && cp docker/docker /usr/local/bin/ && rm -rf docker docker-latest.tgz

RUN go env

WORKDIR /root/source
