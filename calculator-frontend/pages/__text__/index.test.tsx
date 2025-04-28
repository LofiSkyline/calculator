import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import IndexPage from "../index"; 
import { computeExpression } from "@/service/calculator";
import { ConnectError, Code } from "@connectrpc/connect"; 

jest.mock("@/service/calculator"); // 直接mock整个 computeExpression模块

describe("Calculator Page", () => {
  beforeEach(() => {
    jest.clearAllMocks(); // 每次测试前清除mock，避免互相污染
  });

  test("renders calculator page correctly", () => {
    render(<IndexPage />);

    // 检查是否有输入框、运算按钮、结果显示区域
    expect(screen.getByPlaceholderText(/请输入表达式/i)).toBeInTheDocument();
    expect(screen.getByText("+")).toBeInTheDocument();
    expect(screen.getByText("-")).toBeInTheDocument();
    expect(screen.getByText("*")).toBeInTheDocument();
    expect(screen.getByText("/")).toBeInTheDocument();
  });

  test("computes expression successfully", async () => {
    (computeExpression as jest.Mock).mockResolvedValue(3); // Mock返回3
  
    render(<IndexPage />);
  
    const input = screen.getByPlaceholderText(/请输入表达式/i);
    const calculateButton = screen.getByText("计算"); 
  
    fireEvent.change(input, { target: { value: "1 + 2" } });
    fireEvent.click(calculateButton);
  
    await waitFor(() => {
      expect(screen.getByPlaceholderText("计算结果")).toHaveValue("3"); 
    });
  });
  

  test("shows error message on computation failure", async () => {
    (computeExpression as jest.Mock).mockRejectedValue(
      new ConnectError("invalid expression format", Code.InvalidArgument)
    );
  
    jest.spyOn(window, "alert").mockImplementation(() => {});
  
    render(<IndexPage />);
  
    const input = screen.getByPlaceholderText(/请输入表达式/i);
    const calculateButton = screen.getByText("计算");
  
    fireEvent.change(input, { target: { value: "1 +" } });
    fireEvent.click(calculateButton);
  
    await waitFor(() => {
      expect(window.alert).toHaveBeenCalledWith(expect.stringContaining("invalid expression format"));
    });
  });
  
});
