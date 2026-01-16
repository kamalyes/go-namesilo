/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:00:00
 * @FilePath: \go-namesilo\privacy\service.go
 * @Description: Privacy 隐私保护服务
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package privacy

import (
	"github.com/kamalyes/go-namesilo/client"
)

// Service Privacy 隐私保护服务
type Service struct {
	client *client.Client
}

// NewService 创建 Privacy 服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
