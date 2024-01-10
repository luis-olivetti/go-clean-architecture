package event

import "time"

type GetAllOrders struct {
	Name    string
	Payload interface{}
}

func NewGetAllOrders() *GetAllOrders {
	return &GetAllOrders{
		Name: "GetAllOrders",
	}
}

func (e *GetAllOrders) GetName() string {
	return e.Name
}

func (e *GetAllOrders) GetPayload() interface{} {
	return e.Payload
}

func (e *GetAllOrders) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *GetAllOrders) GetDateTime() time.Time {
	return time.Now()
}
