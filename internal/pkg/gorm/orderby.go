package postgresgorm

import "fmt"

type OrderBy struct {
	Field     string           `json:"field"`
	Direction OrderByDirection `json:"direction"`
}

func (o *OrderBy) ToSort() string {
	return fmt.Sprintf("%s %s", o.Field, o.Direction)
}

type OrderByDirection string

const (
	OrderByDirectionAsc  OrderByDirection = "asc"
	OrderByDirectionDesc OrderByDirection = "desc"
)

type OrderByDirectionBsonValue int

const (
	OrderByDirectionBsonValueAsc  OrderByDirectionBsonValue = 1
	OrderByDirectionBsonValueDesc OrderByDirectionBsonValue = -1
)

func (o *OrderByDirection) GetSortDirectionBsonValue() OrderByDirectionBsonValue {
	switch *o {
	case OrderByDirectionAsc:
		return OrderByDirectionBsonValueAsc
	case OrderByDirectionDesc:
		return OrderByDirectionBsonValueDesc
	}
	return OrderByDirectionBsonValueAsc
}
