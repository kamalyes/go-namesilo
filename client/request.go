/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:08:32
 * @FilePath: \go-namesilo\client\request.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	namesilo "github.com/kamalyes/go-namesilo"
)

// DoRequest 执行 API 请求（公开给子模块使用）
func (c *Client) DoRequest(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
	// 构建请求参数
	queryParams := url.Values{}
	queryParams.Set("version", DefaultAPIVersion)
	queryParams.Set("type", DefaultType)
	queryParams.Set("key", c.config.APIKey)

	// 添加其他参数
	for key, value := range params {
		if value != "" {
			queryParams.Set(key, value)
		}
	}

	// 构建完整 URL
	fullURL := fmt.Sprintf("%s/%s?%s", c.config.BaseURL, operation, queryParams.Encode())

	if c.config.Debug {
		fmt.Printf("[NameSilo] Request: %s\n", fullURL)
	}

	// 发送请求
	resp, err := c.httpClient.Get(fullURL).Do(ctx)
	if err != nil {
		return nil, namesilo.NewRequestError(err)
	}

	if c.config.Debug {
		fmt.Printf("[NameSilo] Response: %s\n", string(resp))
	}

	return resp, nil
}

// ParseResponse 解析响应（公开给子模块使用）
func (c *Client) ParseResponse(data []byte, v interface{}) error {
	// 默认尝试 JSON 解析
	if err := json.Unmarshal(data, v); err == nil {
		return nil
	}

	// 如果 JSON 解析失败，尝试 XML 解析
	if err := xml.Unmarshal(data, v); err != nil {
		return namesilo.NewParseError(err)
	}

	return nil
}
