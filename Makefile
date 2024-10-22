.PHONY: run-containers run-proxy-server stop
run-containers:
    docker run --rm -d -p 9001:80 --name server1 kennethreitz/httpbin
    docker run --rm -d -p 9002:80 --name server2 kennethreitz/httpbin
    docker run --rm -d -p 9003:80 --name server3 kennethreitz/httpbin

stop:
 docker stop server1
 docker stop server2
 docker stop server3

run-proxy-server:
 go run cmd/main.go
