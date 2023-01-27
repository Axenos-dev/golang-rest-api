package models

type SearchForUserByEmail struct {
	Email string
}

func (query *SearchForUserByEmail) CreateQuery(email string) {
	query.Email = email
}
