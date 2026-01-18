/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:35:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-19 10:17:17
 * @FilePath: \go-namesilo\forwarding\service.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package forwarding

// Service 转发服务
type Service struct {
	client ClientInterface
}

// NewService 创建转发服务
func NewService(c ClientInterface) *Service {
	return &Service{client: c}
}
