package gates

import "gorm.io/gorm"

type Orm interface {
	Connect()
	GetConnection() *gorm.DB
	Migrate() error
}
