package api

import (
	"testing"

	"etcdm/internal/dto"
)

// 测试GetKeysTree成功场景
func TestKeyApi_GetKeysTree(t *testing.T) {
	api := NewKeyApi()

	tests := []struct {
		name        string
		req         *dto.GetKeysTreeRequest
		wantData    bool
		wantSuccess bool
	}{
		{
			name: "正常请求",
			req: &dto.GetKeysTreeRequest{
				ID: "8e97833d-0475-4cc6-9418-04a6ef5c1ecb",
			},
			wantData:    true,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := api.GetKeysTree(tt.req)
			if resp.Success != tt.wantSuccess {
				t.Errorf("GetKeysTree() success = %v, want %v", resp.Success, tt.wantSuccess)
			}
			if (resp.Data != nil) != tt.wantData {
				t.Errorf("GetKeysTree() data presence = %v, want %v", resp.Data != nil, tt.wantData)
			}
		})
	}
}
