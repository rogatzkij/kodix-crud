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
	ID        uint       `json:"id" bson:"id" jsonapi:"primary,autos"`                 // Уникальный идентификатор
	Brandname string     `json:"brandname" bson:"brandname" jsonapi:"attr, brandname"` // Бренд автомобиля
	Automodel string     `json:"automodel" bson:"automodel" jsonapi:"attr, automodel"` // Модель автомобиля
	Price     uint       `json:"price" bson:"price" jsonapi:"attr, price"`             // Цена автомобиля
	Status    AutoStatus `json:"status" bson:"status" jsonapi:"attr, status"`          // Статус автомобиля
	Mileage   uint       `json:"mileage" bson:"mileage" jsonapi:"attr, mileage"`       // Пробег автомобиля
}
