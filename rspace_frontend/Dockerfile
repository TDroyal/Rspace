# 阶段一：构建前端
FROM node:14 as build-stage
WORKDIR /app
COPY package*.json ./

# 配置 npm 镜像源
RUN npm config set registry https://registry.npmmirror.com

RUN npm install
COPY . .
RUN npm run build


# 阶段二：运行 Nginx 服务
FROM nginx:latest as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
# COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 外部访问容器中的nignx服务通过80端口
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
