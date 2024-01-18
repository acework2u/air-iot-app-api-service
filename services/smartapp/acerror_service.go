package smartapp

import (
	"fmt"
	acrepo "github.com/acework2u/air-iot-app-api-service/repository/smartapp"
)

type acErrorService struct {
	acErrRepo acrepo.AcErrorCodeRepo
}

func NewAcErrorService(acErrRepo acrepo.AcErrorCodeRepo) AcErrorService {
	return &acErrorService{acErrRepo}
}

func (s *acErrorService) GetErrorByCode(code int) (*AcErrorInfo, error) {

	errInfo := AcErrorInfo{}
	rows, err := s.acErrRepo.GetErrorCode(code)
	if err != nil {
		return nil, err
	}
	errInfo = AcErrorInfo{ID: rows.ID, Unit: rows.Unit, Title: rows.Title}
	fmt.Println(rows)

	return &errInfo, nil
}
