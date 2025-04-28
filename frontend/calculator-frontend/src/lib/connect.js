import { createConnectTransport } from "@connectrpc/connect-web";

export const transport = createConnectTransport({
  baseUrl: "http://localhost:8080", // 你的后端地址
});
