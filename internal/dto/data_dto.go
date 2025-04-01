package dto

type ListDataRequest struct {
	PageRequest
	Params ListDataParams `json:"params"`
}

type ListDataParams struct {
	ID     string `json:"id"`
	Prefix string `json:"prefix"`
}

type DataItem struct {
	Key            string `json:"key,omitempty"`
	CreateRevision int64  `json:"create_revision,omitempty"`
	ModRevision    int64  `json:"mod_revision,omitempty"`
	Version        int64  ` json:"version,omitempty"`
	Value          string ` json:"value,omitempty"`
	Lease          int64  ` json:"lease,omitempty"`
}
