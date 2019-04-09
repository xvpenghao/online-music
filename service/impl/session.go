package impl

type SessionService struct {
	BaseService
}

/*
*@Title: 设置session
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/9
 */
func (receiver *SessionService) SetSession(key, value, expire string) error {
	receiver.BeforeLog("SetSession")

	return nil
}

/*
*@Title: 得到session
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/9
 */
func (receiver *SessionService) GetSession(key string) (string, error) {
	receiver.BeforeLog("GetSession")

	return "", nil
}

/*
*@Title: 删除session
*@Description:
*@User: 徐鹏豪
*@Date 2019/4/9
 */
func (receiver *SessionService) DelSession(key string) error {
	receiver.BeforeLog("DelSession")

	return nil
}
