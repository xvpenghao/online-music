package service

const (
	ServiceIBase     = "IBase"
	ServiceIBaseInit = "IBaseInit"
	ServiceIUser     = "IUser"
	ServiceILogin    = "ILogin"
	ServiceISession  = "ISession"
)

var (
	allService = map[string]interface{}{}
)

func SetServiceMap(serviceName string, service interface{}) {
	allService[serviceName] = service
}
