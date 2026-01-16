/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:35:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:35:00
 * @FilePath: \go-namesilo\forwarding\service.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package forwarding

import "github.com/kamalyes/go-namesilo/client"

// Service 转发服务
type Service struct {
	client *client.Client
}

// NewService 创建转发服务实例
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}
