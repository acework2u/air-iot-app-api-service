package smartapp

import "github.com/uptrace/bun"

type AcErrorInfo struct {
	bun.BaseModel `bun:"table:error_code,alias:er"`

	ID    int64  `bun:"id,pk" json:"id"`
	Unit  string `bun:"unit" json:"unit"`
	Title string `bun:"title" json:"title"`
}
type APIErrorCode struct {
	Code   string `json:"code"`
	Unit   string `json:"unit"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Video  string `json:"video"`
	Web    string `json:"web"`
}

type AcErrorService interface {
	GetErrorByCode(code int) (*APIErrorCode, error)
	GetErrors() ([]APIErrorCode, error)
}

type ErrorCodeService interface {
	ErrorCodeList() ([]*APIErrorCode, error)
	ErrorByCode(code string) (*APIErrorCode, error)
}
