package vmysql

import "gorm.io/gorm"

type (
	Config = gorm.Config
	DB     = gorm.DB
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

func Open(dialector gorm.Dialector, opts ...gorm.Option) (db *DB, err error) {
	return gorm.Open(dialector, opts...)
}
