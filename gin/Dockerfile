FROM alpine:3.14

RUN mkdir "/app"
WORKDIR "/app"

# mygindemo 在这里是APP的名字，也就是二进制文件的文件名
COPY ./out/mygindemo /app/app
# COPY resources /app/resources

ENTRYPOINT [ "./app" ]
