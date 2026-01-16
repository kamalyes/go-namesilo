# 🔄 Transfer - 域名转移管理模块

域名转移管理模块提供域名转移相关的所有操作。

## 📋 功能列表

- ✅ 获取域名授权码 (EPP Code)
- ✅ 检查域名转移状态
- ✅ 重新提交转移到注册局
- ✅ 重新发送转移管理员邮件
- ✅ 更改转移 EPP 授权码

## 🚀 快速开始

```go
import "github.com/kamalyes/go-namesilo/transfer"

transferService := transfer.NewService(client)

// 获取授权码
authResp, _ := transferService.RetrieveAuthCode(ctx, &transfer.RetrieveAuthCodeRequest{
    Domain: "example.com",
})
fmt.Printf("授权码: %s\n", authResp.Reply.AuthCode)

// 检查转移状态
statusResp, _ := transferService.CheckTransferStatus(ctx, &transfer.CheckTransferStatusRequest{
    Domain: "example.com",
})
fmt.Printf("状态: %s\n", statusResp.Reply.Transfer.Status)
```

详细文档请参考 [GoDoc](https://pkg.go.dev/github.com/kamalyes/go-namesilo/transfer)
