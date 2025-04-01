package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"etcdm/internal/model"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func BuildConfig(ctx context.Context, conn *model.Connection) clientv3.Config {
	// 创建 etcd 客户端配置
	endpoint := fmt.Sprintf("%s:%d", conn.Host, conn.Port)
	config := clientv3.Config{
		Context:              ctx,
		Endpoints:            []string{endpoint},
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    5 * time.Second,
		DialKeepAliveTimeout: 5 * time.Second,
	}

	// 如果启用了认证，则配置认证
	if conn.EnableAuth {
		config.Username = conn.Username
		config.Password = conn.Password
	}

	// 如果启用了 TLS，则配置 TLS
	if conn.EnableTLS {
		config.TLS = buildTLSConfig(conn)
	}

	return config
}

func buildTLSConfig(conn *model.Connection) *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true, // 跳过 TLS 证书验证
		MinVersion:         tls.VersionTLS12,
		MaxVersion:         tls.VersionTLS13,
		// 如果提供了 CA 证书，则加载并配置
		RootCAs: func() *x509.CertPool {
			if conn.CA != "" {
				pool := x509.NewCertPool()
				pool.AppendCertsFromPEM([]byte(conn.CA))
				return pool
			}
			return nil
		}(),
		// 如果提供了客户端证书和密钥，则加载并配置
		Certificates: func() []tls.Certificate {
			if conn.Cert != "" && conn.Key != "" {
				cert, err := tls.X509KeyPair([]byte(conn.Cert), []byte(conn.Key))
				if err == nil {
					return []tls.Certificate{cert}
				}
			}
			return nil
		}(),
	}
}
