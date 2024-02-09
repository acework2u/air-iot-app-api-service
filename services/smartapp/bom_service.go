package smartapp

import (
	"fmt"
	bomRepo "github.com/acework2u/air-iot-app-api-service/repository/smartapp"
)

type bomService struct {
	bomRepo bomRepo.BomRepository
}

func NewBomService(bomRepo bomRepo.BomRepository) BomService {
	return &bomService{bomRepo: bomRepo}
}

func (s bomService) CheckCompressor(indoor string) ([]*AcBomResponse, error) {

	comRes, err := s.bomRepo.Compressor(indoor)
	//comRes, err := s.bomRepo.Compressors()
	if err != nil {
		return nil, err
	}
	compressor := []*AcBomResponse{}

	for _, items := range comRes {
		yearTxt := fmt.Sprintf("%v", items.Year)
		item := &AcBomResponse{
			Year:     yearTxt,
			Btu:      items.Btu,
			IndItem:  items.IndItem,
			OduItem:  items.OduItem,
			IndModel: items.IndModel,
			OduModel: items.OduModel,
			Compressors: Compressor{
				Brand:  items.Compressors.Brand,
				Model:  items.Compressors.Model,
				ItemNo: items.Compressors.ItemNo,
			},
		}

		compressor = append(compressor, item)

	}

	return compressor, nil
}

func (s *bomService) CompressorList() ([]*AcBomResponse, error) {

	compList, err := s.bomRepo.Compressors()
	if err != nil {
		return nil, err
	}
	var acCompList []*AcBomResponse
	for _, items := range compList {
		item := &AcBomResponse{
			Year:     string(items.Year),
			IndItem:  items.IndItem,
			IndModel: items.IndModel,
		}
		acCompList = append(acCompList, item)
	}

	return acCompList, nil
}
