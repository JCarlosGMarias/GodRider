package main

import (
	"fmt"
	"godrider/controllers"
	"godrider/infrastructures"
	"godrider/services"
	"godrider/webclients"
)

// DIContainer is the responsible for register all structs and connect them as their relationships expect
type DIContainer struct {
	// Infrastructures
	apiUrlsIstructr      infrastructures.APIUrlsInfrastructurer
	userIstructr         infrastructures.UserInfrastructurer
	userProviderIstructr infrastructures.UserProviderInfrastructurer
	providerIstructr     infrastructures.ProviderInfrastructurer
	webClientFactorier   webclients.WebClientFactorier
	// Services
	configSrvr       services.ConfigurationServicer
	userSrvr         services.UserServicer
	providerSrvr     services.ProviderServicer
	userProviderSrvr services.UserProviderServicer
	orderSrvr        services.OrderServicer
	validationSrvr   services.ValidationServicer
	// Controllers
	ConfigCtrlr   controllers.ConfigurationControllerer
	UserCtrlr     controllers.UserControllerer
	ProviderCtrlr controllers.ProviderControllerer
	OrderCtrlr    controllers.OrderControllerer
}

// Register instantiates all infrastructures, services and controllers in order
func (di *DIContainer) Register() error {
	di.registerInfrastructures()
	if err := di.registerServices(); err != nil {
		return err
	}
	return di.registerControllers()
}

func (di *DIContainer) registerInfrastructures() {
	di.apiUrlsIstructr = &infrastructures.APIUrlsInfrastructure{}
	di.userIstructr = &infrastructures.UserInfrastructure{}
	userProviderIstruct := &infrastructures.UserProviderInfrastructure{}
	di.userProviderIstructr = userProviderIstruct
	di.providerIstructr = &infrastructures.ProviderInfrastructure{}
	di.webClientFactorier = &webclients.ClientFactory

	userProviderIstruct.TableName("userprovider")
}

func (di *DIContainer) allInfrastructuresCreated() bool {
	return di.apiUrlsIstructr != nil &&
		di.userIstructr != nil &&
		di.userProviderIstructr != nil &&
		di.providerIstructr != nil &&
		di.webClientFactorier != nil
}

func (di *DIContainer) registerServices() error {
	if !di.allInfrastructuresCreated() {
		return fmt.Errorf("Unable to register services: Check all infrastructure register")
	}

	configSrv := &services.ConfigurationService{}
	configSrv.APIUrlsInfrastructure(&di.apiUrlsIstructr)
	di.configSrvr = configSrv

	userSrv := &services.UserService{}
	userSrv.UserInfrastructure(&di.userIstructr)
	di.userSrvr = userSrv

	providerSrv := &services.ProviderService{}
	providerSrv.ProviderInfrastructure(&di.providerIstructr)
	di.providerSrvr = providerSrv

	userProviderSrv := &services.UserProviderService{}
	userProviderSrv.UserProviderInfrastructure(&di.userProviderIstructr)
	di.userProviderSrvr = userProviderSrv

	orderSrv := &services.OrderService{}
	orderSrv.ProviderInfrastructure(&di.providerIstructr)
	orderSrv.Factory(&di.webClientFactorier)
	di.orderSrvr = orderSrv

	validationSrv := &services.ValidationService{}
	validationSrv.UserSrv(&di.userSrvr)
	di.validationSrvr = validationSrv

	return nil
}

func (di *DIContainer) allServicesCreated() bool {
	return di.configSrvr != nil &&
		di.userSrvr != nil &&
		di.providerSrvr != nil &&
		di.userProviderSrvr != nil &&
		di.orderSrvr != nil &&
		di.validationSrvr != nil
}

func (di *DIContainer) registerControllers() error {
	if !di.allInfrastructuresCreated() {
		return fmt.Errorf("Unable to register controllers: Check all service register")
	}

	configCtrl := &controllers.ConfigurationController{}
	configCtrl.ConfigSrv(&di.configSrvr)
	configCtrl.ValidationSrv(&di.validationSrvr)
	di.ConfigCtrlr = configCtrl

	userCtrl := &controllers.UserController{}
	userCtrl.UserSrv(&di.userSrvr)
	userCtrl.ValidationSrv(&di.validationSrvr)
	di.UserCtrlr = userCtrl

	providerCtrl := &controllers.ProviderController{}
	providerCtrl.ProviderSrv(&di.providerSrvr)
	providerCtrl.UserProviderSrv(&di.userProviderSrvr)
	providerCtrl.ValidationSrv(&di.validationSrvr)
	di.ProviderCtrlr = providerCtrl

	orderCtrl := &controllers.OrderController{}
	orderCtrl.OrderSrv(&di.orderSrvr)
	orderCtrl.ValidationSrv(&di.validationSrvr)
	di.OrderCtrlr = orderCtrl

	return nil
}
