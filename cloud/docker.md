## Tips and tricks
+ set timezone to UTC in Dockerfile
    ```
    ENV TZ=UTC
    RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
    ```
+ `docker -p <dest_port>:<container_port>`
+ `docker images --digests` show docker image sha256
+ dock post-installation
    - `sudo usermod -aG docker $USER`
    - In case docker warning config.json [permission denied](https://askubuntu.com/questions/747778/docker-warning-config-json-permission-denied):
        - `sudo chown "$USER":"$USER" /home/"$USER"/.docker -R`
        - `sudo chmod g+rwx "/home/$USER/.docker" -R`

## docker inspect with jq
+ `docker inspect ebdb795dc32d | jq '.[0]' | jq keys`: show all keys
+ `docker inspect ebdb795dc32d | jq '.[0]' | jq .NetworkSettings`: show values for
    a specific key (or ` jq -r '.[0].NetworkSettings.IPAddress'`)


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

3. [FreshRSS](https://github.com/FreshRSS/FreshRSS/tree/master/Docker#run-freshrss):
```
# You can optionally run from the directory containing the FreshRSS source code:
cd ./FreshRSS/

# The data will be saved on the host in `./data/`
mkdir -p ./data/

sudo docker run -d --restart unless-stopped --log-opt max-size=10m \
  -v $(pwd)/data:/var/www/FreshRSS/data \
  -e 'CRON_MIN=5,35' \
  -p 8080:80 \
  --name freshrss freshrss/freshrss
```
### Go/Docker Makefile template
+ https://github.com/thockin/go-build-template
