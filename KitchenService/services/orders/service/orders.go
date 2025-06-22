package service

// This is going to be business logic. No need to depend on http or grpc

type OrderService struct {
	// store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
