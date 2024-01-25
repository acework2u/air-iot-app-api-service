package smartapp

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

//type acErrorRepo struct {
//	db  *sql.DB
//	ctx context.Context
//	db2 *bun.DB
//	db3 *gorm.DB
//}

type acErrorRepo struct {
	db  *gorm.DB
	ctx context.Context
}

type DbType interface {
	*sql.DB | *bun.DB | *gorm.DB
}

func NewAcErrorCodeRepo(Db *gorm.DB) AcErrorCodeRepo {

	return &acErrorRepo{db: Db}
}

func (r *acErrorRepo) GetErrorCode(code int) (*ErrorCode, error) {

	acErr := ErrorCode{}
	res := r.db.First(&acErr, "code = ?", code)
	if res.RowsAffected > 0 {
		return &acErr, nil
	}
	return nil, res.Error
}

func (r *acErrorRepo) AcErrorCodeList() ([]ErrorCode, error) {
	acErrors := []ErrorCode{}
	res := r.db.Find(&acErrors)
	if res.RowsAffected > 0 {
		return acErrors, nil
	}
	return nil, res.Error
}
