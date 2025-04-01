package db

import "etcdm/internal/model"

type IDB[T model.IModel] interface {
	Conn() string
	Get(id string) (T, error)
	List() ([]T, error)
	Create(t T) error
	Update(t T) error
	Delete(id string) error
}
