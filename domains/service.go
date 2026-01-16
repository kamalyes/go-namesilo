/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:56:27
 * @FilePath: \go-namesilo\domains\service.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"github.com/kamalyes/go-namesilo/client"
)

// Service 域名服务
type Service struct {
	client *client.Client
}

// NewService 创建域名服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
