package entity

type OrderRepositoryInterface interface {
	ListOrder() ([]Order, error)
	Save(order *Order) error
	GetTotal() (int, error)
}
