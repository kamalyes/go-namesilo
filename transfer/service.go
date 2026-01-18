/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\transfer\service.go
 * @Description: Transfer 域名转移服务
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package transfer

// Service Transfer 域名转移服务
type Service struct {
	client ClientInterface
}

// NewService 创建 Transfer 服务
func NewService(c ClientInterface) *Service {
	return &Service{
		client: c,
	}
}
