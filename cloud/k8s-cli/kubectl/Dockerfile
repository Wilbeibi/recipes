FROM alpine:3.8
ENV KUBE_VERSION="v1.10.6"
ADD https://storage.googleapis.com/kubernetes-release/release/${KUBE_VERSION}/bin/linux/amd64/kubectl /usr/local/bin/kubectl

RUN apk add --no-cache curl ca-certificates \
    && chmod +x /usr/local/bin/kubectl \
    && rm -rf /var/cache/apk/*
ENTRYPOINT ["kubectl"]