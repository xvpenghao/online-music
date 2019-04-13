package impl

import (
	"online-music/service"
)

func init() {
	InitServiceMap()
}

func InitServiceMap() {
	service.SetServiceMap(service.ServiceIBase, new(BaseService))
	service.SetServiceMap(service.ServiceIBaseInit, new(BaseServiceInit))
	service.SetServiceMap(service.ServiceIUser, new(UserService))
	service.SetServiceMap(service.ServiceILogin, new(LoginService))
	service.SetServiceMap(service.ServiceISession, new(SessionService))
	service.SetServiceMap(service.ServiceISongCover, new(SongCoverService))
	service.SetServiceMap(service.ServiceISong, new(SongService))
}
