package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type DataReceiverType uint8

const (
	ReceiverContent DataReceiverType = iota
	ReceiverFile
	ReceiverUrl
)

const (
	RJson = "json"
	RUrl  = "url"
)

// the data parse tool
type DataReceiver interface {
	Name() string
	Receiver(DataReceiverType, string) DataReceiver
	GetData() map[string]interface{}
}

// the url data receiver
type BaseReceiver struct {
	vMap map[string]interface{}
}

// the data parse grammar, support json/url
type JsonReceiver struct {
	BaseReceiver
}

// base json parse
func (c *JsonReceiver) parseJson(vByte []byte) *JsonReceiver {
	var jsonData interface{}
	er := json.Unmarshal(vByte, &jsonData)
	if er == nil {
		rv := reflect.ValueOf(jsonData)
		if rv.Kind() == reflect.Map {
			mr := rv.MapRange()
			if c.vMap == nil {
				c.vMap = map[string]interface{}{}
			}
			for mr.Next() {
				c.vMap[fmt.Sprintf("%v", mr.Key().Interface())] = mr.Value().Interface()
			}
		}
	}
	return c
}

//parse json string
func (c *JsonReceiver) JsonStr(vStr string) *JsonReceiver {
	return c.parseJson([]byte(vStr))
}

// parse json from json-file
func (c *JsonReceiver) JsonFile(filename string) *JsonReceiver {
	bys, er := ioutil.ReadFile(filename)
	if er == nil {
		return c.parseJson(bys)
	}
	return c
}

// parse json from json-url, only http.get
func (c *JsonReceiver) JsonUrl(vUrl string) *JsonReceiver {
	if bys := GetUrlContent(vUrl); bys != nil {
		return c.parseJson(bys)
	}
	return c
}

func (c *JsonReceiver) Name() string {
	return "json"
}

// receiver data
func (c *JsonReceiver) Receiver(vType DataReceiverType, content string) DataReceiver {
	switch vType {
	case ReceiverContent:
		c.JsonStr(content)
	case ReceiverFile:
		c.JsonFile(content)
	case ReceiverUrl:
		c.JsonUrl(content)
	}
	return c
}

//get finally data by parse
func (c *JsonReceiver) GetData() map[string]interface{} {
	return c.vMap
}

// get url content
func GetUrlContent(vUrl string) []byte {
	resp, er := http.Get(vUrl)
	if er == nil {
		if resp.StatusCode == http.StatusOK {
			bys, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				return bys
			}
		}
	}
	return nil
}

// the url data receiver
type UrlReceiver struct {
	BaseReceiver
}

//parse json string
func (c *UrlReceiver) UrlStr(vStr string) *UrlReceiver {
	return c.parse(string(vStr))
}

// parse json from json-file
func (c *UrlReceiver) UrlFile(filename string) *UrlReceiver {
	bys, er := ioutil.ReadFile(filename)
	if er == nil {
		return c.parse(string(bys))
	}
	return c
}

// parse json from json-url, only http.get
func (c *UrlReceiver) UrlUrl(vUrl string) *UrlReceiver {
	if bys := GetUrlContent(vUrl); bys != nil {
		return c.parse(string(bys))
	}
	return c
}

// base json parse
func (c *UrlReceiver) parse(vStr string) *UrlReceiver {
	if u, er := url.ParseQuery(vStr); er == nil {
		if vJson, err := json.Marshal(u); err == nil {
			var jsonData interface{}
			er := json.Unmarshal(vJson, &jsonData)
			if er == nil {
				rv := reflect.ValueOf(jsonData)
				if rv.Kind() == reflect.Map {
					mr := rv.MapRange()
					if c.vMap == nil {
						c.vMap = map[string]interface{}{}
					}
					for mr.Next() {
						c.vMap[fmt.Sprintf("%v", mr.Key().Interface())] = mr.Value().Interface()
					}
				}
			}
		}
	}
	return c
}

func (c *UrlReceiver) Name() string {
	return "url"
}

// receiver data
func (c *UrlReceiver) Receiver(vType DataReceiverType, content string) DataReceiver {
	switch vType {
	case ReceiverContent:
		c.UrlStr(content)
	case ReceiverFile:
		c.UrlFile(content)
	case ReceiverUrl:
		c.UrlUrl(content)
	}
	return c
}

//get finally data by parse
func (c *UrlReceiver) GetData() map[string]interface{} {
	return c.vMap
}

//get DataReceiver by different type.
func NewDataReceiver(vType string) (DataReceiver, error) {
	var dr DataReceiver = nil
	var er error = nil
	switch strings.ToLower(vType) {
	case RJson:
		dr = &JsonReceiver{}
	case RUrl:
		dr = &UrlReceiver{}
	default:
		er = errors.New("DataReceiver type only support: json, url")
	}
	return dr, er
}
