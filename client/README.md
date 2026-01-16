# 🔌 Client HTTP 客户端

NameSilo API 的 HTTP 客户端封装,提供请求/响应处理、错误处理、日志记录等功能。

## 📦 安装

```bash
go get -u github.com/kamalyes/go-namesilo
```

## 🚀 快速开始

### 基础用法

```go
import "github.com/kamalyes/go-namesilo/client"

// 创建客户端
c, err := client.New("your-api-key")
if err != nil {
    log.Fatal(err)
}
```

### 自定义配置

```go
import (
    "time"
    "github.com/kamalyes/go-namesilo/client"
)

c, err := client.New(
    "your-api-key",
    client.WithTimeout(30 * time.Second),      // 自定义超时
    client.WithBaseURL("https://api.namesilo.com/api"), // 自定义 API 地址
    client.WithDebug(true),                     // 开启调试模式
    client.WithRetryConfig(3, 1*time.Second),  // 重试配置: 3次, 间隔1秒
)
if err != nil {
    log.Fatal(err)
}
```

## 🎯 配置选项

### WithAPIKey

覆盖默认 API Key(优先级高于 New 的第一个参数)。

```go
client.New("default-key", client.WithAPIKey("override-key"))
// 结果: 使用 "override-key"
```

### WithBaseURL

自定义 API 基础地址。

```go
client.WithBaseURL("https://sandbox.namesilo.com/api")
```

**默认值**: `https://www.namesilo.com/api`

### WithTimeout

设置 HTTP 请求超时时间。

```go
client.WithTimeout(60 * time.Second)
```

**默认值**: `30 秒`

### WithDebug

开启调试模式,输出详细的请求/响应日志。

```go
client.WithDebug(true)
```

**默认值**: `false`

### WithRetryConfig

配置请求重试机制。

```go
client.WithRetryConfig(
    5,                  // 最大重试次数
    2 * time.Second,   // 重试间隔
)
```

**默认值**: 
- 重试次数: `3`
- 重试间隔: `1 秒`

### WithLogger

自定义日志记录器。

```go
type CustomLogger struct{}

func (l *CustomLogger) Printf(format string, v ...interface{}) {
    log.Printf("[MY-LOG] "+format, v...)
}

func (l *CustomLogger) Println(v ...interface{}) {
    log.Println(append([]interface{}{"[MY-LOG]"}, v...)...)
}

client.New("api-key", client.WithLogger(&CustomLogger{}))
```

## 💡 配置优先级

API Key 配置遵循以下优先级(从高到低):

1. **WithAPIKey 选项** - `client.WithAPIKey("key")`
2. **New 函数参数** - `client.New("key")`
3. **环境变量** - `NAMESILO_API_KEY`

```go
// 示例 1: WithAPIKey 优先级最高
os.Setenv("NAMESILO_API_KEY", "env-key")
c, _ := client.New("param-key", client.WithAPIKey("option-key"))
// 结果: 使用 "option-key"

// 示例 2: 参数优先于环境变量
os.Setenv("NAMESILO_API_KEY", "env-key")
c, _ := client.New("param-key")
// 结果: 使用 "param-key"

// 示例 3: 使用环境变量
os.Setenv("NAMESILO_API_KEY", "env-key")
c, _ := client.New("")
// 结果: 使用 "env-key"
```

## 🧪 测试支持

### Mock 客户端

用于单元测试的 Mock 客户端。

```go
import (
    "testing"
    "github.com/kamalyes/go-namesilo/client"
)

func TestYourFunction(t *testing.T) {
    // 创建 Mock 客户端
    mockClient := client.NewMockClient()
    
    // 配置 Mock 响应
    mockClient.SetResponse(&client.MockResponse{
        StatusCode: 200,
        Body: `<?xml version="1.0"?>
        <namesilo>
            <request>
                <operation>registerDomain</operation>
                <ip>127.0.0.1</ip>
            </request>
            <reply>
                <code>300</code>
                <detail>success</detail>
            </reply>
        </namesilo>`,
    })
    
    // 使用 Mock 客户端进行测试
    // ...
}
```

## 🔍 调试技巧

### 开启详细日志

```go
c, _ := client.New("api-key", client.WithDebug(true))

// 输出示例:
// [DEBUG] Request: POST https://www.namesilo.com/api/registerDomain
// [DEBUG] Params: domain=example.com&years=1&private=1
// [DEBUG] Response: {"code":300,"detail":"success"}
```

### 自定义日志格式

```go
type JSONLogger struct{}

func (l *JSONLogger) Printf(format string, v ...interface{}) {
    log.Printf(`{"level":"debug","msg":"`+format+`"}`, v...)
}

func (l *JSONLogger) Println(v ...interface{}) {
    log.Println(map[string]interface{}{
        "level": "debug",
        "msg":   fmt.Sprint(v...),
    })
}

c, _ := client.New("api-key", 
    client.WithDebug(true),
    client.WithLogger(&JSONLogger{}),
)
```

## 🛡️ 错误处理

客户端内部会自动处理以下错误:

- ✅ HTTP 请求失败
- ✅ 响应解析失败
- ✅ API 错误码识别
- ✅ 超时重试

```go
resp, err := someService.SomeMethod(ctx, req)
if err != nil {
    // 统一的错误处理
    log.Printf("API 调用失败: %v", err)
    return
}
```

## 📖 相关资源

- [NameSilo API 文档](https://www.namesilo.com/api-reference)
- [返回主文档](../)

## ⚠️ 注意事项

1. **API Key 安全**: 
   - 切勿在代码中硬编码 API Key
   - 推荐使用环境变量或配置文件
   - 在版本控制系统中忽略包含 Key 的文件

2. **超时设置**: 
   - 根据网络环境调整超时时间
   - 批量操作建议增加超时时间

3. **重试机制**: 
   - 重试仅适用于幂等操作
   - 域名注册等不可逆操作需谨慎

4. **日志安全**: 
   - 生产环境关闭调试模式
   - 避免在日志中泄露敏感信息

5. **并发使用**: 
   - 客户端实例是并发安全的
   - 可以在多个 goroutine 中共享使用

## 🎨 最佳实践

### 单例模式

```go
var (
    once       sync.Once
    apiClient  *client.Client
    clientErr  error
)

func GetClient() (*client.Client, error) {
    once.Do(func() {
        apiClient, clientErr = client.New(
            os.Getenv("NAMESILO_API_KEY"),
            client.WithTimeout(30*time.Second),
        )
    })
    return apiClient, clientErr
}
```

### 上下文超时

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

resp, err := service.SomeMethod(ctx, req)
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        log.Println("请求超时")
    }
    return err
}
```

### 优雅关闭

```go
c, err := client.New("api-key")
if err != nil {
    log.Fatal(err)
}
defer c.Close() // 如果客户端有 Close 方法

// 或使用 context 取消
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// 监听中断信号
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

go func() {
    <-sigCh
    cancel() // 取消所有进行中的请求
}()
```
