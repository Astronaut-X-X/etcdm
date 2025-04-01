package api

import (
	"etcdm/internal/dto"
	"etcdm/internal/service"
)

type DataApi struct {
	keyService        *service.KeyService
	connectionService *service.ConnectionService
	dataService       *service.DataService
}

func NewDataApi() *DataApi {
	return &DataApi{
		keyService:        service.NewKeyService(),
		connectionService: service.NewConnectionService(),
		dataService:       service.NewDataService(),
	}
}

func (a *DataApi) ListData(req *dto.ListDataRequest) *dto.BaseResponse {
	conn, err := a.connectionService.GetConnection(req.Params.ID)
	if err != nil {
		return dto.FailResp(err.Error())
	}

	data, total, err := a.dataService.ListData(conn, req)
	if err != nil {
		return dto.FailResp(err.Error())
	}
	return dto.SuccessRespPage(data, total)
}
