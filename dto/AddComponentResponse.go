package dto

type AddComponentResponse struct {
	Component_Id int		`json:"component_id"`
	Name 		 string 	`json:"component_name"`
	Project_id 	int 		`json:"project_id"`

}