package dto

import "etcdm/internal/model"

type GetConnectionRequest struct {
	ID string `json:"id"`
}

type ListConnectionRequest struct {
	PageRequest
}

type CreateConnectionRequest struct {
	model.Connection
}

type UpdateConnectionRequest struct {
	model.Connection
}

type DeleteConnectionRequest struct {
	Id string `json:"id" binding:"required"`
}

type TestConnectionRequest struct {
	model.Connection
}
