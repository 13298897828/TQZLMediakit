# TQZLMediakit - ZLMediaKit HLS 直播回放工具

这是一个用于与 ZLMediaKit 服务器交互的 Go 语言工具，提供 HLS 直播回放功能。

## 功能特性

- 获取 ZLMediaKit 服务器上的流列表
- 获取指定流的详细信息
- 生成 HLS 拉流地址
- 播放 HLS 流
- 配置文件支持

## 安装和使用

### 1. 克隆仓库

```bash
git clone https://github.com/13298897828/TQZLMediakit.git
cd TQZLMediakit
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置

修改 `config/config.yaml` 中的配置：

```yaml
# ZLMediaKit 配置文件
server:
  # ZLMediaKit 服务器地址
  base_url: "http://127.0.0.1:8080"
  # 请求超时时间（秒）
  timeout: 30
```

### 4. 运行

```bash
go run main.go
```

### 5. 输出示例

```
配置文件读取成功
当前流列表: 2 个
App: live, Stream: test1, 在线时间: 3600 秒
App: live, Stream: test2, 在线时间: 1800 秒
播放地址: http://127.0.0.1:8080/live/test1/hls/index.m3u8
```

## API 说明

### ZLMediaKitClient

```go
// 创建一个新的 ZLMediaKit 客户端
func NewZLMediaKitClient(baseURL string, timeout time.Duration) *ZLMediaKitClient
```

### 获取流列表

```go
// 获取流列表
func (c *ZLMediaKitClient) GetStreamList() ([]StreamInfo, error)
```

### 获取流详情

```go
// 获取流详情
func (c *ZLMediaKitClient) GetStreamDetail(app, stream string) (*StreamInfo, error)
```

### 获取 HLS 拉流地址

```go
// 获取 HLS 拉流地址
func (c *ZLMediaKitClient) GetHLSPullURL(app, stream string) string
```

### 播放 HLS 流

```go
// 播放 HLS 流
func (c *ZLMediaKitClient) PlaybackHLS(app, stream string) (string, error)
```

## 依赖

- Go 1.21+
- github.com/spf13/viper：用于读取配置文件
- 网络连接到 ZLMediaKit 服务器

## 注意事项

- 确保 ZLMediaKit 服务器的 API 端口（默认 8080）可以访问
- 如果使用 HTTPS，需要修改 BaseURL 为 https 协议
- 可以根据需要调整超时时间（当前默认 30 秒）

## License

MIT