# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20.5

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到容器中的 /app 目录下
COPY . .

# 构建程序
RUN go build -o main .

# 暴露程序运行的端口
EXPOSE 8081

# 运行程序
CMD ["./main"]