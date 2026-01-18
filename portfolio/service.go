/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:00:00
 * @FilePath: \go-namesilo\portfolio\service.go
 * @Description: Portfolio 域名组合管理服务
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package portfolio

// Service Portfolio 服务
type Service struct {
	client ClientInterface
}

// NewService 创建 Portfolio 服务
func NewService(c ClientInterface) *Service {
	return &Service{
		client: c,
	}
}
