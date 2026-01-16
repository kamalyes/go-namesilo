# 📁 Portfolio - 域名组合管理模块

域名组合管理模块允许您将域名组织到不同的组合(文件夹)中。

## 📋 功能列表

- ✅ 列出所有域名组合
- ✅ 创建新域名组合
- ✅ 删除域名组合
- ✅ 关联域名到组合

## 🚀 快速开始

```go
import "github.com/kamalyes/go-namesilo/portfolio"

portfolioService := portfolio.NewService(client)

// 列出所有组合
listResp, _ := portfolioService.List(ctx, &portfolio.PortfolioListRequest{})
for _, p := range listResp.Reply.Portfolios {
    fmt.Printf("组合: %s (包含 %d 个域名)\n", p.Name, p.DomainCount)
}

// 创建新组合
portfolioService.Add(ctx, &portfolio.PortfolioAddRequest{
    Portfolio: "my-domains",
})

// 关联域名到组合
portfolioService.DomainAssociate(ctx, &portfolio.PortfolioDomainAssociateRequest{
    Domains:   []string{"example.com", "test.com"},
    Portfolio: "my-domains",
})
```

详细文档请参考 [GoDoc](https://pkg.go.dev/github.com/kamalyes/go-namesilo/portfolio)
