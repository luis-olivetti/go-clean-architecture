package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetAll() ([]*Order, error)
}
