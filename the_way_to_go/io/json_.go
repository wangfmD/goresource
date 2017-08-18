package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Js struct {
	data interface{}
}

// NewJson ...
func NewJson(data string) *Js {
	j := new(Js)
	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		return j
	}
	j.data = f
	return j
}

func (j *Js) GetMapData() map[string]interface{} {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

// Get ...
func (j *Js) Get(key string) *Js {
	m := j.GetMapData()
	if v, ok := m[key]; ok {
		j.data = v
		return j
	}
	j.data = nil
	return j
}

func (j *Js) Type() {
	fmt.Println(reflect.TypeOf(j.data))
}

func main() {
	data := "{\"a\":1,\"b\":2}"
	var f map[string]int
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	// fmt.Println(f["a"])
	// for key, value := range f {
	// 	fmt.Println(key, ":", value)
	// }
	m1 := make(map[string]string)
	m1 = map[string]string{"a": "1", "b": "2"}
	m1["key1"] = "11"
	fmt.Println(m1)
	fmt.Println(m1["a"])
	if _, ok := m1["key1"]; ok {
		fmt.Println("OK")
	} else {
		fmt.Println("False")
	}

	fmt.Println("---")
	json := `{"name":"light","weigth":"maybe55kg","result":["light","fish","dylan"]}`
	j1 := NewJson(json)
	fmt.Println(*j1)
	// j2 := j1.Get("n1ame")
	// fmt.Println(*j2)
	j1.Type()
}
