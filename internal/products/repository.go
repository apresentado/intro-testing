package products

import (
	"context"
	"errors"
)
var(
	errNotFound = errors.New("product not found")
)

// repository provides access to the product repository

// solo tiene un metodo y nos retorna un producto en base a su id

type Repository interface {
	GetByID( ctx context.Context ,id int) (Products, error)
}

// SliceRerpository is a in-memory repository

type SliceRepository struct {
	products []Products
}

func NewSliceRepository(Products []Products) *SliceRepository {
	return &SliceRepository{ // retorna la instancia
		products: Products,
	}
}

func (r *SliceRepository ) GetByID( ctx context.Context ,id int) (product Products, err error) {
	// Esto define un método llamado GetByID en un tipo SliceRepository que implementa la interfaz Repository.
	//iterar sobre los productos y retornar el que tenga el id que estamos buscando
	for index := range r.products {
		if r.products[index].ID == id {
			return r.products[index], nil
		}
	}
	//el producto no fue encontrado
	err = errNotFound
	return 
	
}
