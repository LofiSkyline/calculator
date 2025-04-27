package calculator

import (
    "context"
    "errors"
    "strconv"
    "strings"

    "github.com/bufbuild/connect-go" // ✅ 这里确实要导入
    calculatorv1 "github.com/LofiSkyline/calculator/gen"
)

type CalculatorServer struct{}

func (s *CalculatorServer) ComputeExpression(ctx context.Context, req *connect.Request[calculatorv1.ExpressionRequest]) (*connect.Response[calculatorv1.ExpressionResponse], error) {
    expr := req.Msg.GetExpression()
    if expr == "" {
        return nil, errors.New("expression is empty")
    }

    result, err := evaluate(expr)
    if err != nil {
        return nil, err
    }

    return connect.NewResponse(&calculatorv1.ExpressionResponse{
        Result: result,
    }), nil
}

// evaluate 函数不变
func evaluate(expression string) (float64, error) {
    tokens := strings.Fields(expression)
    if len(tokens)%2 == 0 {
        return 0, errors.New("invalid expression format")
    }

    result, err := strconv.ParseFloat(tokens[0], 64)
    if err != nil {
        return 0, err
    }

    for i := 1; i < len(tokens); i += 2 {
        op := tokens[i]
        nextVal, err := strconv.ParseFloat(tokens[i+1], 64)
        if err != nil {
            return 0, err
        }

        switch op {
        case "+":
            result += nextVal
        case "-":
            result -= nextVal
        case "*":
            result *= nextVal
        case "/":
            if nextVal == 0 {
                return 0, errors.New("division by zero")
            }
            result /= nextVal
        default:
            return 0, errors.New("unsupported operator: " + op)
        }
    }

    return result, nil
}
