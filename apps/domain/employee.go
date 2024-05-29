package domain

type Employee struct {
	EmployeeCode  string `json:"employee_code"`
	DepartementId int    `json:"departement_id"`
	PositionId    int    `json:"position_id"`
	Superior      int    `json:"superior"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
}
