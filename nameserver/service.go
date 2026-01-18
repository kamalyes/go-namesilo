/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:49:28
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:58:52
 * @FilePath: \go-namesilo\nameserver\service.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package nameserver

// Service 域名服务器服务
type Service struct {
	client ClientInterface
}

// NewService 创建域名服务器服务
func NewService(c ClientInterface) *Service {
	return &Service{
		client: c,
	}
}
