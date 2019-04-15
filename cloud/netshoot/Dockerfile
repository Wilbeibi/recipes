FROM alpine:3.7

RUN set -ex \
    && apk update \
    && apk add --no-cache \
    bash \
    git \
    tcpdump \
    bind-tools \
    netcat-openbsd \
    iputils \
    curl \
    drill \
    jq \
    nmap \
    vim 

RUN apk add hping3 --update-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/testing

RUN mv /usr/sbin/tcpdump /usr/bin/tcpdump
ENTRYPOINT ["sh"]
