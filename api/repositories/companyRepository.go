package repositories

type CompanyRepositoryI interface {
	Create()
	Update()
	Delete()
	FindOne()
}

type CompanyRepository struct {
}

func InitCompanyRepository() CompanyRepositoryI {
	return &CompanyRepository{}
}

func (r *CompanyRepository) Create() {

}

func (r *CompanyRepository) Update() {

}

func (r *CompanyRepository) Delete() {

}

func (r *CompanyRepository) FindOne() {

}
