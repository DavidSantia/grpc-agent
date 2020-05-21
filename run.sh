if [ -z "$(docker images -q fluentd-nr)" ]; then
    echo 'Please build fluend-nr image first'
    exit
fi

if [ -z "$(docker images -q grpc-agent)" ]; then
    echo 'Please build grpc-agentr image first'
    exit
fi

if [ -z "$(docker ps -a -q --filter name=fluentd)" ]; then
    echo 'Starting docker log forwarding'
    docker run -d --rm --name fluentd -p 24224:24224 fluentd-nr
fi

if [ -z "$(docker ps -a -q --filter name=grpc-agent)" ]; then
    echo 'Starting grpc-agent server'
    docker run -d --rm --name grpc-agent -p 9000:9000 --log-driver=fluentd --log-opt tag="nrlogs" grpc-agent
fi
