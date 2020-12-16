package repository

type InstaPickerRepository struct {
}

func ProvideInstaPickerRepository() *InstaPickerRepository {
	return &InstaPickerRepository{}
}

func (repository InstaPickerRepository) DoSomething() {

}
