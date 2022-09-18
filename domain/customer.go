package domain

import "github.com/ajiththiyar/banking/errs"

type Customer struct {
	Id          string
	City        string
	Name        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
