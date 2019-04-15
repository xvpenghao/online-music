package dbModel

import "time"

const (
	//查询用户信息根据用户id
	QUEYR_USER_INFO_BY_Id = ``
	//登录
	QUERY_LOGIN_USER_INFO = `select  tbu.* 
                             from tb_login tbl 
                             inner join tb_user tbu on tbl.user_id = tbu.user_id
                             where (tbl.login_name = ? or tbl.login_email = ?) 
                                   and tbl.password = ?`
	QUERY_USER_COUNTS_BY_UID = `select count(*)
                                from tb_user tbu
                                where tbu.user_id <> ? and (tbu.user_name = ? or tbu.email = ?)`
)

type User struct {
	//用户id
	Id string `gorm:"column:user_id"`
	//用户姓名
	Name string `gorm:"column:user_name"`
	//邮箱
	Email string `gorm:"column:email"`
	//性别
	Gender string `gorm:"column:gender"`
	//年龄
	Age int `gorm:"column:age"`
	//生日
	Birthday time.Time `gorm:"column:birthday"`
	//删除的状态 1删除 2不删除
	DelState int `gorm:"column:del_status"`
	//创建时间
	CreatTime time.Time `gorm:"column:create_time"`
	//创建人
	CreateUser string `gorm:"column:create_user"`
	//创建人ID
	CreateUserId string `gorm:"column:create_user_id"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time"`
	//更新人
	UpdateUser string `gorm:"column:update_user"`
	//更新人ID
	UpdateUserId string `gorm:"column:update_user_id"`
}

type UserLogin struct {
	//登录id
	LoginId string `gorm:"column:login_id"`
	//用户id
	UserId string `gorm:"column:user_id"`
	//登录名
	LoginName string `gorm:"column:login_name"`
	//登陆邮箱
	LoginEmail string `gorm:"column:login_email"`
	//登陆密码
	Password string `gorm:"column:password"`
	//创建时间
	CreatTime time.Time `gorm:"column:create_time"`
	//创建人
	CreateUser string `gorm:"column:create_user"`
	//创建人ID
	CreateUserId string `gorm:"column:create_user_id"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time"`
	//更新人
	UpdateUser string `gorm:"column:update_user"`
	//更新人ID
	UpdateUserId string `gorm:"column:update_user_id"`
}

//返回登录信息
type LoginUserInfo struct {
	LoginId string `gorm:"column:login_id"`
	//用户id
	UserId string `gorm:"column:user_id"`
	//登录名
	LoginName string `gorm:"column:login_name"`
	//登陆邮箱
	LoginEmail string `gorm:"column:login_email"`
}
