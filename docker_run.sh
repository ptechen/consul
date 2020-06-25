docker pull $1
docker stop  etl
docker container prune -f
docker run --name $2 -d \
-v /etc/localtime:/etc/localtime \
-v /opt/log:/opt/log \
--restart=always \
--network=host \
--env CONSUL_HTTP_ADDR=127.0.0.1:8500 \
--env SERVER_ADDR=127.0.0.1 \
--env SERVER_PORT=8080 \
$1