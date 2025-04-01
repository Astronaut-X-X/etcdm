package client

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

type GetAllKeyParams struct {
	Keyword string `json:"keyword"`
}

func GetAllKey[T []string, P GetAllKeyParams](c *Client, param P) (T, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()

	ops := []clientv3.OpOption{clientv3.WithPrefix(), clientv3.WithKeysOnly()}
	resp, err := c.Client.Get(ctx, "", ops...)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
	}

	return keys, nil
}
