# go学习笔记

##集合

1.定义集合,是key=>value的形式

`map_variable := make(map[key_data_type]value_data_type)`
`map_variable := make(map[键的类型]键值的类型)`

interface{} 可以适配任何类型

2.例子
```go
//接口成功返回
func Ok(data map[string]interface{}, msg string) map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = 200
	result["msg"] = msg
	result["data"] = data
	return result
}
```

## 切片

1.定义切片，切片是不固定长度的数组的抽象

`slice1 := make([]type, len)`

2.切片初始化

`s := [] int {1,2,3 } `

直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

3.例子
```go
//可以批量插入数据整理
data := []models.Category{
			{Name: "slene"},
			{Name: "slene1"},
			{Name: "slene2"},
		}
//定义的结构体模型，models.Category 可作为一种类型
```