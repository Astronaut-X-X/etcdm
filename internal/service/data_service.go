package service

import (
	"etcdm/internal/client"
	"etcdm/internal/dto"
	"etcdm/internal/model"
	"etcdm/internal/pkg"
)

type DataService struct {
}

func NewDataService() *DataService {
	return &DataService{}
}

func (s *DataService) ListData(conn *model.Connection, req *dto.ListDataRequest) ([]*client.KV, int64, error) {
	param := &client.ListDataParams{
		Prefix: req.Params.Prefix,
	}
	resp, err := client.Execute(conn, client.ListData, param)
	if err != nil {
		return nil, 0, err
	}

	data, total := pkg.NewPageHelper(resp, req.Page, req.Size).GetPage()
	return data, total, nil
}
