package dto



type UpdateComponentRequest struct {

	Name string		    `json:"component_name"`
	Project_id int	`json:"project_Id"`
}