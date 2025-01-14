package main

import "fmt"

func CreatMap() {
	//map键值对没有顺序
	var step = make(map[string]string)
	step = map[string]string{
		"第一名": "深圳",
		"第二名": "上海",
		"第三名": "北京",
	}
	fmt.Println(step)
	//添加键值对
	step["第四名"] = "广州"
	fmt.Println(step)
	//第二种创建键值对的方法
	var step1 = map[string]string{
		"第一名": "深圳",
		"第二名": "上海",
		"第三名": "北京",
	}
	fmt.Println(step1)
	//键值对引用类型，可以对所指向的键值对进行修改
	var step2 = step1
	step2["第五名"] = "天津"
	fmt.Println(step2)
	fmt.Println(step1)
	//删除map中的键值对
	delete(step2, "第五名")
	fmt.Println(step1)
	fmt.Println(step2)
	//判断是否存在这个键值对，不存在的话value就是空的
	//step1["第五名"]这个函数返回值有两个，一个是数值，一个是是否存在
	var value, ok = step1["第五名"]
	fmt.Println(ok, value)
	var value1, ok1 = step1["第三名"]
	fmt.Println(ok1, value1)
}
func main() {
	CreatMap()
}
