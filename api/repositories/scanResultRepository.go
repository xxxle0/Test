package repositories

type ScanResultRepositoryI interface {
	Create()
	FindOne()
	FindMany()
	Delete()
	Update()
}

type ScanResultRepository struct {
}

func InitScanResultRepository() ScanResultRepositoryI {
	return &ScanResultRepository{}
}

func (r *ScanResultRepository) Create() {

}

func (r *ScanResultRepository) FindOne() {

}

func (r *ScanResultRepository) FindMany() {

}

func (r *ScanResultRepository) Delete() {

}

func (r *ScanResultRepository) Update() {

}
