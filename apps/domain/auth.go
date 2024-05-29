package domain

type AuthRequest struct {
	EmployeeCode string `json:"employee_code" validate:"required"`
	Password     string `json:"password" validate:"required"`
}
