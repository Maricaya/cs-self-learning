package main

import (
    // "fmt"
    "log"
    "net/http"
	"encoding/json"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    http.HandleFunc("/get", handleGet)
	http.HandleFunc("/post", handlePost)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World12221!"))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
    // 处理创建新用户的请求
    var user = []User{
        {ID: 1, Name: "Alice", Age: 20},
        {ID: 2, Name: "Bob", Age: 30},
        {ID: 3, Name: "Charlie", Age: 40},
    }

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // 保存用户到数据库或者其他存储介质
    // ...
    // 返回 JSON 响应
    json.NewEncoder(w).Encode(user)
}

// curl -X POST -H "Content-Type: application/json" -d '[{"id": 4, "name": "David", "age": 25}, {"id": 5, "name": "Emma", "age": 35}]' http://localhost:8081/post