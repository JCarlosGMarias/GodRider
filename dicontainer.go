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
	configSvcr       services.ConfigurationServicer
	userSvcr         services.UserServicer
	providerSvcr     services.ProviderServicer
	userProviderSvcr services.UserProviderServicer
	orderSvcr        services.OrderServicer
	validationSvcr   services.ValidationServicer
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
	apiUrlsIstruct := &infrastructures.APIUrlsInfrastructure{}
	apiUrlsIstruct.TableName("apiurl")
	di.apiUrlsIstructr = apiUrlsIstruct

	userIstruct := &infrastructures.UserInfrastructure{}
	userIstruct.TableName("user")
	di.userIstructr = userIstruct

	providerIstruct := &infrastructures.ProviderInfrastructure{}
	providerIstruct.TableName("provider")
	di.providerIstructr = providerIstruct

	userProviderIstruct := &infrastructures.UserProviderInfrastructure{}
	userProviderIstruct.TableName("userprovider")
	di.userProviderIstructr = userProviderIstruct

	di.webClientFactorier = &webclients.ClientFactory
}

func (di *DIContainer) allInfrastructuresCreated() bool {
	return di.apiUrlsIstructr != nil &&
		di.userIstructr != nil &&
		di.providerIstructr != nil &&
		di.userProviderIstructr != nil &&
		di.webClientFactorier != nil
}

func (di *DIContainer) registerServices() error {
	if !di.allInfrastructuresCreated() {
		return fmt.Errorf("Unable to register services: Check all infrastructure register")
	}

	configSvc := &services.ConfigurationService{}
	configSvc.APIUrlsInfrastructure(&di.apiUrlsIstructr)
	di.configSvcr = configSvc

	userSvc := &services.UserService{}
	userSvc.UserInfrastructure(&di.userIstructr)
	di.userSvcr = userSvc

	providerSvc := &services.ProviderService{}
	providerSvc.ProviderInfrastructure(&di.providerIstructr)
	di.providerSvcr = providerSvc

	userProviderSvc := &services.UserProviderService{}
	userProviderSvc.UserProviderInfrastructure(&di.userProviderIstructr)
	di.userProviderSvcr = userProviderSvc

	orderSvc := &services.OrderService{}
	orderSvc.ProviderInfrastructure(&di.providerIstructr)
	orderSvc.Factory(&di.webClientFactorier)
	di.orderSvcr = orderSvc

	validationSvc := &services.ValidationService{}
	validationSvc.UserSrv(&di.userSvcr)
	di.validationSvcr = validationSvc

	return nil
}

func (di *DIContainer) allServicesCreated() bool {
	return di.configSvcr != nil &&
		di.userSvcr != nil &&
		di.providerSvcr != nil &&
		di.userProviderSvcr != nil &&
		di.orderSvcr != nil &&
		di.validationSvcr != nil
}

func (di *DIContainer) registerControllers() error {
	if !di.allInfrastructuresCreated() {
		return fmt.Errorf("Unable to register controllers: Check all service register")
	}

	configCtrl := &controllers.ConfigurationController{}
	configCtrl.ConfigSrv(&di.configSvcr)
	configCtrl.ValidationSrv(&di.validationSvcr)
	di.ConfigCtrlr = configCtrl

	userCtrl := &controllers.UserController{}
	userCtrl.UserSrv(&di.userSvcr)
	userCtrl.ValidationSrv(&di.validationSvcr)
	di.UserCtrlr = userCtrl

	providerCtrl := &controllers.ProviderController{}
	providerCtrl.ProviderSrv(&di.providerSvcr)
	providerCtrl.UserProviderSrv(&di.userProviderSvcr)
	providerCtrl.ValidationSrv(&di.validationSvcr)
	di.ProviderCtrlr = providerCtrl

	orderCtrl := &controllers.OrderController{}
	orderCtrl.OrderSrv(&di.orderSvcr)
	orderCtrl.ValidationSrv(&di.validationSvcr)
	di.OrderCtrlr = orderCtrl

	return nil
}
