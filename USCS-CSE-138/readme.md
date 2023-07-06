# 测试
```bash
curl -X POST -H "Content-Type: application/json" -d '[{"id": 4, "name": "David", "age": 25}, {"id": 5, "name": "Emma", "age": 35}]' http://localhost:8081/post
```

# docker
```bash
cd /assigment-1

docker build -t myapp .
#使用以下命令运行 Docker 容器：
docker run -p 8081:8081 myapp
```