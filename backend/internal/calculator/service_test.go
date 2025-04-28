package calculator

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    calculator "github.com/LofiSkyline/calculator/gen"
    "github.com/bufbuild/connect-go" 
)



// 测试 ComputeExpression 函数
func TestComputeExpression(t *testing.T) {
    server := &CalculatorServer{}

    tests := []struct {
        name       string
        expression string
        wantResult float64
        wantErr    bool
    }{
        {"Addition", "1 + 2", 3, false},
        {"Subtraction", "5 - 3", 2, false},
        {"Multiplication", "4 * 3", 12, false},
        {"Division", "10 / 2", 5, false},
        {"DivisionByZero", "10 / 0", 0, true},
        {"InvalidExpression", "abc + 1", 0, true},
        {"IncompleteExpression", "1 +", 0, true},
    }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := connect.NewRequest(&calculator.ExpressionRequest{
				Expression: tt.expression,
			})			
			resp, err := server.ComputeExpression(context.Background(), req)
			
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, resp.Msg.Result)
			}
		})
	}	
}

