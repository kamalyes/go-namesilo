/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:10:15
 * @FilePath: \go-namesilo\client\mock.go
 * @Description: Mock 客户端 - 用于测试
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"fmt"
)

// MockClient Mock 客户端，用于测试
type MockClient struct {
	baseClient        *Client
	MockDoRequestFunc func(ctx context.Context, operation string, params map[string]string) ([]byte, error)
}

// DoRequest Mock 的 DoRequest 方法
func (m *MockClient) DoRequest(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
	if m.MockDoRequestFunc != nil {
		return m.MockDoRequestFunc(ctx, operation, params)
	}
	return m.baseClient.DoRequest(ctx, operation, params)
}

// DoPublicRequest Mock 的 DoPublicRequest 方法
func (m *MockClient) DoPublicRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	// Mock 实现可以在测试中通过设置自定义函数来覆盖
	return m.baseClient.DoPublicRequest(ctx, method, path, body, result)
}

// ParseResponse 解析响应（直接调用内部 client 的方法）
func (m *MockClient) ParseResponse(data []byte, v interface{}) error {
	return m.baseClient.ParseResponse(data, v)
}

// AsClient 返回一个包装后的 *Client，可以用于 Service
// 这个方法通过替换内部的 DoRequest 来实现 Mock
func (m *MockClient) AsClient() *Client {
	// 注意：Go 语言的限制导致无法在运行时替换结构体方法
	// 因此这个方法只是返回基础 Client，实际的 Mock 功能需要在测试中使用适配器模式
	return m.baseClient
}

// NewMockClient 创建一个 Mock 客户端 (XML 格式)
func NewMockClient() *MockClient {
	client, _ := New("mock-api-key", WithDebug(false), WithResponseType(ResponseTypeXML))
	return &MockClient{
		baseClient: client,
	}
}

// NewMockJSONClient 创建一个 Mock 客户端 (JSON 格式)
func NewMockJSONClient() *MockClient {
	client, _ := New("mock-api-key", WithDebug(false), WithResponseType(ResponseTypeJSON))
	return &MockClient{
		baseClient: client,
	}
}

// MockResponseFunc 动态生成 Mock 响应的函数类型
type MockResponseFunc func(operation string, params map[string]string) ([]byte, error)

// WithMockResponse 设置 Mock 响应函数
func (m *MockClient) WithMockResponse(fn MockResponseFunc) *MockClient {
	m.MockDoRequestFunc = func(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
		return fn(operation, params)
	}
	return m
}

// MockSuccessXMLResponse 动态生成成功的 XML 响应
func MockSuccessXMLResponse(operation string, content string) MockResponseFunc {
	return func(op string, params map[string]string) ([]byte, error) {
		if op != operation {
			return nil, fmt.Errorf("unexpected operation: %s, expected: %s", op, operation)
		}
		response := fmt.Sprintf(`<namesilo>
<request>
<operation>%s</operation>
<ip>185.220.236.2</ip>
</request>
<reply>
<code>300</code>
<detail>success</detail>
%s
</reply>
</namesilo>`, operation, content)
		return []byte(response), nil
	}
}

// MockErrorResponse 动态生成错误响应
func MockErrorResponse(code int, detail string) MockResponseFunc {
	return func(operation string, params map[string]string) ([]byte, error) {
		response := fmt.Sprintf(`<namesilo>
<request>
<operation>%s</operation>
<ip>185.220.236.2</ip>
</request>
<reply>
<code>%d</code>
<detail>%s</detail>
</reply>
</namesilo>`, operation, code, detail)
		return []byte(response), nil
	}
}

// MockJSONResponse 动态生成 JSON 响应
func MockJSONResponse(operation string, jsonContent string) MockResponseFunc {
	return func(op string, params map[string]string) ([]byte, error) {
		if op != operation {
			return nil, fmt.Errorf("unexpected operation: %s, expected: %s", op, operation)
		}
		response := fmt.Sprintf(`{
	"request": {
		"operation": "%s",
		"ip": "185.220.236.2"
	},
	"reply": {
		"code": 300,
		"detail": "success",
		%s
	}
}`, operation, jsonContent)
		return []byte(response), nil
	}
}

// MockConditionalResponse 根据参数条件动态生成响应
func MockConditionalResponse(conditions map[string]MockResponseFunc, defaultFn MockResponseFunc) MockResponseFunc {
	return func(operation string, params map[string]string) ([]byte, error) {
		// 根据 operation 选择对应的响应函数
		if fn, ok := conditions[operation]; ok {
			return fn(operation, params)
		}
		// 使用默认响应
		if defaultFn != nil {
			return defaultFn(operation, params)
		}
		return nil, fmt.Errorf("no mock response for operation: %s", operation)
	}
}
