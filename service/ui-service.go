package service

import (
	repository "github.com/Angieuski/Desafio-Edu/Repository"
	"github.com/Angieuski/Desafio-Edu/entity"
)

type UiService interface {
	Save(entity.Ui) entity.Ui
	Update(ui entity.Ui)
	Delete(ui entity.Ui)
	FindAll() []entity.Ui
	FindByEmail(email string) entity.Ui
}
type UiServiceStruct struct {
	UiRepository repository.UiRepository
}

func New(repo repository.UiRepository) UiService {
	return &UiServiceStruct{
		UiRepository: repo,
	}
}

func (service *UiServiceStruct) FindByEmail(email string) entity.Ui {
	var user entity.Ui
	service.UiRepository.FindByEmail(email, &user)
	return user
}

func (service *UiServiceStruct) Save(Ui entity.Ui) entity.Ui {
	service.UiRepository.Save(Ui)
	return Ui
}

func (service *UiServiceStruct) Update(Ui entity.Ui) {
	service.UiRepository.Update(Ui)
}

func (service *UiServiceStruct) Delete(Ui entity.Ui) {
	service.UiRepository.Delete(Ui)
}

func (service *UiServiceStruct) FindAll() []entity.Ui {
	return service.UiRepository.FindAll()
}
