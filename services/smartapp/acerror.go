package smartapp

import "github.com/uptrace/bun"

type AcErrorInfo struct {
	bun.BaseModel `bun:"table:error_code,alias:er"`

	ID    int64  `bun:"id,pk" json:"id"`
	Unit  string `bun:"unit" json:"unit"`
	Title string `bun:"title" json:"title"`
}
type APIErrorCode struct {
	Code   int64
	Unit   string
	Title  string
	Detail string
	Video  string
	Web    string
}

type AcErrorService interface {
	GetErrorByCode(code int) (*APIErrorCode, error)
	GetErrors() ([]APIErrorCode, error)
}
