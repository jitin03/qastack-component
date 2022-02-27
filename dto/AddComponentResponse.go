package dto

type AddComponentResponse struct {
	Component_Id string `json:"component_id"`
	Name         string `json:"component_name"`
	Project_id   string `json:"project_id"`
	CreateDate   string `json:"create_date"`
	UpdateDate   string `json:"update_date"`
}
