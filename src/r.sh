#! /bin/bash
podman stop godoku; podman rm godoku
podman rmi godoku:0.1
podman images | grep "none"| awk '{print $3}' | xargs podman rmi -f
podman build -t godoku:0.1 ./
podman run -d --name godoku -p 8080:8080 godoku:0.1
podman logs -f godoku
