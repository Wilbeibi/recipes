FROM alpine:3.7

RUN apk add --no-cache \
    tcpdump \
    bind-tools \
    netcat-openbsd \
    iputils \
    curl \
    drill \
    jq \
    vim 

RUN mv /usr/sbin/tcpdump /usr/bin/tcpdump
ENTRYPOINT [ "/bin/bash" ]