package products

import "context"

//service es una logica ded negocio que se encarga de manejar los productos y su registro
type Service interface {
	//GetByID returns a product by its ID
	GetByID(ctx context.Context, id int) (Products, error)
}
//DefaultService is a default implementation of the Service interface

type DefaultService struct {
	// repository is used to access the product repository
	repository Repository
}

// NewDefaultService creates a new instance of the DefaultService

func NewDefaultService(repository Repository) *DefaultService {
	return &DefaultService{
		repository: repository,
	}
}

// GetByID returns a product by its ID

func (s *DefaultService) GetByID(ctx context.Context, id int) (Products, error) {
	return s.repository.GetByID(ctx,id) // se lo envio al repository para que cumpla con la funcion 
}