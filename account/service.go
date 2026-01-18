/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:30:00
 * @FilePath: \go-namesilo\account\service.go
 * @Description: 账户服务
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package account

// Service 账号服务
type Service struct {
	client ClientInterface
}

// NewService 创建账号服务
func NewService(c ClientInterface) *Service {
	return &Service{
		client: c,
	}
}
