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
	//rows, err := s.acErrRepo.GetErrorCode(code)
	//if err != nil {
	//	return nil, errors.New("Error code ไม่ถูกต้อง หรือ ไม่มีข้อมูลนี้ในระบบ")
	//}
	//errInfo = APIErrorCode{
	//	Code:   rows.Code,
	//	Unit:   rows.Unit,
	//	Title:  rows.Title,
	//	Detail: rows.Detail,
	//	Video:  rows.UrlVideo,
	//	Web:    rows.UrlWeb,
	//}

	return &errInfo, nil
}

func (s *acErrorService) GetErrors() ([]APIErrorCode, error) {

	//rows, err := s.acErrRepo.AcErrorCodeList()
	//if err != nil {
	//	return nil, err
	//}
	//res := []APIErrorCode{}
	//for _, items := range rows {
	//	item := APIErrorCode{
	//		Code:   items.Code,
	//		Unit:   items.Unit,
	//		Title:  items.Title,
	//		Detail: items.Detail,
	//		Video:  items.UrlVideo,
	//		Web:    items.UrlWeb,
	//	}
	//	res = append(res, item)
	//}
	//return res, nil
	return nil, nil
}

type errCodeService struct {
	errCodeRepo acrepo.ErrorCodeRepo
}

func (s errCodeService) ErrorCodeList() ([]*APIErrorCode, error) {

	res, err := s.errCodeRepo.ErrorCodeList()
	if err != nil {
		return nil, err
	}
	errCodeList := []*APIErrorCode{}
	for _, items := range res {
		item := &APIErrorCode{
			Code:   items.Code,
			Unit:   items.Unit,
			Title:  items.Title,
			Detail: items.Detail,
			Video:  items.VideoUrl,
			Web:    items.WebUrl,
		}

		errCodeList = append(errCodeList, item)

	}
	return errCodeList, nil
}
func (s errCodeService) ErrorByCode(code string) (*APIErrorCode, error) {

	res, err := s.errCodeRepo.AcErrCode(code)
	if err != nil {
		return nil, errors.New("error code is wrong or no data")
	}
	acCodeErr := &APIErrorCode{
		Code:   res.Code,
		Title:  res.Title,
		Detail: res.Detail,
		Unit:   res.Unit,
		Video:  res.VideoUrl,
		Web:    res.WebUrl,
	}

	return acCodeErr, nil
}

func NewErrorCodeService(errCodeRepo acrepo.ErrorCodeRepo) ErrorCodeService {
	return &errCodeService{errCodeRepo: errCodeRepo}
}
