package service

type TableService struct{}

func InitTableService() *TableService {
	service := &TableService{}
	return service
}

func (s TableService) FindTable() {
}
