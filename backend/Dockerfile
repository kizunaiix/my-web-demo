FROM alpine:3.14

RUN mkdir "/appHome"
WORKDIR "/appHome"

# mygindemo 在这里是APP的名字，也就是二进制文件的文件名
COPY ./bin/gin-app-be /appHome/app
# COPY resources /app/resources

ENTRYPOINT [ "./app" ]
