import { useState } from "react";
import { computeExpression } from "../src/service/calculator";

export default function Home() {
  const [expression, setExpression] = useState("");
  const [result, setResult] = useState<number | null>(null);
  const [loading, setLoading] = useState(false);

  const handleCalculate = async () => {
    if (expression.trim() === "") {
      alert("请输入要计算的表达式！");
      return;
    }

    try {
      setLoading(true);
      const res = await computeExpression(expression);
      setResult(res);
    } catch (error) {
      console.error("计算出错:", error);
      alert("计算失败，请检查服务器连接！");
    } finally {
      setLoading(false);
    }
  };

  const handleInsertOperator = (operator: string) => {
    setExpression((prev) => prev + " " + operator + " ");
  };

  return (
    <div style={{ padding: "2rem", fontFamily: "Arial, sans-serif" }}>
      <h1>🧮 简易计算器</h1>

      <div style={{ marginBottom: "1rem" }}>
        <input
          type="text"
          value={expression}
          onChange={(e) => setExpression(e.target.value)}
          placeholder="请输入表达式，例如 5 + 7 * 2"
          style={{ width: "300px", marginRight: "10px", padding: "8px" }}
        />

        <button onClick={handleCalculate} disabled={loading}>
          {loading ? "计算中..." : "计算"}
        </button>
      </div>

      {/* 新增运算符按钮区域 */}
      <div style={{ marginBottom: "1rem" }}>
        {["+", "-", "*", "/"].map((op) => (
          <button
            key={op}
            onClick={() => handleInsertOperator(op)}
            style={{
              marginRight: "10px",
              padding: "8px 12px",
              fontSize: "16px",
            }}
          >
            {op}
          </button>
        ))}
      </div>

      <div style={{ marginTop: "1rem" }}>
        <input
          type="text"
          value={result !== null ? result : ""}
          readOnly
          placeholder="计算结果"
          style={{ width: "300px", padding: "8px" }}
        />
      </div>
    </div>
  );
}
