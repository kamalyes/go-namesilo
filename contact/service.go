/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 22:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:30:00
 * @FilePath: \go-namesilo\contact\service.go
 * @Description: 联系人服务
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package contact

// Service 联系人服务
type Service struct {
	client ClientInterface
}

// NewService 创建联系人服务
func NewService(c ClientInterface) *Service {
	return &Service{
		client: c,
	}
}
