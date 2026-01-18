/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-19 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-19 00:00:00
 * @FilePath: \go-namesilo\types\client.go
 * @Description: 客户端接口定义
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package types

import (
	"context"
)

// ClientInterface 定义客户端接口
type ClientInterface interface {
	DoRequest(ctx context.Context, operation string, params map[string]string) ([]byte, error)
	DoPublicRequest(ctx context.Context, method, url string, body interface{}, result interface{}) error
	ParseResponse(data []byte, v interface{}) error
}
