/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:20:00
 * @FilePath: \go-namesilo\dns\service.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package dns

import (
	"github.com/kamalyes/go-namesilo/client"
)

// Service DNS 服务
type Service struct {
	client *client.Client
}

// NewService 创建 DNS 服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
