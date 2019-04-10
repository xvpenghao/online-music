package service

type (
	ISessionService interface {
		IBaseService
		//设置session
		SetSession(key, value, expire string) error
		//得到session
		GetSession(key string) (string, error)
		//删除session
		DelSession(key string) error
	}
)

func NewSessionService(init IBaseServiceInit) ISessionService {
	temp := allService[ServiceISession]
	result := temp.(ISessionService)
	result.SetInitInfo(init)
	return result
}
