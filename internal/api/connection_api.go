package api

import (
	"etcdm/internal/dto"
	"etcdm/internal/model"
	"etcdm/internal/service"
)

type ConnectionApi struct {
	connectionService *service.ConnectionService
}

func NewConnectionApi() *ConnectionApi {
	return &ConnectionApi{
		connectionService: service.NewConnectionService(),
	}
}

// 获取单个连接信息
func (api *ConnectionApi) GetConnection(req *dto.GetConnectionRequest) *dto.BaseResponse {
	if req == nil {
		return dto.FailResp("invalid request")
	}

	conn, err := api.connectionService.GetConnection(req.ID)
	if err != nil {
		return dto.FailResp(err.Error())
	}

	if conn == nil {
		return dto.FailResp("connection not found")
	}

	return dto.SuccessRespData(conn)
}

// 获取连接列表
func (api *ConnectionApi) ListConnections(req *dto.ListConnectionRequest) *dto.BaseResponse {
	if req == nil {
		return dto.FailResp("invalid request")
	}

	conns, _, err := api.connectionService.ListConnections()
	if err != nil {
		return dto.FailResp(err.Error())
	}

	return dto.SuccessRespData(conns)
}

// 创建新连接
func (api *ConnectionApi) CreateConnection(req *dto.CreateConnectionRequest) *dto.BaseResponse {
	if req == nil {
		return dto.FailResp("invalid request")
	}

	conn := &model.Connection{
		Id:         req.Id,
		Name:       req.Name,
		Host:       req.Host,
		Port:       req.Port,
		EnableAuth: req.EnableAuth,
		Username:   req.Username,
		Password:   req.Password,
		EnableTLS:  req.EnableTLS,
		CA:         req.CA,
		Cert:       req.Cert,
		Key:        req.Key,
	}

	if err := api.connectionService.CreateConnection(conn); err != nil {
		return dto.FailResp(err.Error())
	}

	return dto.SuccessRespData(conn.Id)
}

// 更新连接信息
func (api *ConnectionApi) UpdateConnection(req *dto.UpdateConnectionRequest) *dto.BaseResponse {
	if req == nil {
		return dto.FailResp("invalid request")
	}

	conn := &model.Connection{
		Id:         req.Id,
		Name:       req.Name,
		Host:       req.Host,
		Port:       req.Port,
		EnableAuth: req.EnableAuth,
		Username:   req.Username,
		Password:   req.Password,
		EnableTLS:  req.EnableTLS,
		CA:         req.CA,
		Cert:       req.Cert,
		Key:        req.Key,
	}

	if err := api.connectionService.UpdateConnection(conn); err != nil {
		return dto.FailResp(err.Error())
	}

	return dto.SuccessResp()
}

// 删除连接
func (api *ConnectionApi) DeleteConnection(req *dto.DeleteConnectionRequest) *dto.BaseResponse {
	if req == nil {
		return dto.FailResp("invalid request")
	}

	if err := api.connectionService.DeleteConnection(req.Id); err != nil {
		return dto.FailResp(err.Error())
	}

	return dto.SuccessResp()
}

// 连接测试
func (api *ConnectionApi) TestConnection(req *dto.TestConnectionRequest) *dto.BaseResponse {
	if req == nil {
		return dto.FailResp("invalid request")
	}

	success, err := api.connectionService.TestConnection(&req.Connection)
	if err != nil {
		return dto.FailResp(err.Error())
	}
	if !success {
		return dto.FailResp("connection test FailResp")
	}

	return dto.SuccessResp()
}
