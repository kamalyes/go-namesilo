/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-19 10:15:46
 * @FilePath: \go-namesilo\dns\service.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package dns

// Service DNS 服务
type Service struct {
	client ClientInterface
}

// NewService 创建 DNS 服务
func NewService(c ClientInterface) *Service {
	return &Service{
		client: c,
	}
}
