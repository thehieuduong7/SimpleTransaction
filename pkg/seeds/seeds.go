// pkg/seeds/seeds.go
package seeds

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"internBE.com/pkg/seed"
)

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateUser",
			Run: func(dbConn *gorm.DB) error {
				for i := 0; i < 10; i++ {
					CreateUser(dbConn, gofakeit.Name(), gofakeit.Phone(), gofakeit.Email())
				}
				return nil
			},
		},
		seed.Seed{
			Name: "CreateAccount",
			Run: func(dbConn *gorm.DB) error {
				for i := 0; i < 20; i++ {
					CreateAccount(dbConn, gofakeit.Number(1, 10), gofakeit.Price(10000, 500000))
				}
				return nil
			},
		},
		seed.Seed{
			Name: "CreateTransaction",
			Run: func(dbConn *gorm.DB) error {
				for i := 0; i < 20; i++ {
					CreatTransaction(dbConn, gofakeit.Number(1, 10), gofakeit.Number(1, 10), gofakeit.Company(), gofakeit.Price(10000, 500000))
				}
				return nil
			},
		},
	}
}
