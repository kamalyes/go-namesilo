/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:29:32
 * @FilePath: \go-namesilo\client\client.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"github.com/kamalyes/go-toolbox/pkg/httpx"
)

// Client NameSilo API 客户端
type Client struct {
	config     *Config
	httpClient *httpx.Client
}

// New 创建一个新的 NameSilo 客户端
func New(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, ErrAPIKeyRequired
	}

	config := &Config{
		APIKey:  apiKey,
		BaseURL: DefaultBaseURL,
		Timeout: DefaultTimeout,
		Debug:   true,
	}

	// 应用选项
	for _, opt := range opts {
		opt(config)
	}

	// 创建 HTTP 客户端
	httpClient := httpx.NewClient(
		httpx.WithTimeout(config.Timeout),
	)

	client := &Client{
		config:     config,
		httpClient: httpClient,
	}

	return client, nil
}

// GetConfig 获取客户端配置
func (c *Client) GetConfig() *Config {
	return c.config
}

// GetHTTPClient 获取 HTTP 客户端
func (c *Client) GetHTTPClient() *httpx.Client {
	return c.httpClient
}
