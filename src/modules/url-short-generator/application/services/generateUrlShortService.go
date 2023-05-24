package application_services

type GenerateUrlShortService struct {
}

func NewGenerateUrlShortService() GenerateUrlShortService {
	return GenerateUrlShortService{}
}
func (g GenerateUrlShortService) Execute(url string) string {
	return ""
}
