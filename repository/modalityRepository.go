package repository

import (
	"log"

	"github.com/GustavoFreitas22/ludus_sys/config"
	"github.com/GustavoFreitas22/ludus_sys/model"
	"github.com/GustavoFreitas22/ludus_sys/query"
)

func CreateNewModality(name string) {

	statement, err := config.Datasource.Prepare(query.INSERT_MODALITY)
	checkErr(err)

	result, err := statement.Exec(name)
	checkErr(err)

	linesAffect, err := result.RowsAffected()
	checkErr(err)

	log.Println("Rows affects: ", linesAffect)
}

func FindAllModalities() []model.Modality {

	var allModalities []model.Modality

	statement, err := config.Datasource.Query(query.SELECT_ALL_MODALITIES)
	checkErr(err)

	for statement.Next() {
		var modality model.Modality

		err = statement.Scan(&modality.Id, &modality.Name)
		checkErr(err)

		allModalities = append(allModalities, modality)
	}

	return allModalities
}

func FindModalityById(id int) model.Modality {

	var modality model.Modality

	err := config.Datasource.QueryRow(query.SELECT_MODALITY_BY_ID, id).Scan(&modality.Name)
	checkErr(err)

	return modality
}

func UpdateModality(id int, name string) {
	statement, err := config.Datasource.Prepare(query.UPDATE_MODALITY)
	checkErr(err)

	result, err := statement.Exec(name, id)
	checkErr(err)

	linesAffect, err := result.RowsAffected()
	checkErr(err)

	log.Println("Rows affects: ", linesAffect)

}

func DeleteModality(id int) {
	statement, err := config.Datasource.Prepare(query.DELETE_MODALITY)
	checkErr(err)

	result, err := statement.Exec(id)
	checkErr(err)

	linesAffect, err := result.RowsAffected()
	checkErr(err)

	log.Println("Rows affects: ", linesAffect)
}

func checkErr(err error) {
	if err != nil {
		log.Println("Error to execute statemnet: ", err)
	}
}
