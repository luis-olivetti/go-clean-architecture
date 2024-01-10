package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type GetAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	Event           events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetAllOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	Event events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{
		OrderRepository: OrderRepository,
		Event:           Event,
		EventDispatcher: EventDispatcher,
	}
}

func (c *GetAllOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var output []OrderOutputDTO
	for _, order := range orders {
		output = append(output, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	c.Event.SetPayload(output)
	c.EventDispatcher.Dispatch(c.Event)

	return output, nil
}
