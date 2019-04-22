package dbModel

const (
	//统计网站使用群体
	QUERY_WEBSITE_USE_GROUP = `select '儿童' as wsug,count(*)as cnt
                               from tb_user 
                               where age BETWEEN 5 and 11
                               UNION all
                               select '少年' as wsug,count(*)as cnt
                               from tb_user 
                               where age BETWEEN 12 and 18
                               UNION all
                               select '青年' as wsug,count(*)as cnt
                               from tb_user 
                               where age BETWEEN 19 and 35
                               UNION all
                               select '中年' as wsug,count(*)as cnt
                               from tb_user 
                               where age BETWEEN 36 and 59
                               UNION all
                               select '老年' as wsug,count(*)as cnt
                               from tb_user 
                               where age >=60`
	//查询性别比例
	QUERY_GENDER_PROPORTION = `select gender,count(*)/(
                                                  select count(*)
                                                  from tb_user)* 100 as cnt
                               from tb_user 
                               group by gender`
)

//网站使用群体
type WebSiteUseGroup struct {
	//群体名称
	GroupName string `grom:"column:wsug"`
	//群体数量
	GroupCounts int `gorm:"column:cnt"`
}

//查询性别比例
type GenderProportion struct {
	//性别
	Gender string `gorm:"column:gender"`
	//性别分类数量
	GenderCounts float64 `gorm:"column:cnt"`
}
