/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:08:32
 * @FilePath: \go-namesilo\client\options.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"time"
)

// Option 客户端选项函数
type Option func(*Config)

// WithBaseURL 设置 API 基础 URL
func WithBaseURL(baseURL string) Option {
	return func(c *Config) {
		c.BaseURL = baseURL
	}
}

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithDebug 开启调试模式
func WithDebug(debug bool) Option {
	return func(c *Config) {
		c.Debug = debug
	}
}
