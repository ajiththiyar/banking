package domain

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ajiththiyar/banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (cd CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findCustomer := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := cd.client.QueryRow(findCustomer, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {

			fmt.Println("Error While scanning the table", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database Error")
		}
	}
	return &c, nil
}

func (cd CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := cd.client.Query(findAllSql)
	if err != nil {
		fmt.Println("Error Querying the Table", err.Error())
		return nil, errs.NewNotFoundError("Issue Querying")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customer Not Found")
			} else {

				fmt.Println("Error While scanning the table", err.Error())
				return nil, errs.NewUnexpectedError("Unexpected Database Error")
			}
		}
		customers = append(customers, c)
	}

	return customers, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
