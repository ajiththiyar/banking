package domain

type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"100", "Akash", "New Delhi", "482929", "3999-12-1", "1"},
		{"101", "Chintu", "New Delhi", "482529", "2000-12-1", "1"},
	}
	return CustomerRepositoryStub{customers}
}
