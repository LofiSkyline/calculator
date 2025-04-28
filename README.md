

# 🧮 简易计算器项目

本项目是一个基于 Go + ConnectRPC + Next.js 技术栈开发的全栈计算器应用。  
前端通过 Connect 协议调用后端接口，支持基础的加减乘除运算。

---

## 🛠️ 技术栈

- **后端：** Go + ConnectRPC
- **前端：** Next.js + TypeScript
- **通信协议：** ConnectRPC（非传统 HTTP/JSON）
- **单元测试：**
  - 后端：Go 自带测试框架（`testing`包）
  - 前端：Jest + React Testing Library

---

## 📐 项目架构

```
calculator/
├── backend/               # 后端Go服务
│   ├── cmd/server/         # 后端入口
│   ├── internal/calculator # 具体计算逻辑
│   ├── api/                # Proto文件
│   ├── gen/                # Proto生成的Go文件
├── calculator-frontend/    # 前端Next.js应用
│   ├── pages/              # 页面(index.tsx)
│   ├── src/
│   │   ├── service/        # ConnectRPC服务封装
│   │   ├── gen/            # Proto生成的前端TS文件
│   ├── jest.config.ts      # Jest测试配置
│   ├── babel.config.js     # Babel配置
```

---

## 🚀 如何启动项目

### 后端（Go）

1. 进入后端目录：

```bash
cd backend
```

2. 安装依赖（如果需要）：

```bash
go mod tidy
```

3. 启动后端服务：

```bash
go run cmd/server/main.go
```

默认监听端口为 `localhost:8080`，ConnectRPC协议。

---

### 前端（Next.js）

1. 进入前端目录：

```bash
cd calculator-frontend
```

2. 安装依赖：

```bash
npm install
```

3. 启动前端开发服务器：

```bash
npm run dev
```

默认运行在 `http://localhost:3000`。

---

## 🧪 运行单元测试

### 后端测试（Go）

```bash
cd backend
go test ./...
```

---

### 前端测试（Next.js）

```bash
cd calculator-frontend
npm run test
```

（已配置 Jest + Testing Library，支持 TSX 单测）

---

## 📦 附加说明

- 本项目 ConnectRPC通信基于 Connect 协议标准，客户端自动生成。
- 前端和后端均配备了基础单元测试。
- 项目符合远程调用、组件解耦、单元测试覆盖的标准开发要求。

---

