package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type tFirst struct {
	id interface{int}
	name interface{string}
	isHere interface{bool}
	t interdace{null}
	arr inerface{[]}
	str interface{string, string}
}

type tSecond struct {
	calendar, autoContinue, buttons, status interface{bool}
	delay interface{int}
	list interface{[]}
}

func main(){
	file, _ := ioutil.ReadFile("test.json")
	data := tFirst{}
	_ = json.Unmarshal([]byte(file), &data)
	
	var first tFirst
	first.id, first.name, first.isHere, first.t, first.arr,  first.str = data.id, data.name, data.isHere, data.t, data.arr, data.str
	
	var second tSecond
	second.calendar, second.autoContinue, second.buttons, second.status, second.delay, second.list = data.calendar, data.autoContinue, data.buttons, data.status, data.delay, data.list
}

