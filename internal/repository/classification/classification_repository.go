package classificationRepository

import "context"

type ClassificationRepository interface {
	GetDiseaseDetailByPlantAndDisease(ctx context.Context, plant string, disease_unsplitted string) (int, error)
	GetConditionId(ctx context.Context, condition string) (int, error)
	InsertImagePath(ctx context.Context, diseaseDetailID int, conditionID int, image_path string) error
}
