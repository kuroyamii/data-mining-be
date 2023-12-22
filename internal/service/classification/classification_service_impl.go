package classificationService

import (
	"context"
	classificationRepository "datamining-be/internal/repository/classification"
	"datamining-be/pkg/utilities"
	"log"
)

type classificationService struct {
	cr classificationRepository.ClassificationRepository
}

func (cs classificationService) InsertImagePath(ctx context.Context, imagePath string, plant string, condition string, disease string) error {
	res, err := cs.cr.GetDiseaseDetailByPlantAndDisease(ctx, plant, disease)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return err
	}

	condid, err := cs.cr.GetConditionId(ctx, condition)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return err
	}

	err = cs.cr.InsertImagePath(ctx, res, condid, imagePath)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return err
	}

	return nil
}

func NewClassificationService(cr classificationRepository.ClassificationRepository) classificationService {
	return classificationService{
		cr: cr,
	}
}
