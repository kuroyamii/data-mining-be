package classificationService

import "context"

type ClassificationService interface {
	InsertImagePath(ctx context.Context, imagePath string, plant string, condition string, disease string) error
}
