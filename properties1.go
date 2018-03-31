package cstruct

import (
	"fmt"
	"os"
	"reflect"
	"sync"
)

type StructProperties struct {
	Prop  []*Properties
	stype reflect.Type
}

var (
	propertiesMu  sync.RWMutex
	propertiesMap = make(map[reflect.Type]*StructProperties)
)

func GetProperties(t reflect.Type) *StructProperties {
	if t.Kind() != reflect.Struct {
		panic("cstruct: type must have kind struct")
	}

	propertiesMu.RLock()
	sprop, ok := propertiesMap[t]
	propertiesMu.RUnlock()
	if ok {
		return sprop
	}

	propertiesMu.Lock()
	sprop = getPropertiesLocked(t)
	propertiesMu.Unlock()
	return sprop
}

func getPropertiesLocked(t reflect.Type) *StructProperties {
	if prop, ok := propertiesMap[t]; ok {
		return prop
	}

	prop := new(StructProperties)
	propertiesMap[t] = prop
	prop.Prop = make([]*Properties, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		p := new(Properties)
		name := f.Name
		p.init(f.Type, name, f.Tag.Get("c"), &f, false)
		prop.Prop[i] = p
		if p.enc == nil {
			fmt.Fprintln(os.Stderr, "cstruct: no encoder for", f.Name, f.Type.String(), "[GetProperties]")
		}
	}

	return prop
}
