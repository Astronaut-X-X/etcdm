package api

import (
	"etcdm/internal/dto"
	"etcdm/internal/model"
	"testing"
)

func TestConnectionApi_GetConnection(t *testing.T) {
	api := NewConnectionApi()

	tests := []struct {
		name        string
		req         *dto.GetConnectionRequest
		wantData    bool
		wantSuccess bool
	}{
		{
			name:        "请求为空",
			req:         nil,
			wantData:    false,
			wantSuccess: false,
		},
		{
			name: "正常请求",
			req: &dto.GetConnectionRequest{
				ID: "test-id",
			},
			wantData:    true,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := api.GetConnection(tt.req)
			if resp.Success != tt.wantSuccess {
				t.Errorf("GetConnection() success = %v, want %v", resp.Success, tt.wantSuccess)
			}
			if (resp.Data != nil) != tt.wantData {
				t.Errorf("GetConnection() data presence = %v, want %v", resp.Data != nil, tt.wantData)
			}
		})
	}
}

func TestConnectionApi_ListConnections(t *testing.T) {
	api := NewConnectionApi()

	tests := []struct {
		name        string
		req         *dto.ListConnectionRequest
		wantData    bool
		wantSuccess bool
	}{
		{
			name:        "请求为空",
			req:         nil,
			wantData:    false,
			wantSuccess: false,
		},
		{
			name: "正常请求",
			req: &dto.ListConnectionRequest{
				PageRequest: dto.PageRequest{
					Page: 1,
					Size: 10,
				},
			},
			wantData:    true,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := api.ListConnections(tt.req)
			if resp.Success != tt.wantSuccess {
				t.Errorf("ListConnections() success = %v, want %v", resp.Success, tt.wantSuccess)
			}
			if (resp.Data != nil) != tt.wantData {
				t.Errorf("ListConnections() data presence = %v, want %v", resp.Data != nil, tt.wantData)
			}
		})
	}
}

func TestConnectionApi_CreateConnection(t *testing.T) {
	api := NewConnectionApi()

	tests := []struct {
		name        string
		req         *dto.CreateConnectionRequest
		wantData    bool
		wantSuccess bool
	}{
		{
			name:        "请求为空",
			req:         nil,
			wantData:    false,
			wantSuccess: false,
		},
		{
			name: "正常请求",
			req: &dto.CreateConnectionRequest{
				Connection: model.Connection{
					Name: "test-conn",
					Host: "localhost",
					Port: 2379,
				},
			},
			wantData:    true,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := api.CreateConnection(tt.req)
			if resp.Success != tt.wantSuccess {
				t.Errorf("CreateConnection() success = %v, want %v", resp.Success, tt.wantSuccess)
			}
			if (resp.Data != nil) != tt.wantData {
				t.Errorf("CreateConnection() data presence = %v, want %v", resp.Data != nil, tt.wantData)
			}
		})
	}
}

func TestConnectionApi_UpdateConnection(t *testing.T) {
	api := NewConnectionApi()

	tests := []struct {
		name        string
		req         *dto.UpdateConnectionRequest
		wantSuccess bool
	}{
		{
			name:        "请求为空",
			req:         nil,
			wantSuccess: false,
		},
		{
			name: "正常请求",
			req: &dto.UpdateConnectionRequest{
				Connection: model.Connection{
					Id:   "2c3fe09d-8fea-4949-a61e-e1d5a3a71378",
					Name: "test-conn-0000",
					Host: "localhost",
					Port: 2379,
				},
			},
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := api.UpdateConnection(tt.req)
			if resp.Success != tt.wantSuccess {
				t.Errorf("UpdateConnection() success = %v, want %v", resp.Success, tt.wantSuccess)
			}
		})
	}
}

func TestConnectionApi_DeleteConnection(t *testing.T) {
	api := NewConnectionApi()

	tests := []struct {
		name        string
		req         *dto.DeleteConnectionRequest
		wantSuccess bool
	}{
		{
			name:        "请求为空",
			req:         nil,
			wantSuccess: false,
		},
		{
			name: "正常请求",
			req: &dto.DeleteConnectionRequest{
				Id: "test-id",
			},
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := api.DeleteConnection(tt.req)
			if resp.Success != tt.wantSuccess {
				t.Errorf("DeleteConnection() success = %v, want %v", resp.Success, tt.wantSuccess)
			}
		})
	}
}

func TestConnectionApi_TestConnection(t *testing.T) {
	api := NewConnectionApi()

	tests := []struct {
		name        string
		req         *dto.TestConnectionRequest
		wantSuccess bool
	}{
		// {
		// 	name:        "请求为空",
		// 	req:         nil,
		// 	wantSuccess: false,
		// },
		// {
		// 	name: "正常请求",
		// 	req: &dto.TestConnectionRequest{
		// 		Connection: model.Connection{
		// 			Name: "test-conn",
		// 			Host: "localhost",
		// 			Port: 2379,
		// 		},
		// 	},
		// 	wantSuccess: true,
		// },
		{
			name: "错误请求",
			req: &dto.TestConnectionRequest{
				Connection: model.Connection{
					Name: "test-conn",
					Host: "localhostxxx",
					Port: 2379,
				},
			},
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := api.TestConnection(tt.req)
			if resp.Success != tt.wantSuccess {
				t.Errorf("TestConnection() success = %v, want %v", resp.Success, tt.wantSuccess)
			}
		})
	}
}
