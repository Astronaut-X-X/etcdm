package client

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

// =========== =========== =========== list data =========== =========== ===========

// 定义一个接口来约束参数类型
type PrefixGetter interface {
	GetPrefix() string
}

// 修改 ListDataParams 实现接口
type ListDataParams struct {
	Prefix string `json:"prefix"`
}

func (p ListDataParams) GetPrefix() string {
	return p.Prefix
}

type KV struct {
	Key            string `json:"key,omitempty"`
	CreateRevision int64  `json:"create_revision,omitempty"`
	ModRevision    int64  `json:"mod_revision,omitempty"`
	Version        int64  ` json:"version,omitempty"`
	Value          string ` json:"value,omitempty"`
	Lease          int64  ` json:"lease,omitempty"`
}

// 修改函数签名，添加类型约束
func ListData[T []*KV, P PrefixGetter](c *Client, param P) (T, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()

	ops := []clientv3.OpOption{clientv3.WithPrefix()}
	resp, err := c.Client.Get(ctx, param.GetPrefix(), ops...)
	if err != nil {
		return nil, err
	}

	keys := make([]*KV, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		keys = append(keys, &KV{
			Key:            string(kv.Key),
			CreateRevision: kv.CreateRevision,
			ModRevision:    kv.ModRevision,
			Version:        kv.Version,
			Value:          string(kv.Value),
			Lease:          kv.Lease,
		})
	}

	return keys, nil
}

// =========== =========== =========== put data =========== =========== ===========

// 定义一个接口来约束参数类型
type KeyValueGetter interface {
	GetValue() string
	GetKey() string
}

type PutDataParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p PutDataParams) GetValue() string {
	return p.Value
}
func (p PutDataParams) GetKey() string {
	return p.Key
}

func PutData[T bool, P KeyValueGetter](c *Client, param P) (T, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()

	_, err := c.Client.Put(ctx, param.GetKey(), param.GetValue())
	if err != nil {
		return false, err
	}

	return true, nil
}

// =========== =========== =========== delete data =========== =========== ===========

func DeleteData[T bool, P string](c *Client, key P) (T, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()

	_, err := c.Client.Delete(ctx, string(key))
	if err != nil {
		return false, err
	}

	return true, nil
}
