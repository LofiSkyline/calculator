import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { CalculatorService } from "../gen/calculator_connect";
import { ExpressionResponse } from "../gen/calculator_pb";
import { transport } from "../lib/connect";

const client = createPromiseClient(CalculatorService, transport);

export async function computeExpression(expression: string): Promise<number> {
  const response = await client.computeExpression({ expression }) as ExpressionResponse;
  return response.result;
}
