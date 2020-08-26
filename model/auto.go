package model

type AutoStatus string // Статус автомобиля
const (
	TRANSIT      AutoStatus = "transit"      // В пути
	STOCK        AutoStatus = "stock"        // На складе
	SOLD_OUT     AutoStatus = "sold out"     // Продан
	DISCONTINUED AutoStatus = "discontinued" // Снят с продажи
)

func (as AutoStatus) Check() bool {
	switch as {
	case TRANSIT, STOCK, SOLD_OUT, DISCONTINUED:
		return true
	default:
		return false
	}
}

type Auto struct {
	ID        uint       `json:"id" bson:"id"`               // Уникальный идентификатор
	Brandname string     `json:"brandname" bson:"brandname"` // Бренд автомобиля
	Automodel string     `json:"automodel" bson:"automodel"` // Модель автомобиля
	Price     uint       `json:"price" bson:"price"`         // Цена автомобиля
	Status    AutoStatus `json:"status" bson:"status"`       // Статус автомобиля
	Mileage   uint       `json:"mileage" bson:"mileage"`     // Пробег автомобиля
}
