#! /bin/bash

npm run build

docker build -t kizunaiix/my-web-demo-frontend:latest .

# 构建容器
# docker run -itd --name my-web-demo-frontend -p 3000:3000 kizunaiix/my-web-demo-frontend:latest
