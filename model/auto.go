package model

type AutoStatus string // Статус автомобиля
const (
	TRANSIT      AutoStatus = "transit"      // В пути
	STOCK        AutoStatus = "stock"        // На складе
	SOLD_OUT     AutoStatus = "sold out"     // Продан
	DISCONTINUED AutoStatus = "discontinued" // Снят с продажи
)

type Auto struct {
	ID        uint       `json:"id"`        // Уникальный идентификатор
	Brandname string     `json:"brandname"` // Бренд автомобиля
	Model     string     `json:"model"`     // Модель автомобиля
	Price     uint       `json:"price"`     // Цена автомобиля
	Status    AutoStatus `json:"status"`    // Статус автомобиля
	Mileage   uint       `json:"mileage"`   // Пробег автомобиля
}
