## Dockerfile tips
+ set timezone to UTC
    ```
    ENV TZ=UTC
    RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
    ```
+ `docker -p <dest_port>:<container_port>`

## Useful images
1. Docker management UI: 
```
docker run -d -p 9000:9000 \
--restart always \
-v /var/run/docker.sock:/var/run/docker.sock \
-v /opt/portainer:/data portainer/portainer
```

2. Code search:
```
docker run \
--publish 7080:7080 --rm \
--volume /tmp/sourcegraph/config:/etc/sourcegraph \
--volume /tmp/sourcegraph/data:/var/opt/sourcegraph \
sourcegraph/server
```
### Go/Docker Makefile template
+ https://github.com/thockin/go-build-template
