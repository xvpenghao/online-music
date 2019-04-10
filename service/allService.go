package service

const (
	ServiceIBase      = "IBase"
	ServiceIBaseInit  = "IBaseInit"
	ServiceIUser      = "IUser"
	ServiceILogin     = "ILogin"
	ServiceISession   = "ISession"
	ServiceISongCover = "ISongCover"
	ServiceIChannel   = "IChannel"
)

var (
	allService = map[string]interface{}{}
)

func SetServiceMap(serviceName string, service interface{}) {
	allService[serviceName] = service
}
