package providerSet

type FooBarService struct {
	*BarRepository
	*FooRepository
}

func NewFooBarService(repoBar *BarRepository, repoFoo *FooRepository) *FooBarService {
	return &FooBarService{
		BarRepository: repoBar,
		FooRepository: repoFoo,
	}
}