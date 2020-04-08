package main

import (
	"encoding/json"
	"fmt"
)


// TAG：标签，给结构体的每一个字段，打上标签，冒号前面是类型，后面是标签名
// json:key,key为 '-'表示不序列化
// omitempty：忽略空值
// type 结构体与类型不一致时，需要指定 结构提对应的类型
// 一个 标准的 msg
type Msg struct {
	Action string
	Value  int32
}

func main() {
	s := &Msg{}
	s.Action = "leaf"


	// json 序列化的时候 属性为小些 会被忽略
	da, _ := json.Marshal(s)

	fmt.Println(da)

	fmt.Println(string(da))

}
