package main

import (
    "log"
    "net/http"

    // "github.com/bufbuild/connect-go" // ConnectRPC框架
    // calculatorv1 "github.com/LofiSkyline/calculator/gen" // .pb.go结构
    calculatorconnect "github.com/LofiSkyline/calculator/gen/calculatorconnect" // .connect.go接口
    "github.com/LofiSkyline/calculator/internal/calculator" // 你的服务逻辑
)

func main() {
    // 创建 CalculatorServer 实例
    calculatorServer := &calculator.CalculatorServer{}

    // 用 ConnectRPC 框架注册服务
    mux := http.NewServeMux()

    path, handler := calculatorconnect.NewCalculatorServiceHandler(calculatorServer)
    mux.Handle(path, handler)

    // 启动 HTTP 服务器
    addr := ":8080"
    log.Printf("Server listening on %s ...", addr)
    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}

