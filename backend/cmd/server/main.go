package main

import (
    "log"
    "net/http"
    "github.com/LofiSkyline/calculator/internal/calculator"
    calculatorconnect "github.com/LofiSkyline/calculator/gen/calculatorconnect"
    //"github.com/bufbuild/connect-go"
)

func main() {
    calculatorServer := &calculator.CalculatorServer{}

    mux := http.NewServeMux()

    path, handler := calculatorconnect.NewCalculatorServiceHandler(calculatorServer)
    mux.Handle(path, handler)

    // 🌟 加一个跨域处理器（重点）
    corsMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        mux.ServeHTTP(w, r)
    })

    addr := ":8080"
    log.Printf("Server listening on %s ...", addr)
    if err := http.ListenAndServe(addr, corsMux); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}

