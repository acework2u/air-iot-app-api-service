package smartapp

import (
	"errors"
	smartappRep "github.com/acework2u/air-iot-app-api-service/repository/smartapp"
)

type diagnosticBoardService struct {
	diagnosticRepo smartappRep.DiagnosticBoardRepo
}

func NewDiagnosticService(diagnosticRepo smartappRep.DiagnosticBoardRepo) DiagnosticService {
	return &diagnosticBoardService{diagnosticRepo: diagnosticRepo}
}

func (s *diagnosticBoardService) CheckDiagnosticBoard(filter2 *DiagnosticFilter) (*DiagnosticResponse, error) {
	filter := smartappRep.DiagnosticFilter{
		Btu:    filter2.Btu,
		CompId: filter2.CompId,
	}

	res, err := s.diagnosticRepo.DiagnosticBoard(filter)
	if err != nil {
		return nil, errors.New("ไม่พบข้อมูล ในระบบ")
	}

	diagResponse := &DiagnosticResponse{
		Btu:       res.Btu,
		CompId:    res.CompId,
		CompItem:  res.CompItem,
		CompModel: res.CompModel,
	}

	return diagResponse, nil
}
func (s *diagnosticBoardService) DiagnosticBoards() ([]*DiagnosticResponse, error) {
	res, err := s.diagnosticRepo.DiagnosticBoards()
	if err != nil {
		return nil, err
	}
	compList := []*DiagnosticResponse{}
	for _, items := range res {
		item := &DiagnosticResponse{
			Btu:       items.Btu,
			CompId:    items.CompId,
			CompItem:  items.CompItem,
			CompModel: items.CompModel,
		}
		compList = append(compList, item)
	}

	return compList, nil
}
