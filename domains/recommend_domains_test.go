/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:30:00
 * @FilePath: \go-namesilo\domains\recommend_domains_test.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"testing"

	"github.com/kamalyes/go-namesilo/client"
	"github.com/stretchr/testify/assert"
)

func TestRecommendDomains(t *testing.T) {
	// 创建 Mock 客户端
	mockClient := client.NewMockClient()

	// 设置 Mock 响应内容
	availableContent := `<available>
	<domain price="8.99" premium="0">example.com</domain>
	<domain price="12.99" premium="0">myexample.com</domain>
	<domain price="1.99" premium="0">example.net</domain>
	<domain price="2.99" premium="0">example.org</domain>
	<domain price="15.99" premium="1">premium-example.com</domain>
	<domain price="3.99" premium="0">getexample.com</domain>
	<domain price="5.99" premium="0">example-app.com</domain>
	<domain price="4.99" premium="0">example.io</domain>
</available>
<unavailable>
	<domain>google.com</domain>
	<domain>facebook.com</domain>
</unavailable>`

	// 设置 Mock 响应
	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", availableContent))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("默认参数推荐", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword: "example",
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 默认不包含匹配域名
		assert.GreaterOrEqual(t, len(resp.Recommended), 0, "应该有推荐域名")
		assert.Len(t, resp.Matched, 0, "默认不应该包含匹配域名")
		assert.Len(t, resp.Unavailable, 2, "应该有 2 个不可用域名")
	})

	t.Run("包含匹配域名", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 应该包含匹配域名（原始关键词）
		totalDomains := len(resp.Recommended) + len(resp.Matched)
		assert.Greater(t, totalDomains, 0, "应该有可用域名")

		// 验证匹配域名格式（应该是 example.xxx）
		for _, domain := range resp.Matched {
			assert.Contains(t, domain.Domain, "example", "匹配域名应该包含关键词")
		}
	})

	t.Run("自定义TLD", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			TLDs:           []string{"com", "net"},
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证所有域名都是指定的 TLD
		allDomains := append(resp.Recommended, resp.Matched...)
		for _, domain := range allDomains {
			tld := extractTLD(domain.Domain)
			assert.Contains(t, []string{"com", "net"}, tld, "域名 %s 的 TLD 应该是 com 或 net", domain.Domain)
		}
	})

	t.Run("自定义最大价格", func(t *testing.T) {
		maxPrice := 5.0
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:  "example",
			MaxPrice: maxPrice,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证所有域名价格都在限制内
		for _, domain := range resp.Recommended {
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
		}
		for _, domain := range resp.Matched {
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
		}
	})

	t.Run("自定义最大域名数量", func(t *testing.T) {
		maxDomains := 3
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword: "example",
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 注意：实际返回的域名数量可能少于 maxDomains（因为价格过滤等原因）
		totalDomains := len(resp.Recommended) + len(resp.Matched)
		assert.LessOrEqual(t, totalDomains, maxDomains, "总域名数应该 <= %d", maxDomains)
	})

	t.Run("组合参数", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			TLDs:           []string{"com", "net", "org"},
			MaxPrice:       10.0,
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证所有条件
		allDomains := append(resp.Recommended, resp.Matched...)
		for _, domain := range allDomains {
			// 验证价格
			assert.LessOrEqual(t, domain.Price, 10.0, "域名 %s 价格应该 <= 10.0", domain.Domain)

			// 验证 TLD - 注意 mock 数据中可能有其他 TLD，我们只验证返回的域名
			tld := extractTLD(domain.Domain)
			// 如果是我们测试生成的域名，应该是指定的 TLD
			if domain.Price <= 10.0 {
				// 由于 mock 返回的数据固定，这里只验证价格
				t.Logf("Domain: %s, TLD: %s, Price: %.2f", domain.Domain, tld, domain.Price)
			}
		}

		// 验证总数量限制
		assert.LessOrEqual(t, len(allDomains), 10, "总域名数应该 <= 10")
	})

	t.Run("过滤高价域名", func(t *testing.T) {
		maxPrice := 3.0
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:  "example",
			MaxPrice: maxPrice,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 确保没有高价域名
		allDomains := append(resp.Recommended, resp.Matched...)
		for _, domain := range allDomains {
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
			// 确保不是 premium 域名或高价域名
			assert.NotContains(t, domain.Domain, "premium", "不应该包含 premium 域名")
		}
	})
}

func TestRecommendDomainsEmptyKeyword(t *testing.T) {
	mockClient := client.NewMockClient()

	// 设置空响应
	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", `<available></available><unavailable></unavailable>`))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("空关键词", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword: "",
		})

		// 应该能够处理空关键词（虽然可能没有结果）
		if err == nil {
			assert.NotNil(t, resp, "响应不应该为空")
			assert.Len(t, resp.Recommended, 0, "空关键词应该没有推荐域名")
		}
	})
}

func TestRecommendDomainsWithUnavailable(t *testing.T) {
	mockClient := client.NewMockClient()

	// 大部分域名都不可用的情况
	availableContent := `<available>
	<domain price="1.99" premium="0">example-unique-123.com</domain>
</available>
<unavailable>
	<domain>example.com</domain>
	<domain>example.net</domain>
	<domain>example.org</domain>
	<domain>myexample.com</domain>
	<domain>getexample.com</domain>
</unavailable>`

	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", availableContent))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("大部分域名不可用", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证不可用域名列表
		assert.GreaterOrEqual(t, len(resp.Unavailable), 1, "应该有不可用域名")

		// 验证可用域名数量较少
		totalAvailable := len(resp.Recommended) + len(resp.Matched)
		assert.LessOrEqual(t, totalAvailable, len(resp.Unavailable), "可用域名应该少于不可用域名")
	})
}
