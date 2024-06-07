package providerSet

type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

func NewFooService(repo FooRepository) *FooService {
	return &FooService{
		FooRepository: &repo,
	}
}