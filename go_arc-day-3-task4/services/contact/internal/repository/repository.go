package repository

import "architecture_go/services/contact/internal/domain"

type IRepository interface {
    NewContact(id int, fullName string, phoneNumber string)

    GetContact(id int) (domain.Contact)

	SetFullName(id int, lastName, firstName, middleName string)

	GetFullName() string

	SetPhoneNumber(id int, phoneNumber string)
}
