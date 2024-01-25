package smartapp

import (
	"github.com/uptrace/bun"
)

type AcErrorCodeInfo struct {
	ID    int64  `bun:"id,pk" json:"id"`
	Unit  string `bun:"unit" json:"unit"`
	Title string `bun:"title" json:"title"`
}

type ErrorCode2 struct {
	bun.BaseModel `bun:"table:error_code,alias:er"`

	ID    int64  `bun:"id,pk" json:"id"`
	Unit  string `bun:"unit" json:"unit"`
	Title string `bun:"title" json:"title"`
}

type ErrorCode struct {
	//gorm.Model
	Id       uint `gorm:primary`
	Code     int64
	Unit     string
	Title    string
	Detail   string
	UrlVideo string
	UrlWeb   string
}

type AcErrorCodeRepo interface {
	GetErrorCode(code int) (*ErrorCode, error)
}
