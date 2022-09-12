package xini

// @Date：   2018/9/30 0030 11:04
// @Author:  Joshua Conero
// @Name:    抽象容器

// Container  abstract data container
type Container struct {
	Data        map[any]any
	_eventGetFn map[any]func() any
}

// GetData the data of `Container`
func (c *Container) GetData() map[any]any {
	if c.Data == nil {
		c.Data = map[any]any{}
	}
	return c.Data
}

func (c *Container) Get(key any) (bool, any) {
	data := c.GetData()
	value, has := data[key]
	if !has {
		if c._eventGetFn == nil {
			c._eventGetFn = map[any]func() any{}
		}
		if fn, evtHas := c._eventGetFn[key]; evtHas {
			value = fn()
			has = true
		}
	}
	return has, value
}

// GetDef get value by key with default.
func (c *Container) GetDef(key any, def any) any {
	return c.Value(key, nil, def)
}

// SetFunc set func to container
func (c *Container) SetFunc(key string, fn func() any) *Container {
	if c._eventGetFn == nil {
		c._eventGetFn = map[any]func() any{}
	}
	c._eventGetFn[key] = fn
	return c
}

// HasKey checkout if keys exist
func (c *Container) HasKey(keys ...any) bool {
	data := c.GetData()
	for _, key := range keys {
		if _, has := data[key]; has {
			return true
		}
	}
	return false
}

// Value the container get or set value
// Container.Value(key interface{})  			get data value
// Container.Value(key, value interface{})  	set data value
// Container.Value(key, nil, def interface{})  	get data value with default c.GetDef(key, def interface{})
func (c *Container) Value(params ...any) any {
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

// Set set container value
func (c *Container) Set(key, value any) *Container {
	c.GetData()
	c.Data[key] = value
	return c
}

// Del del key from container
func (c *Container) Del(key any) bool {
	if c.HasKey(key) {
		delete(c.Data, key)
		return true
	}
	return false
}

// Merge merge data from map set
func (c *Container) Merge(data map[any]any) *Container {
	for k, v := range data {
		c.Set(k, v)
	}
	return c
}

// Reset reset container will del all data
func (c *Container) Reset() *Container {
	c.Data = map[any]any{}
	return c
}
