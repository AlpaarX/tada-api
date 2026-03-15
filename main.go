package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)


func main() {
  items := []map[string]interface{}{
    {
      "id": 1,
      "name": "自転車",
      "description": "2018年モデルのロードバイクです。軽量で高速な走行が可能です。",
      "price": 120000,
    },
    {
      "id": 2,
      "name": "ノートパソコン",
      "description": "高性能なノートパソコンで、仕事やゲームに最適です。",
      "price": 150000,
    },
    {
      "id": 3,
      "name": "スマートフォン",
      "description": "最新モデルのスマートフォンで、カメラ性能が優れています。",
      "price": 80000,
    },
  }

  r := gin.Default()

  r.GET("/items", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, items)
  })

  // Start server on port 8081
  // Server will listen on 0.0.0.0:8081 (localhost:8081 on Windows)
  r.Run(":8081")
}