package service

import (
	"sync"
)

type Attr struct {
	lock *sync.Mutex // 扯犊子玩意，同协程都不可重入
	attr map[string]interface{}
}

func NewAttr(params map[string]interface{}) *Attr {
	attr := new(Attr)
	attr.lock = new(sync.Mutex)
	attr.attr = make(map[string]interface{})
	for k, v := range params {
		attr.Set(k, v)
	}
	return attr
}
func (attr *Attr) Copy(srcAttr IAttr) {
	attrs := srcAttr.CopyToMap()
	for k, v := range attrs {
		attr.Set(k, v)
	}
}

func (attr *Attr) CopyFromMap(newAttr map[string]interface{}) {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	attr.attr = newAttr
}
func (attr *Attr) CopyToMap() map[string]interface{} {
	result := make(map[string]interface{})
	defer attr.lock.Unlock()
	attr.lock.Lock()
	for k, v := range attr.attr {
		result[k] = v
	}
	return result
}

func (attr *Attr) Set(key string, value interface{}) {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	attr.attr[key] = value
}

func (attr *Attr) Get(key string) (value interface{}) {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	return attr.get(key)
}

func (attr *Attr) get(key string) (value interface{}) {
	if v, ok := attr.attr[key]; ok {
		return v
	}
	return nil
}

func (attr *Attr) GetBool(key string) bool {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		wantValue, convertError := ConvertInterface2Bool(v)
		if convertError == nil {
			return wantValue
		}
	}
	return false
}

func (attr *Attr) GetInt8(key string) int8 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		wantValue, convertError := ConvertInterface2Int8(v)
		if convertError == nil {
			return wantValue
		}
	}
	return -1
}

func (attr *Attr) GetInt16(key string) int16 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		wantValue, convertError := ConvertInterface2Int16(v)
		if convertError == nil {
			return wantValue
		}
	}
	return -1
}

func (attr *Attr) GetInt32(key string) int32 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		wantValue, convertError := ConvertInterface2Int32(v)
		if convertError == nil {
			return wantValue
		}
	}
	return -1
}

func (attr *Attr) GetInt64(key string) int64 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		wantValue, convertError := ConvertInterface2Int64(v)
		if convertError == nil {
			return wantValue
		}
	}
	return -1
}

func (attr *Attr) GetInt(key string) int {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		wantValue, convertError := ConvertInterface2Int(v)
		if convertError == nil {
			return wantValue
		}
	}
	return -1
}

func (attr *Attr) GetString(key string) string {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		wantValue, convertError := ConvertInterface2String(v)
		if convertError == nil {
			return wantValue
		}
	}
	return ""
}
