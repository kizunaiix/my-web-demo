
# 使用官方 Node.js 镜像作为基础镜像
FROM node:lts-alpine3.20

# 设置工作目录
WORKDIR /app

# 复制项目编译后的文件到工作目录
COPY .next ./.next
COPY node_modules ./node_modules
COPY package.json ./package.json

# 安装依赖
# RUN npm install

# 暴露应用运行的端口，默认是 3000
EXPOSE 3000

# 启动 Next.js 应用
CMD ["npm", "start"]



######################### 旧版本：nginx版本 #########################
# 
# FROM nginx
# 
# # 添加作者信息
# LABEL maintainer="kizunaiix"
# 
# # 添加镜像描述
# LABEL description="based on next.js."
# 
# # 静态文件
# COPY my-next-app/out /app/reactbuild
# 
# # nginx配置文件
# COPY nginx/nginx.conf /etc/nginx/nginx.conf
# COPY nginx/default.conf /etc/nginx/conf.d/default.conf
# 
# # 等docker run的时候再写端口
###################################################################
