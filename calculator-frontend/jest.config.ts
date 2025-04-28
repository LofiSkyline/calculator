import type { Config } from 'jest';

const config: Config = {
  testEnvironment: "jsdom",
  setupFilesAfterEnv: ["<rootDir>/jest.setup.ts"],
  moduleNameMapper: {
    "^@/(.*)$": "<rootDir>/src/$1",
    "\\.(css|less|sass|scss)$": "identity-obj-proxy" // mock掉CSS模块
  },
  transform: {
    "^.+\\.(ts|tsx|js|jsx)$": "babel-jest", // 👈 用 babel-jest 处理所有文件！
  },
  moduleFileExtensions: ["ts", "tsx", "js", "jsx"],
  testMatch: ["**/__tests__/**/*.test.ts?(x)", "**/?(*.)+(test).ts?(x)"],
};

export default config;
