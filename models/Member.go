package models

type Member struct {
	ID    int
	Name  string
	Email string
}

var Members = []Member{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	{ID: 3, Name: "Sam Wilson", Email: "sam@example.com"},
}
