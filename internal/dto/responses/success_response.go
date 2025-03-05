package responses

type successResponse[T any] struct {
	Data T `json:"data"`
}

func NewSuccessResponse[T any](data T) *successResponse[T] {
	return &successResponse[T]{Data: data}
}
