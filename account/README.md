# 💰 Account - 账户管理

账户管理模块提供账户余额、订单、价格查询等功能。

## 📋 功能列表

- ✅ 查询账户余额
- ✅ 添加账户资金  
- ✅ 查询订单详情
- ✅ 列出订单
- ✅ 列出即将到期的域名
- ✅ 统计即将到期的域名数量

## 🚀 快速开始

```go
import "github.com/kamalyes/go-namesilo/account"

accountService := account.NewService(client)
```

## 💡 使用示例

### 查询账户余额

```go
balance, err := accountService.GetBalance(ctx, &account.GetAccountBalanceRequest{})
if err == nil {
    fmt.Printf("账户余额: $%s\n", balance.Balance)
}
```

### 查询即将到期的域名

```go
expiringReq := &account.ListExpiringDomainsRequest{
    DaysCount: 30,
    Page:      1,
    PageSize:  10,
}
expiring, err := accountService.ListExpiringDomains(ctx, expiringReq)
for _, domain := range expiring.Domains {
    fmt.Printf("域名 %s 将在 %s 到期\n", domain.Domain, domain.Expires)
}
```

### 统计即将到期的域名数量

```go
countReq := &account.CountExpiringDomainsRequest{DaysCount: 30}
count, err := accountService.CountExpiringDomains(ctx, countReq)
if err == nil {
    fmt.Printf("30天内到期的域名数量: %d\n", count.Count)
}
```

### 查询订单详情

```go
orderReq := &account.OrderDetailsRequest{OrderNumber: "12345"}
order, err := accountService.GetOrderDetails(ctx, orderReq)
```

### 列出订单

```go
listReq := &account.ListOrdersRequest{
    Page:     1,
    PageSize: 20,
}
orders, err := accountService.ListOrders(ctx, listReq)
for _, order := range orders.Orders {
    fmt.Printf("订单 %s: $%.2f\n", order.OrderNumber, order.OrderAmount)
}
```

## 📖 API 文档

- [NameSilo Account API](https://www.namesilo.com/api-reference#account)
- [GoDoc](https://pkg.go.dev/github.com/kamalyes/go-namesilo/account)
