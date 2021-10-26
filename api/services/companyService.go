package services

type CompanyServiceI interface {
	CreateCompany()
	GetCompanyById()
}

type CompanyService struct {
}

func InitCompanyService() CompanyServiceI {
	return &CompanyService{}
}

func (s *CompanyService) CreateCompany() {

}

func (s *CompanyService) GetCompanyById() {

}
