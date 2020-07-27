package xini

// @Date：   2018/9/30 0030 11:04
// @Author:  Joshua Conero
// @Name:    抽象容器

// 抽象参数容器
type Container struct {
	Data        map[interface{}]interface{}
	_eventGetFn map[interface{}]func() interface{}
}

// data 数据检测
// 获取容器的全部数据，并且可实例化数据
func (c *Container) GetData() map[interface{}]interface{} {
	if c.Data == nil {
		c.Data = map[interface{}]interface{}{}
	}
	return c.Data
}

// 获取值
func (c *Container) Get(key interface{}) (bool, interface{}) {
	data := c.GetData()
	value, has := data[key]
	if !has {
		if c._eventGetFn == nil {
			c._eventGetFn = map[interface{}]func() interface{}{}
		}
		if fn, evtHas := c._eventGetFn[key]; evtHas {
			value = fn()
			has = true
		}
	}
	return has, value
}

// 获取值，且含默认值
func (c *Container) GetDef(key interface{}, def interface{}) interface{} {
	return c.Value(key, nil, def)
}

// get 参数注册
func (c *Container) GetFunc(key string, fn func() interface{}) *Container {
	if c._eventGetFn == nil {
		c._eventGetFn = map[interface{}]func() interface{}{}
	}
	c._eventGetFn[key] = fn
	return c
}

// 是否存在键值
func (c *Container) HasKey(key interface{}) bool {
	data := c.GetData()
	_, has := data[key]
	return has
}

// 容器值得获取/设置
// Container.Value(key interface{})  			获取数据
// Container.Value(key, value interface{})  	设置数据
// Container.Value(key, nil, def interface{})  	获取值并带参数，可使用新的方法 c.GetDef(key, def interface{})
func (c *Container) Value(params ...interface{}) interface{} {
	if len(params) > 2 { // key, nil, def
		if has, value := c.Get(params[0].(string)); has {
			return value
		}
		return params[2]
	} else if len(params) > 1 { // key, value
		c.Set(params[0].(string), params[1])
	} else if len(params) == 1 { // key
		if has, value := c.Get(params[0].(string)); has {
			return value
		}
	}
	return nil
}

// 设置容器参数
func (c *Container) Set(key, value interface{}) *Container {
	c.GetData()
	c.Data[key] = value
	return c
}

// 删除容器的值
func (c *Container) Del(key interface{}) bool {
	if c.HasKey(key) {
		delete(c.Data, key)
		return true
	}
	return false
}

// 数据合并
func (c *Container) Merge(data map[interface{}]interface{}) *Container {
	for k, v := range data {
		c.Set(k, v)
	}
	return c
}

// 重置容器的值
func (c *Container) Reset() *Container {
	c.Data = map[interface{}]interface{}{}
	return c
}
