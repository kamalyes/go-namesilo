/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 23:55:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:15:16
 * @FilePath: \go-namesilo\client\mock_test.go
 * @Description: Mock 测试客户端 - 动态生成响应
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockClient Mock 客户端，用于测试
type MockClient struct {
	*Client
	MockDoRequestFunc func(ctx context.Context, operation string, params map[string]string) ([]byte, error)
}

// DoRequest Mock 的 DoRequest 方法
func (m *MockClient) DoRequest(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
	if m.MockDoRequestFunc != nil {
		return m.MockDoRequestFunc(ctx, operation, params)
	}
	return m.Client.DoRequest(ctx, operation, params)
}

// NewMockClient 创建一个 Mock 客户端 (XML 格式)
func NewMockClient() *MockClient {
	client, _ := New("mock-api-key", WithDebug(false), WithResponseType(ResponseTypeXML))
	return &MockClient{
		Client: client,
	}
}

// NewMockJSONClient 创建一个 Mock 客户端 (JSON 格式)
func NewMockJSONClient() *MockClient {
	client, _ := New("mock-api-key", WithDebug(false), WithResponseType(ResponseTypeJSON))
	return &MockClient{
		Client: client,
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

// TestMockDynamicResponse 测试动态生成响应
func TestMockDynamicResponse(t *testing.T) {
	mockClient := NewMockClient()

	// 动态生成域名列表内容
	domainsContent := `<domains>
<domain created="2025-08-01" expires="2026-08-01">example1.com</domain>
<domain created="2025-02-27" expires="2027-02-27">example2.com</domain>
<domain created="2026-01-09" expires="2027-01-09">example3.com</domain>
</domains>`

	// 设置动态 Mock 响应
	mockClient.WithMockResponse(MockSuccessXMLResponse("listDomains", domainsContent))

	ctx := context.Background()
	params := map[string]string{
		"page":      "1",
		"page_size": "10",
	}

	resp, err := mockClient.DoRequest(ctx, "listDomains", params)

	assert.NoError(t, err, "请求应该成功")
	assert.Contains(t, string(resp), "example1.com", "响应应该包含 example1.com")
	assert.Contains(t, string(resp), "<code>300</code>", "响应应该包含成功代码")

	t.Logf("动态生成的响应:\n%s", string(resp))
}

// TestMockErrorResponse 测试错误响应
func TestMockErrorResponse(t *testing.T) {
	mockClient := NewMockClient()

	// 设置错误响应
	mockClient.WithMockResponse(MockErrorResponse(110, "Invalid API Key"))

	ctx := context.Background()
	resp, err := mockClient.DoRequest(ctx, "listDomains", nil)

	assert.NoError(t, err, "DoRequest 不应该返回错误")
	assert.Contains(t, string(resp), "<code>110</code>", "响应应该包含错误代码")
	assert.Contains(t, string(resp), "Invalid API Key", "响应应该包含错误详情")

	t.Logf("错误响应:\n%s", string(resp))
}

// TestMockJSONResponse 测试 JSON 响应
func TestMockJSONResponse(t *testing.T) {
	mockClient := NewMockJSONClient()

	// 动态生成 JSON 内容
	jsonContent := `"domains": ["example1.com", "example2.com", "example3.com"]`

	mockClient.WithMockResponse(MockJSONResponse("listDomains", jsonContent))

	ctx := context.Background()
	resp, err := mockClient.DoRequest(ctx, "listDomains", nil)

	assert.NoError(t, err, "请求应该成功")
	assert.Contains(t, string(resp), "example1.com", "响应应该包含域名")
	assert.Contains(t, string(resp), `"code": 300`, "响应应该包含成功代码")

	t.Logf("JSON 响应:\n%s", string(resp))
}

// TestMockConditionalResponse 测试条件响应
func TestMockConditionalResponse(t *testing.T) {
	mockClient := NewMockClient()

	// 为不同操作设置不同的响应
	conditions := map[string]MockResponseFunc{
		"listDomains": MockSuccessXMLResponse("listDomains", `<domains>
<domain created="2025-08-01" expires="2026-08-01">test1.com</domain>
</domains>`),
		"getDomainInfo": MockSuccessXMLResponse("getDomainInfo", `<contact>
<contact_id>123456</contact_id>
</contact>`),
	}

	mockClient.WithMockResponse(MockConditionalResponse(conditions, MockErrorResponse(500, "Unknown operation")))

	ctx := context.Background()

	// 测试 listDomains
	resp1, err := mockClient.DoRequest(ctx, "listDomains", nil)
	assert.NoError(t, err)
	assert.Contains(t, string(resp1), "test1.com")
	t.Logf("listDomains 响应:\n%s", string(resp1))

	// 测试 getDomainInfo
	resp2, err := mockClient.DoRequest(ctx, "getDomainInfo", nil)
	assert.NoError(t, err)
	assert.Contains(t, string(resp2), "123456")
	t.Logf("getDomainInfo 响应:\n%s", string(resp2))

	// 测试未定义的操作
	resp3, err := mockClient.DoRequest(ctx, "unknownOperation", nil)
	assert.NoError(t, err)
	assert.Contains(t, string(resp3), "Unknown operation")
	t.Logf("unknownOperation 响应:\n%s", string(resp3))
}

// TestMockWithRealData 测试使用真实数据
func TestMockWithRealData(t *testing.T) {
	mockClient := NewMockClient()

	// 使用示例域名数据
	realDomainsContent := `<domains>
<domain created="2025-11-24" expires="2026-11-24">example1.com</domain>
<domain created="2026-01-07" expires="2027-01-07">example2.com</domain>
<domain created="2025-09-17" expires="2026-09-18">example3.com</domain>
</domains>`

	mockClient.WithMockResponse(MockSuccessXMLResponse("listDomains", realDomainsContent))

	ctx := context.Background()
	resp, err := mockClient.DoRequest(ctx, "listDomains", nil)

	assert.NoError(t, err)

	// 解析并验证
	type Domain struct {
		Name    string `xml:",chardata"`
		Created string `xml:"created,attr"`
		Expires string `xml:"expires,attr"`
	}

	type Result struct {
		Code    int      `xml:"reply>code"`
		Detail  string   `xml:"reply>detail"`
		Domains []Domain `xml:"reply>domains>domain"`
	}

	var result Result
	err = mockClient.ParseResponse(resp, &result)
	assert.NoError(t, err)
	assert.Equal(t, 300, result.Code)
	assert.Equal(t, 3, len(result.Domains))

	t.Logf("解析到 %d 个域名", len(result.Domains))
	for _, domain := range result.Domains {
		t.Logf("  - %s (创建: %s, 过期: %s)", domain.Name, domain.Created, domain.Expires)
	}
}

// TestMockJSONRealData 测试使用真实 JSON 数据
func TestMockJSONRealData(t *testing.T) {
	mockClient := NewMockJSONClient()

	// 示例 JSON 响应数据（精简版）
	realJSONResponse := `{
    "request": {
        "operation": "listDomains",
        "ip": "192.0.2.1"
    },
    "reply": {
        "code": 300,
        "detail": "success",
        "domains": [
            {"domain": "example1.com", "created": "2025-08-01", "expires": "2026-08-01"},
            {"domain": "example2.com", "created": "2025-02-27", "expires": "2027-02-27"},
            {"domain": "example3.com", "created": "2026-01-09", "expires": "2027-01-09"}
        ]
    }
}`

	mockClient.MockDoRequestFunc = func(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
		return []byte(realJSONResponse), nil
	}

	ctx := context.Background()
	resp, err := mockClient.DoRequest(ctx, "listDomains", nil)

	assert.NoError(t, err)
	assert.Contains(t, string(resp), "example1.com")
	assert.Contains(t, string(resp), `"code": 300`)

	// 解析 JSON
	type JSONDomain struct {
		Domain  string `json:"domain"`
		Created string `json:"created"`
		Expires string `json:"expires"`
	}

	type JSONResult struct {
		Reply struct {
			Code    int          `json:"code"`
			Detail  string       `json:"detail"`
			Domains []JSONDomain `json:"domains"`
		} `json:"reply"`
	}

	var result JSONResult
	err = mockClient.ParseResponse(resp, &result)
	assert.NoError(t, err)
	assert.Equal(t, 300, result.Reply.Code)
	assert.Equal(t, "success", result.Reply.Detail)
	assert.Equal(t, 3, len(result.Reply.Domains))

	// 验证几个关键域名
	assert.Equal(t, "example1.com", result.Reply.Domains[0].Domain)
	assert.Equal(t, "2025-08-01", result.Reply.Domains[0].Created)
	assert.Equal(t, "example2.com", result.Reply.Domains[1].Domain)

	t.Logf("JSON 格式解析到 %d 个域名", len(result.Reply.Domains))
}

// TestMockJSONFullRealData 测试使用完整真实 JSON 数据(所有域名)
func TestMockJSONFullRealData(t *testing.T) {
	mockClient := NewMockJSONClient()

	// 完整的示例 JSON 响应（精简版）
	fullJSONResponse := `{
    "request": {
        "operation": "listDomains",
        "ip": "192.0.2.1"
    },
    "reply": {
        "code": 300,
        "detail": "success",
        "domains": [
            {"domain": "example1.com", "created": "2025-08-01", "expires": "2026-08-01"},
            {"domain": "example2.com", "created": "2025-02-27", "expires": "2027-02-27"},
            {"domain": "example3.com", "created": "2026-01-13", "expires": "2027-01-13"},
            {"domain": "example4.com", "created": "2026-01-14", "expires": "2027-01-14"}
        ]
    }
}`

	mockClient.MockDoRequestFunc = func(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
		return []byte(fullJSONResponse), nil
	}

	ctx := context.Background()
	resp, err := mockClient.DoRequest(ctx, "listDomains", nil)

	assert.NoError(t, err)

	type JSONDomain struct {
		Domain  string `json:"domain"`
		Created string `json:"created"`
		Expires string `json:"expires"`
	}

	type JSONResult struct {
		Reply struct {
			Code    int          `json:"code"`
			Domains []JSONDomain `json:"domains"`
		} `json:"reply"`
	}

	var result JSONResult
	err = mockClient.ParseResponse(resp, &result)
	assert.NoError(t, err)
	assert.Equal(t, 300, result.Reply.Code)
	assert.Equal(t, 4, len(result.Reply.Domains))

	// 验证关键域名
	assert.Equal(t, "example1.com", result.Reply.Domains[0].Domain)
	assert.Equal(t, "example3.com", result.Reply.Domains[2].Domain)

	t.Logf("完整 JSON 数据解析到 %d 个域名", len(result.Reply.Domains))
}
