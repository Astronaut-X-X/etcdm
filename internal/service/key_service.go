package service

import (
	"etcdm/internal/client"
	"etcdm/internal/dto"
	"etcdm/internal/model"
	"strings"
)

type KeyService struct {
}

func NewKeyService() *KeyService {
	return &KeyService{}
}

func (s *KeyService) GetKeysTree(conn *model.Connection, req *dto.GetKeysTreeRequest) ([]*dto.KeyTreeItem, error) {
	resp, err := client.Execute(conn, client.GetAllKey, client.GetAllKeyParams{Keyword: req.Keyword})
	if err != nil {
		return nil, err
	}

	res := make([]*dto.KeyTreeItem, 0)
	map_ := make(map[string]*dto.KeyTreeItem)
	buildParentItem := func(parentKey string) *dto.KeyTreeItem {
		words := strings.Split(parentKey, "/")
		var parentItem *dto.KeyTreeItem
		var item *dto.KeyTreeItem
		var isExist bool
		for i, word := range words {
			key := strings.Join(words[:i+1], "/")
			item, isExist = map_[key]
			if isExist {
				parentItem = item
				continue
			}
			item = &dto.KeyTreeItem{
				Key:      strings.Join(words[:i+1], "/"),
				Label:    word,
				Children: nil,
			}
			map_[item.Key] = item
			if i == 0 {
				res = append(res, item)
				parentItem = item
			} else {
				parentItem.Children = append(parentItem.Children, item)
			}
		}
		return item
	}

	for _, key := range resp {
		words := strings.Split(key, "/")
		parentKey := strings.Join(words[:len(words)-1], "/")
		parentItem, isExist := map_[parentKey]
		if !isExist {
			parentItem = buildParentItem(parentKey)
		}
		if !req.WithSuffix {
			continue
		}
		item := &dto.KeyTreeItem{
			Key:   key,
			Label: words[len(words)-1],
		}
		parentItem.Children = append(parentItem.Children, item)
	}

	return res, nil
}
