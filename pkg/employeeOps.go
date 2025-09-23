package pkg

import (
	"attandance/models"
	"attandance/utils"
	"encoding/json"
	"fmt"
	"os"
)

/*
Adds employee to file.
*/
func AddEmployee(emp *models.Employee) error {
	rawContents, err := utils.FileReadOpration()
	if err != nil {
		return fmt.Errorf("error while reading employee file, err : %s", err.Error())
	}
	emps := make([]models.Employee, 0)
	err = json.Unmarshal(rawContents, &emps)
	if err != nil {
		return fmt.Errorf("error while unmarshalling employee file contents, err : %s", err.Error())
	}
	emps = append(emps, *emp)

	fc, err := json.Marshal(emps)
	if err != nil {
		return fmt.Errorf("error while marshalling employee file contents, err : %s", err.Error())
	}
	err = os.WriteFile(utils.EmployeeFile, fc, os.FileMode(0666))
	if err != nil {
		return fmt.Errorf("error while writing employee file contents, err : %s", err.Error())
	}
	return nil
}
func FindEmployee(id string) (models.Employee, error) {
	rawContents, err := utils.FileReadOpration()
	if err != nil {
		return models.Employee{}, fmt.Errorf("error while reading employee file: %s", err.Error())
	}

	emps := make([]models.Employee, 0)
	err = json.Unmarshal(rawContents, &emps)
	if err != nil {
		return models.Employee{}, fmt.Errorf("error while unmarshalling employee file contents: %s", err.Error())
	}

	for _, emp := range emps {
		if emp.ID == id {
			return emp, nil
		}
	}

	return models.Employee{}, fmt.Errorf("employee with id %s not found", id)
}
