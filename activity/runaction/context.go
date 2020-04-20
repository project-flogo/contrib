package runaction

import (
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/app/resource"
	"github.com/project-flogo/core/support/service"
)

type actionInitContext struct {
	resManager  *resource.Manager
	servManager *service.Manager
	settings    map[string]interface{}
}

func (a actionInitContext) ResourceManager() *resource.Manager {
	return a.resManager
}

func (a actionInitContext) ServiceManager() *service.Manager {
	return a.servManager
}

func (a actionInitContext) RuntimeSettings() map[string]interface{} {
	return a.settings
}

func getInitContext() action.InitContext {

	//Deafult size.
	resources := make(map[string]*resource.Resource, 10)

	resManager := resource.NewManager(resources)

	servManager := service.NewServiceManager()

	return actionInitContext{resManager: resManager, servManager: servManager, settings: make(map[string]interface{})}
}
