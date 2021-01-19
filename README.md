# Envoy grpc_json_transcoder example project 

1. Build proto (Optional, I have done this for you but you have to rebuild it if you make changes to the protofile)

    ```
    sudo pip install grpcio requests
    cd service
    ./script/gen
    ./script/gen_kv.pb
    cd -
    ```

2. Build go service (Optional, I have built this for you, feel free to rebuild it if security is a concern for you)

    ```
    script/bootstrap
    ./script/build
    ```

3. Startup docker-compose

    ```
    docker-compose up --build
    ```

4. Do request

    ```
    docker-compose exec python /client/client.py set foo bar
    docker-compose exec python /client/client.py get foo
    docker-compose exec python /client/client.py count
    ```

### Explanation

1. Python client sends JSON request to Envoy proxy at port 9001. 
2. Envoy proxy convert the JSON request using the provided proto and the grpc_json_transcoder filter. 
3. The grpc requests are then forwarded to your microservice cluster which is running a grpc server. 


### TODO

* Build proto in grpc container instead.
* Guide on deploying to GCP

### Caveats

* Using envoy 1.17 as the docker base image. envoyproxy/envoy-dev:latest is recommended. 

## References

* https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/grpc_json_transcoder/v3/transcoder.proto#envoy-v3-api-msg-extensions-filters-http-grpc-json-transcoder-v3-grpcjsontranscoder
