package dto

type GetKeysTreeRequest struct {
	ID         string `json:"id"`
	Keyword    string `json:"keyword"`
	WithSuffix bool   `json:"with_suffix"`
}

type KeyTreeItem struct {
	Label    string         `json:"label"`
	Key      string         `json:"key"`
	Children []*KeyTreeItem `json:"children"`
}
