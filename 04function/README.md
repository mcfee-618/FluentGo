## 函数

go 语言函数定义格式如下：
```
    func function_name( [parameter list] ) [return_types] {
       函数体
    }

    func max(a int32, b int32) int32 {
        if a > b {
            return a
        } else {
            return b
        }
    }
```
与许多后端语言不同的是， go 语言中的函数是支持多返回值的。


## 方法

go 语言还有一种特殊的函数，叫做方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：

```
    func (variable_name variable_data_type) function_name() [return_type]{
    /* 函数体*/
    }
```


