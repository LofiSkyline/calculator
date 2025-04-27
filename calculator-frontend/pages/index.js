import { useState } from 'react';

export default function Home() {
  const [expression, setExpression] = useState('');
  const [result, setResult] = useState(null);
  const [error, setError] = useState('');

  const handleAppendOperator = (operator) => {
    setExpression((prev) => prev + ' ' + operator + ' ');
  };

  const handleCalculate = async () => {
    if (!expression.trim()) {
      setError('请输入表达式');
      return;
    }

    try {
      const response = await fetch('http://localhost:8080/calculator.CalculatorService/ComputeExpression', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          expression: expression,
        }),
      });

      if (!response.ok) {
        throw new Error('服务器返回错误');
      }

      const data = await response.json();
      setResult(data.result);
      setError('');
    } catch (err) {
      console.error('请求出错:', err);
      setError('计算失败，请检查后端服务或输入格式');
    }
  };

  return (
    <div style={{ padding: '50px', fontFamily: 'Arial, sans-serif' }}>
      <h1>简单计算器</h1>
      <div style={{ marginBottom: '20px' }}>
        <input
          type="text"
          placeholder="请输入数字，点击符号或手动输入表达式"
          value={expression}
          onChange={(e) => setExpression(e.target.value)}
          style={{ width: '300px', padding: '8px', marginRight: '10px' }}
        />
        <button onClick={handleCalculate} style={{ padding: '8px 16px' }}>
          计算
        </button>
      </div>

      <div style={{ marginBottom: '20px' }}>
        <button onClick={() => handleAppendOperator('+')} style={{ marginRight: '10px', padding: '8px' }}>+</button>
        <button onClick={() => handleAppendOperator('-')} style={{ marginRight: '10px', padding: '8px' }}>-</button>
        <button onClick={() => handleAppendOperator('*')} style={{ marginRight: '10px', padding: '8px' }}>×</button>
        <button onClick={() => handleAppendOperator('/')} style={{ padding: '8px' }}>÷</button>
      </div>

      {result !== null && (
        <div style={{ marginTop: '20px', fontSize: '20px' }}>
          <strong>计算结果：</strong> {result}
        </div>
      )}

      {error && (
        <div style={{ marginTop: '20px', color: 'red' }}>
          {error}
        </div>
      )}
    </div>
  );
}
