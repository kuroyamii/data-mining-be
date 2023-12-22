package classificationRepository

import (
	"context"
	"database/sql"
	"datamining-be/pkg/utilities"
	"log"
)

type classificationRepository struct {
	db *sql.DB
}

const GET_DISEASE_DETAIL = `
SELECT pdd.id FROM plant_disease_details pdd 
INNER JOIN plants p ON p.id = pdd.plant_id
INNER JOIN diseases d ON d.id = pdd.disease_id 
WHERE p.name=? AND d.name=?
;
`

const GET_CONDITION_ID = `
SELECT id FROM conditions
WHERE name=?
`
const INSERT_IMAGE_PATH = `
INSERT INTO plant_images(file_name, plant_disease_detail_id,condition_id) VALUES
(?,?,?)
`

func (cr classificationRepository) GetDiseaseDetailByPlantAndDisease(ctx context.Context, plant string, disease_unsplitted string) (int, error) {
	result, err := cr.db.QueryContext(ctx, GET_DISEASE_DETAIL, plant, disease_unsplitted)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return 0, err
	}
	var res int

	for result.Next() {
		err = result.Scan(&res)
	}
	return res, nil

}
func (cr classificationRepository) GetConditionId(ctx context.Context, condition string) (int, error) {
	result, err := cr.db.QueryContext(ctx, GET_CONDITION_ID, condition)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return 0, err
	}
	var res int

	for result.Next() {
		err = result.Scan(&res)
		if err != nil {
			log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
			return 0, err
		}
	}
	return res, nil
}

func (cr classificationRepository) InsertImagePath(ctx context.Context, diseaseDetailID int, conditionID int, image_path string) error {
	stmt, err := cr.db.PrepareContext(ctx, INSERT_IMAGE_PATH)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return err
	}
	var res sql.Result
	res, err = stmt.ExecContext(ctx, image_path, diseaseDetailID, conditionID)
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		log.Printf("%v -> %v", utilities.Red("ERROR"), err.Error())
		return err
	}

	return nil
}

func NewClassificationRepository(db *sql.DB) classificationRepository {
	return classificationRepository{
		db: db,
	}
}
