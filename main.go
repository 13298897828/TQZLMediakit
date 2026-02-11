package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ZLMediaKitClient 是 ZLMediaKit 的客户端
type ZLMediaKitClient struct {
	BaseURL string
	Timeout time.Duration
}

// NewZLMediaKitClient 创建一个新的 ZLMediaKit 客户端
func NewZLMediaKitClient(baseURL string, timeout time.Duration) *ZLMediaKitClient {
	return &ZLMediaKitClient{
		BaseURL: baseURL,
		Timeout: timeout,
	}
}

// StreamInfo 是流的信息
type StreamInfo struct {
	App         string `json:"app"`
	Stream      string `json:"stream"`
	OriginURL   string `json:"origin_url"`
	CreateTime  string `json:"create_time"`
	AliveSecond int    `json:"alive_second"`
	BytesSpeed  int    `json:"bytes_speed"`
	TotalBytes  int64  `json:"total_bytes"`
}

// GetStreamList 获取流列表
func (c *ZLMediaKitClient) GetStreamList() ([]StreamInfo, error) {
	url := fmt.Sprintf("%s/index/api/getMediaList", c.BaseURL)
	client := &http.Client{Timeout: c.Timeout}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Code    int          `json:"code"`
		Message string       `json:"msg"`
		Data    []StreamInfo `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, fmt.Errorf("获取流列表失败: %s", result.Message)
	}

	return result.Data, nil
}

// GetHLSPullURL 获取 HLS 拉流地址
func (c *ZLMediaKitClient) GetHLSPullURL(app, stream string) string {
	return fmt.Sprintf("%s/%s/%s/hls/index.m3u8", c.BaseURL, app, stream)
}

// GetStreamDetail 获取流详情
func (c *ZLMediaKitClient) GetStreamDetail(app, stream string) (*StreamInfo, error) {
	url := fmt.Sprintf("%s/index/api/getMediaInfo?app=%s&stream=%s", c.BaseURL, app, stream)
	client := &http.Client{Timeout: c.Timeout}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Code    int          `json:"code"`
		Message string       `json:"msg"`
		Data    *StreamInfo  `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, fmt.Errorf("获取流详情失败: %s", result.Message)
	}

	return result.Data, nil
}

// PlaybackHLS 播放 HLS 流
func (c *ZLMediaKitClient) PlaybackHLS(app, stream string) (string, error) {
	// 检查流是否存在
	_, err := c.GetStreamDetail(app, stream)
	if err != nil {
		return "", err
	}

	// 返回 HLS 拉流地址
	return c.GetHLSPullURL(app, stream), nil
}

func main() {
	// 示例用法
	client := NewZLMediaKitClient("http://127.0.0.1:8080", 30*time.Second)

	// 获取流列表
	streams, err := client.GetStreamList()
	if err != nil {
		fmt.Printf("获取流列表失败: %v\n", err)
		return
	}

	fmt.Printf("当前流列表: %d 个\n", len(streams))
	for _, stream := range streams {
		fmt.Printf("App: %s, Stream: %s, 在线时间: %d 秒\n", stream.App, stream.Stream, stream.AliveSecond)
	}

	// 播放示例流
	if len(streams) > 0 {
		playbackURL, err := client.PlaybackHLS(streams[0].App, streams[0].Stream)
		if err != nil {
			fmt.Printf("播放流失败: %v\n", err)
			return
		}

		fmt.Printf("播放地址: %s\n", playbackURL)
	}
}