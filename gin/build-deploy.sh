#! /bin/bash

go build -o ./out/mygindemo ./cmd/*
docker build -t kizunaiix/my-web-demo-backend:latest .

# 构建容器
# docker run -itd --name my-web-demo-backend -p 9000:9000 kizunaiix/my-web-demo-backend:latest
