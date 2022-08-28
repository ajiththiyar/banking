package domain

type Customer struct {
	Id          string
	City        string
	Name        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
