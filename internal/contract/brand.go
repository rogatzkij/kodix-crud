package contract

type Brand interface {
	CreateBrand(brand string) error
	CheckBrand(brand string) (bool, error)
	CreateModel(brand, model string) error
	CheckModel(brand, model string) (bool, error)
}
