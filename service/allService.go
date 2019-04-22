package service

const (
	ServiceIBase      = "IBase"
	ServiceIBaseInit  = "IBaseInit"
	ServiceIUser      = "IUser"
	ServiceILogin     = "ILogin"
	ServiceISession   = "ISession"
	ServiceISongCover = "ISongCover"
	ServiceIChannel   = "IChannel"
	ServiceISong      = "ISong"
	ServiceIData      = "IData"
)

var (
	allService = map[string]interface{}{}
)

func SetServiceMap(serviceName string, service interface{}) {
	allService[serviceName] = service
}
