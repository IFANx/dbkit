package common

type Provider interface {
	GenerateDatabase()
	GenerateTable() Table
}
