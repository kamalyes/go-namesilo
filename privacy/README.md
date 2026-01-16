# 🔒 Privacy - WHOIS 隐私保护模块

隐私保护模块用于管理域名的 WHOIS 隐私保护服务。

## 📋 功能列表

- ✅ 添加域名 WHOIS 隐私保护
- ✅ 移除域名 WHOIS 隐私保护

## 🚀 快速开始

```go
import "github.com/kamalyes/go-namesilo/privacy"

privacyService := privacy.NewService(client)

// 添加隐私保护
privacyService.AddPrivacy(ctx, &privacy.AddPrivacyRequest{
    Domain: "example.com",
})

// 移除隐私保护  
privacyService.RemovePrivacy(ctx, &privacy.RemovePrivacyRequest{
    Domain: "example.com",
})
```

详细文档请参考 [GoDoc](https://pkg.go.dev/github.com/kamalyes/go-namesilo/privacy)
