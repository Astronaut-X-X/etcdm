package db

import (
	"encoding/json"
	"etcdm/internal/model"
	"etcdm/internal/pkg"
	"fmt"
	"os"
)

func NewFileDB[T model.IModel](fileName string) IDB[T] {
	return &FileDB[T]{
		FileName: fileName,
	}
}

var _ IDB[model.IModel] = FileDB[model.IModel]{}

type FileDB[T model.IModel] struct {
	FileName string
}

func (b FileDB[T]) Conn() string {
	return b.FileName
}

func (b FileDB[T]) Get(id string) (T, error) {
	var t T

	result, err := b.List()
	if err != nil {
		return t, err
	}

	for _, v := range result {
		if v.GetId() == id {
			return v, nil
		}
	}

	return t, fmt.Errorf("%s not exists", id)
}

func (b FileDB[T]) List() ([]T, error) {
	var result []T

	data, err := pkg.ReadFile(b.Conn())
	if err != nil {
		if !os.IsNotExist(err) {
			return result, err
		}
		data = []byte("6CixH7PPkF491pu827gQTw==")
	}

	decrypted, err := pkg.Decrypt(string(data))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal([]byte(decrypted), &result); err != nil {
		return result, err
	}

	return result, nil
}

func (b FileDB[T]) Create(t T) error {

	result, err := b.List()
	if err != nil {
		return err
	}
	result = append(result, t)

	return b.save(result)
}

func (b FileDB[T]) Update(t T) error {
	result, err := b.List()
	if err != nil {
		return err
	}

	for i, v := range result {
		if v.GetId() == t.GetId() {
			result[i] = t
		}
	}

	return b.save(result)
}

func (b FileDB[T]) Delete(id string) error {
	result, err := b.List()
	if err != nil {
		return err
	}

	for i, v := range result {
		if v.GetId() == id {
			result = append(result[:i], result[i+1:]...)
			break
		}
	}

	return b.save(result)
}

func (b FileDB[T]) save(r []T) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	encrypted, err := pkg.Encrypt(string(data))
	if err != nil {
		return err
	}

	if err := pkg.WriteFile(b.Conn(), []byte(encrypted)); err != nil {
		return err
	}

	return nil
}
