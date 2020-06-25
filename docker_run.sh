docker pull registry.cn-hangzhou.aliyuncs.com/huayong/alarm:latest
docker stop  etl
docker container prune -f
docker run --name alarm -d \
-v /etc/localtime:/etc/localtime \
-v /opt/log:/opt/log \
-p 8081:8000 \
--restart=always \
--env CONSUL_HTTP_ADDR=127.0.0.1:8500 \
--env SERVER_ADDR=127.0.0.1 \
--env SERVER_PORT=8081 \
registry.cn-hangzhou.aliyuncs.com/huayong/alarm:latest
