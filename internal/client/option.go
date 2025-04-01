package client

import "context"

type Option[T any, P any] func(*Client, P) (T, error)

// 测试与etcd服务器的连接
func WithTestConnection[T bool, P any](c *Client, _ P) (T, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()

	_, err := c.Client.Status(ctx, c.Endpoints()[0])
	if err != nil {
		return false, err
	}

	return true, nil
}
