package provider

type simpleRepository struct {
}

func NewSimpleRepository() *simpleRepository {
	return &simpleRepository{}
}

type simpleService struct {
	*simpleRepository
}

func NewSimpleService(repo *simpleRepository) *simpleService {
	return &simpleService{
		simpleRepository: repo,
	}
}