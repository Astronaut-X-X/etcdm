package api

import (
	"etcdm/internal/dto"
	"etcdm/internal/service"
)

type KeyApi struct {
	keyService        *service.KeyService
	connectionService *service.ConnectionService
}

func NewKeyApi() *KeyApi {
	return &KeyApi{
		keyService:        service.NewKeyService(),
		connectionService: service.NewConnectionService(),
	}
}

func (a *KeyApi) GetKeysTree(req *dto.GetKeysTreeRequest) *dto.BaseResponse {
	conn, err := a.connectionService.GetConnection(req.ID)
	if err != nil {
		return dto.FailResp(err.Error())
	}

	keys, err := a.keyService.GetKeysTree(conn, req)
	if err != nil {
		return dto.FailResp(err.Error())
	}

	return dto.SuccessRespData(keys)
}
