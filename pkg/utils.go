package pkg

type LimitOffset struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

func Paginate(limit, offset string) (*LimitOffset, error) {
	var pagination *LimitOffset
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}

	pagination = &LimitOffset{
		Limit:  limit,
		Offset: offset,
	}

	return pagination, nil
}
