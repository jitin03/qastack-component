package service

import (
	"qastack-components/domain"
	"qastack-components/dto"
	"qastack-components/errs"
)

type ComponentService interface {
	AddComponent(request dto.AddComponentRequest) (*dto.AddComponentResponse, *errs.AppError)
	AllComponent() ([]dto.AddComponentResponse, *errs.AppError)
	DeleteComponent(id int) (dto.DeleteComponentResponse, *errs.AppError)
	UpdateComponent(id int, request dto.UpdateComponentRequest) (dto.UpdateComponentResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.ComponentRepository
}

func (s DefaultUserService) DeleteComponent(id int) (dto.DeleteComponentResponse, *errs.AppError) {

	err := s.repo.DeleteComponent(id)
	if err != nil {
		return dto.DeleteComponentResponse{
			Message: "Not Delete",
		}, err
	}
	message := dto.DeleteComponentResponse{
		Message: "Deleted",
	}
	return message, nil
}

func (s DefaultUserService) UpdateComponent(id int, req dto.UpdateComponentRequest) (dto.UpdateComponentResponse, *errs.AppError) {
	c := domain.Component{
		Name:       req.Name,
		Project_Id: req.Project_id,

	}
	err := s.repo.UpdateComponent(id, c)
	if err != nil {
		return dto.UpdateComponentResponse{
			Message: "Not Updated",
		}, err
	}

	message := dto.UpdateComponentResponse{
		Message: "Component is updated ",
	}
	return message, nil

}
func (s DefaultUserService) AllComponent() ([]dto.AddComponentResponse, *errs.AppError) {

	components, err := s.repo.AllComponent()
	if err != nil {
		return nil, err
	}
	response := make([]dto.AddComponentResponse, 0)
	for _, component := range components {
		response = append(response, component.ToDto())
	}
	return response, err
}

func (s DefaultUserService) AddComponent(req dto.AddComponentRequest) (*dto.AddComponentResponse, *errs.AppError) {

	c := domain.Component{

		Name:       req.Name,
		Project_Id: req.Project_id,
	}

	if newComponent, err := s.repo.AddComponent(c); err != nil {
		return nil, err
	} else {
		return newComponent.ToAddComponentResponseDto(), nil
	}
}

func NewComponentService(repository domain.ComponentRepository) DefaultUserService {
	return DefaultUserService{repository}
}
