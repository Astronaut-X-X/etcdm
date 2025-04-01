package service

import (
	"etcdm/internal/cfg"
	"etcdm/internal/client"
	"etcdm/internal/db"
	"etcdm/internal/model"
	"etcdm/internal/pkg"
	"fmt"
	"path"
)

const ConnectionDBFileName = "connections.d"

type ConnectionService struct {
	DB db.IDB[*model.Connection]
}

func NewConnectionService() *ConnectionService {
	dbPath := path.Join(cfg.FilePrefix, ConnectionDBFileName)
	return &ConnectionService{
		DB: db.NewFileDB[*model.Connection](dbPath),
	}
}

// 获取单个连接信息
func (s *ConnectionService) GetConnection(id string) (*model.Connection, error) {
	item, err := s.DB.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// 获取连接列表
func (s *ConnectionService) ListConnections() ([]*model.Connection, int64, error) {
	items, err := s.DB.List()
	if err != nil {
		return nil, 0, err
	}
	return items, int64(len(items)), nil
}

// 创建新连接
func (s *ConnectionService) CreateConnection(conn *model.Connection) error {
	if conn.Id == "" {
		conn.Id = pkg.GenerateUUID()
	}
	return s.DB.Create(conn)
}

// 删除连接
func (s *ConnectionService) DeleteConnection(id string) error {
	return s.DB.Delete(id)
}

// 更新连接信息
func (s *ConnectionService) UpdateConnection(conn *model.Connection) error {
	if conn.Id == "" {
		return fmt.Errorf("connection id cannot be empty")
	}
	return s.DB.Update(conn)
}

// 测试连接
func (s *ConnectionService) TestConnection(conn *model.Connection) (bool, error) {
	return client.Execute(conn, client.WithTestConnection, any(nil))
}
