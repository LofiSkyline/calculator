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

    // 先处理乘除，生成一个新的中间token数组
    var stack []string
    result, err := strconv.ParseFloat(tokens[0], 64)
    if err != nil {
        return 0, err
    }
    stack = append(stack, strconv.FormatFloat(result, 'f', -1, 64))

    i := 1
    for i < len(tokens) {
        op := tokens[i]
        num, err := strconv.ParseFloat(tokens[i+1], 64)
        if err != nil {
            return 0, err
        }

        switch op {
        case "*":
            prev, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
            prev = prev * num
            stack[len(stack)-1] = strconv.FormatFloat(prev, 'f', -1, 64)
        case "/":
            if num == 0 {
                return 0, errors.New("division by zero")
            }
            prev, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
            prev = prev / num
            stack[len(stack)-1] = strconv.FormatFloat(prev, 'f', -1, 64)
        case "+", "-":
            stack = append(stack, op)
            stack = append(stack, strconv.FormatFloat(num, 'f', -1, 64))
        default:
            return 0, errors.New("unsupported operator: " + op)
        }

        i += 2
    }

    // 再处理加减
    finalResult, err := strconv.ParseFloat(stack[0], 64)
    if err != nil {
        return 0, err
    }
    i = 1
    for i < len(stack) {
        op := stack[i]
        num, err := strconv.ParseFloat(stack[i+1], 64)
        if err != nil {
            return 0, err
        }

        switch op {
        case "+":
            finalResult += num
        case "-":
            finalResult -= num
        default:
            return 0, errors.New("unsupported operator in final evaluation: " + op)
        }

        i += 2
    }

    return finalResult, nil
}