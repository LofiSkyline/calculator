package calculator

import (
    "context"
    "errors"
    "strconv"
    "strings"

    "github.com/bufbuild/connect-go"
    calculatorv1 "github.com/LofiSkyline/calculator/gen"
)

type CalculatorServer struct{}

// ComputeExpression 处理表达式计算请求
func (s *CalculatorServer) ComputeExpression(
    ctx context.Context,
    req *connect.Request[calculatorv1.ExpressionRequest],
) (*connect.Response[calculatorv1.ExpressionResponse], error) {
    expr := req.Msg.GetExpression()
    if expr == "" {
        return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("expression is empty"))
    }

    result, err := evaluate(expr)
    if err != nil {
        return nil, connect.NewError(connect.CodeInvalidArgument, err)
    }

    return connect.NewResponse(&calculatorv1.ExpressionResponse{
        Result: result,
    }), nil
}

// evaluate 简单解析表达式，支持 + - * /，不支持括号优先级
func evaluate(expression string) (float64, error) {
    tokens := strings.Fields(expression)
    if len(tokens)%2 == 0 {
        return 0, errors.New("invalid expression format")
    }

    var stack []string
    result, err := strconv.ParseFloat(tokens[0], 64)
    if err != nil {
        return 0, errors.New("invalid number: " + tokens[0])
    }
    stack = append(stack, strconv.FormatFloat(result, 'f', -1, 64))

    i := 1
    for i < len(tokens) {
        op := tokens[i]
        numStr := tokens[i+1]
        num, err := strconv.ParseFloat(numStr, 64)
        if err != nil {
            return 0, errors.New("invalid number: " + numStr)
        }

        switch op {
        case "*":
            prev, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
            prev *= num
            stack[len(stack)-1] = strconv.FormatFloat(prev, 'f', -1, 64)
        case "/":
            if num == 0 {
                return 0, errors.New("division by zero")
            }
            prev, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
            prev /= num
            stack[len(stack)-1] = strconv.FormatFloat(prev, 'f', -1, 64)
        case "+", "-":
            stack = append(stack, op)
            stack = append(stack, strconv.FormatFloat(num, 'f', -1, 64))
        default:
            return 0, errors.New("unsupported operator: " + op)
        }

        i += 2
    }

    finalResult, err := strconv.ParseFloat(stack[0], 64)
    if err != nil {
        return 0, errors.New("invalid calculation result")
    }

    i = 1
    for i < len(stack) {
        op := stack[i]
        num, err := strconv.ParseFloat(stack[i+1], 64)
        if err != nil {
            return 0, errors.New("invalid number in final evaluation: " + stack[i+1])
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
