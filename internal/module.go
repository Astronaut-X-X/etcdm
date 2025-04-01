package internal

import "etcdm/internal/api"

func BindApi() []interface{} {
	return []interface{}{
		api.NewConnectionApi(),
		api.NewKeyApi(),
		api.NewDataApi(),
	}
}
