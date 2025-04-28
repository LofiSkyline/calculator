package main

import (
    "log"
    "net/http"
    "strings"

    "github.com/LofiSkyline/calculator/internal/calculator"
    calculatorconnect "github.com/LofiSkyline/calculator/gen/calculatorconnect"
)

func main() {
    // 实例化服务
    calculatorServer := &calculator.CalculatorServer{}

    // 创建 mux
    mux := http.NewServeMux()

    // 注册 CalculatorService
    path, handler := calculatorconnect.NewCalculatorServiceHandler(calculatorServer)
    mux.Handle(path, handler)

    // CORS + Connect 检查
    corsMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Connect-Protocol-Version")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        // 强制检查 Content-Type 和 Connect-Protocol-Version 使用ConnectRPC协议通信
        contentType := r.Header.Get("Content-Type")
        connectVersion := r.Header.Get("Connect-Protocol-Version")

        if !(strings.HasPrefix(contentType, "application/json") || strings.HasPrefix(contentType, "application/proto+connect")) {
            http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
            return
        }

        if connectVersion != "1" {
            http.Error(w, "Missing or invalid Connect-Protocol-Version header", http.StatusBadRequest)
            return
        }

        mux.ServeHTTP(w, r)
    })

    addr := ":8080"
    log.Printf("✅ Server listening on %s ...", addr)

    if err := http.ListenAndServe(addr, corsMux); err != nil {
        log.Fatalf("❌ failed to start server: %v", err)
    }
}
