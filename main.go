package main

import (
  "net/http"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)


func main() {
  items := []map[string]interface{}{
    {
      "id": 1,
      "title": "自転車",
      "description": "2018年モデルのロードバイクです。軽量で高速な走行が可能です。",
      "price": 120000,
    },
    {
      "id": 2,
      "title": "ノートパソコン",
      "description": "高性能なノートパソコンで、仕事やゲームに最適です。",
      "price": 150000,
    },
    {
      "id": 3,
      "title": "スマートフォン",
      "description": "最新モデルのスマートフォンで、カメラ性能が優れています。",
      "price": 80000,
    },
  }

  r := gin.Default()
  r.Use(cors.New(cors.Config{
  AllowOrigins: []string{"http://localhost:5173"},
  AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
  AllowHeaders: []string{"Origin", "Content-Type"},
  }))

  r.GET("/items", func(c *gin.Context) {
    c.JSON(http.StatusOK, items)
  })

  r.Run(":8081")
}