## 结构体

* go 语言中 结构体定义如下：

```
    type identifier struct {
        field1 type1
        field2 type2
        ...
    }
```

* 对结构体的赋值有两种方式 
    * t := new(T)：变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值
    * t := &T{}：&T{}这种语法叫做 混合字面量语法（composite literal syntax，底层仍然会调用 new ()，这里值的顺序必须按照字段顺序来写。
```
    person := &Person {name : "xiaozhuanlan", age : 12}
```

## 数组

* go中创建数组语法如下：
```
var variable_name [SIZE] variable_type
```

* go中创建数组并初始化如下：
```
var variable_name = [SIZE]variable_type {value1, value2 ...}
```

* for range遍历如下：
```
    for index, value := range array {
        fmt.Printf("array[%d] = %s \n", index, value)
    }
```
假如不需要使用到 index 或者 value ，可以用 _ 进行代替。

## 切片

go 提供了一种类似动态数组结构的数据类型，这种类型就是切片 slice。slice 的本质是一个数据结构，实现了对数组操作的封装。切片 slice 的申明语法如下 ：
```
     var identifier []type
     array = make([]int64,10)  //申明一个为 int64 类型、长度为10的数组
```

* 切片与数组：切片的声明不需要指定长度，数组需要，切片使用make初始化

```
   array = append ( array, 1 ) // 动态追加元素
   array[1:3] // 截取切片 [lower-bound:upper-bound]
   len(slice) // 获取切片长度
   使用range 关键字进行遍历

```

## map

* map 是一种无序的键值对的集合，go语言中语法如下：

```
var map_variable map[key_data_type]value_data_type   
/* 声明变量，默认 map 是 nil */
map_variable := make(map[key_data_type]value_data_type)
// 初始化map，如果不初始化 map，那么就会创建nil map。nil map不能用来存放键值对
```   
对于 map 类型，一定要进行初始化再赋值，字典也可以指定map的初始容量，遍历还是使用range

* 基本操作

```
m[key] = elem
delete(m, key)    //删除某个键
elem, ok = m[key] //检测某个键是否存在
```

