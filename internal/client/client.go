package client

import (
	"context"
	"etcdm/internal/model"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Client struct {
	*clientv3.Client

	ctx    context.Context
	cancel context.CancelFunc

	// option
	timeout time.Duration
}

func NewClient(conn *model.Connection) (*Client, error) {
	// 创建 context
	ctx, cancel := context.WithCancel(context.Background())

	// 创建 etcd 客户端
	config := BuildConfig(ctx, conn)
	cli, err := clientv3.New(config)
	if err != nil {
		cancel()
		return nil, err
	}

	// 创建客户端
	c := &Client{
		Client:  cli,
		ctx:     ctx,
		cancel:  cancel,
		timeout: 5 * time.Second,
	}

	return c, nil
}

func (c *Client) Close() {
	c.cancel()
	c.Client.Close()
}

func Execute[T any, P any](conn *model.Connection, option Option[T, P], params P) (T, error) {
	var resp T

	client, err := NewClient(conn)
	if err != nil {
		return resp, err
	}
	defer client.Close()

	return option(client, params)
}
