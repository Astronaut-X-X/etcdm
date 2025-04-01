package pkg

type PageHelper[T any] struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []T `json:"data"`
}

func NewPageHelper[T any](data []T, page, Size int) *PageHelper[T] {
	return &PageHelper[T]{
		Page:  page,
		Size:  Size,
		Total: len(data),
		Data:  data,
	}
}

func (p *PageHelper[T]) GetPage() ([]T, int64) {
	// 计算起始索引
	start := (p.Page - 1) * p.Size
	if start >= p.Total {
		return []T{}, int64(p.Total)
	}

	// 计算结束索引
	end := start + p.Size
	if end > p.Total {
		end = p.Total
	}

	return p.Data[start:end], int64(p.Total)
}
