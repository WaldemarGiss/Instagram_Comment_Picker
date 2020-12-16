package service

type instaPickerRepository interface {
	DoSomething()
}

type InstaPickerService struct {
	repository instaPickerRepository
}

func ProvideInstaPickerService(repository instaPickerRepository) *InstaPickerService {
	return &InstaPickerService{repository: repository}
}

func (service *InstaPickerService) DoSomething() {

}
