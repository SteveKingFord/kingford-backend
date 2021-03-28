package collection

import (
	"kingford-backend/dto"
	"kingford-backend/global"
	"kingford-backend/repository"
)

type DeleteService struct {
}

func (s *DeleteService) Delete(id string) *dto.Response {
	repo := repository.CollectionRepository{DB: global.DB}
	err := repo.Delete(id)

	if err != nil {
		return &dto.Response{
			Status: 400,

			Msg: err.Error(),
		}
	}
	return &dto.Response{
		Status: 200,
		Data:   true,
		Msg:    "success",
	}
}
