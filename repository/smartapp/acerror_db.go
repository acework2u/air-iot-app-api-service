package smartapp

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

type acErrorRepo struct {
	db  *sql.DB
	ctx context.Context
	db2 *bun.DB
}

func NewAcErrorCodeRepo(db *sql.DB) AcErrorCodeRepo {

	return &acErrorRepo{db: db}
}

func (r *acErrorRepo) GetErrorCode(code int) (*AcErrorCodeInfo, error) {

	fmt.Println("IN Repo code id", code)
	query := "select code_id,unit,title from error_code where code_id=?"
	row := r.db.QueryRow(query, code)
	//defer r.db.Close()

	acErr := AcErrorCodeInfo{}
	err := row.Scan(&acErr.ID, &acErr.Unit, &acErr.Title)
	if err != nil {
		return nil, err
	}
	return &acErr, err

	//acErr := AcErrorCodeInfo{}
	//err := r.db.NewSelect().Table("error_code").Where("code_id = ?", code).Scan(r.ctx, &acErr)
	//if err != nil {
	//	return nil, err
	//}
	//r.db.Ping()
	//if err := r.db.Ping(); err != nil {
	//	log.Fatal(err)
	//}

	//err := r.db.NewSelect().Model(&ErrorCode)
	//rows, err := r.db.QueryContext(r.ctx, "SELECT * FROM error_code where code_id =2")
	//if err != nil {
	//	panic(err)
	//}
	//acErrs := []ErrorCode{}
	//err = r.db.ScanRows(r.ctx, rows, &acErrs)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(acErrs)

	//return nil, nil
}
