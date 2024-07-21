# GinDemo_front

## react

```bash
npx create-next-app {project-name} # 脚手架 -> 之后要把next.config.mjs里改成const nextConfig = { output: 'export', };
npm run dev # 在3000端口启动调试
npm run build #编译，输出会在一个新建的./out 文件夹里面
```

## docker

```bash
# 制作镜像：  
docker build -t kizunaiix/gindemo-nginx:v0.1 .  # -t是打tag ，注意最后有个"." ，指使用当前目录

# (调试)运行这个命令:  ->  调试的时候可以用 -p 8080:80, 这样用的就是宿主机的8080端口了

npm start
    # docker run -itd --name gindemo-nginx -v /home/yaodong/codings/GinDemo_front/reactapp/build:/usr/share/nginx:ro -p 8080:80 kizunaiix/gindemo-nginx:v0.1  没必要用这个命令，用npm就行了

# (部署)运行这个命令：  
docker run -itd --name gindemo_nginx -p 80:80 -d kizunaiix/nginx_gindemo_fe:v1.2  
```
