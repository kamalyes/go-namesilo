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

// Service Privacy 隐私保护服务
type Service struct {
	client ClientInterface
}

// NewService 创建 Privacy 服务
func NewService(c ClientInterface) *Service {
	return &Service{
		client: c,
	}
}
