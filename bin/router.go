package bin

// @Date：   2018/10/30 0030 12:41
// @Author:  Joshua Conero
// @Name:    bin 路由器

type Router struct {
	// 别名映射组
	Alias        map[string]string
	UnfindAction func(action string) // 路由失败
	EmptyAction  func()              // 路由失败
}

// 获取 action参数
func (router *Router) GetAction(action string) string {
	if router.Alias == nil {
		router.Alias = map[string]string{}
	}
	if alias, has := router.Alias[action]; has {
		action = alias
	}
	return action
}
