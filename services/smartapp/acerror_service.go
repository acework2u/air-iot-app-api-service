package smartapp

import (
	"errors"
	acrepo "github.com/acework2u/air-iot-app-api-service/repository/smartapp"
)

type acErrorService struct {
	acErrRepo acrepo.AcErrorCodeRepo
}

func NewAcErrorService(acErrRepo acrepo.AcErrorCodeRepo) AcErrorService {
	return &acErrorService{acErrRepo}
}

func (s *acErrorService) GetErrorByCode(code int) (*APIErrorCode, error) {

	errInfo := APIErrorCode{}
	rows, err := s.acErrRepo.GetErrorCode(code)
	if err != nil {
		return nil, errors.New("Error code ไม่ถูกต้อง หรือ ไม่มีข้อมูลนี้ในระบบ")
	}
	errInfo = APIErrorCode{
		Code:   rows.Code,
		Unit:   rows.Unit,
		Title:  rows.Title,
		Detail: rows.Detail,
		Video:  rows.UrlVideo,
		Web:    rows.UrlWeb,
	}

	return &errInfo, nil
}
