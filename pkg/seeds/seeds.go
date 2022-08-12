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
			Run: func(db *gorm.DB) error {
				for i := 0; i < 30; i++ {
					CreateUser(db, gofakeit.Name(), gofakeit.Phone(), gofakeit.Email())
				}
				return nil
			},
		},
	}
}
