package about

import "github.com/hardzal/portfolio-api-go/services"

type AboutHandler struct {
	AboutService services.AboutService
}

func NewAboutHandler(aboutService services.AboutService) *AboutHandler {
	return &AboutHandler{AboutService: aboutService}
}

func (a *AboutHandler) GetAboutHandler() error {
	panic("belum")
}
