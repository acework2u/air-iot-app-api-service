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

//func (r *acErrorRepo) GetErrorCode(code int) (*AcErrorCodeInfo, error) {
//
//	fmt.Println("IN Repo code id", code)
//	query := "select code_id,unit,title from error_code where code_id=?"
//	row := r.db.QueryRow(query, code)
//	//defer r.db.Close()
//
//	acErr := AcErrorCodeInfo{}
//	err := row.Scan(&acErr.ID, &acErr.Unit, &acErr.Title)
//	if err != nil {
//		return nil, err
//	}
//	return &acErr, err
//
//}
