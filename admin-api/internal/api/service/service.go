package service

import (
	sg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
)

type ServiceManager struct {
	serviceGateway sg.ServiceGateway
}

func NewServiceManager(serviceGateway sg.ServiceGateway) *ServiceManager {
	return &ServiceManager{
		serviceGateway: serviceGateway,
	}
}

// func (sm *ServiceManager) HandleGetAllServices(w http.ResponseWriter, req *http.Request) {
// 	services, err := sm.serviceGateway.GetAllServices()
// 	if err != nil {

// 	}
// }
