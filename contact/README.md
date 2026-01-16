# 👤 Contact - 联系人管理

联系人管理模块提供完整的联系人信息 CRUD 操作和域名关联功能。

## 📋 功能列表

- ✅ 添加联系人
- ✅ 更新联系人信息
- ✅ 删除联系人
- ✅ 列出所有联系人
- ✅ 关联联系人到域名

## 🚀 快速开始

```go
import (
    "context"
    "github.com/kamalyes/go-namesilo/client"
    "github.com/kamalyes/go-namesilo/contact"
)

c, _ := client.New("your-api-key")
contactService := contact.NewService(c)
ctx := context.Background()
```

## 💡 使用示例

### 添加联系人

```go
addReq := &contact.ContactAddRequest{
    FirstName: "John",
    LastName:  "Doe",
    Address:   "123 Main St",
    City:      "New York",
    State:     "NY",
    Zip:       "10001",
    Country:   "US",
    Email:     "john@example.com",
    Phone:     "+12125551234",
}
contactResp, err := contactService.AddContact(ctx, addReq)
if err == nil {
    fmt.Printf("联系人ID: %s\n", contactResp.ContactID)
}
```

### 更新联系人

```go
updateReq := &contact.ContactUpdateRequest{
    ContactID: "123456",
    Email:     "newemail@example.com",
    Phone:     "+12125559999",
}
_, err := contactService.UpdateContact(ctx, updateReq)
```

### 列出联系人

```go
listResp, err := contactService.ListContacts(ctx, &contact.ContactListRequest{})
for _, c := range listResp.Contacts {
    fmt.Printf("%s %s (%s)\n", c.FirstName, c.LastName, c.Email)
}
```

### 关联联系人到域名

```go
associateReq := &contact.ContactDomainAssociateRequest{
    Domain:         "example.com",
    Registrant:     "123456",
    Administrative: "123456",
    Technical:      "123456",
    Billing:        "123456",
}
_, err := contactService.AssociateContactToDomain(ctx, associateReq)
```

## 📖 API 文档

- [NameSilo Contact API](https://www.namesilo.com/api-reference#contact)
- [GoDoc](https://pkg.go.dev/github.com/kamalyes/go-namesilo/contact)
