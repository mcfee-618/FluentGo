package main
import ( 
"fmt"
"reflect"
"unsafe"
)

func main()  {
	var name string
	name="feipeixuan"
	fmt.Println(name)
	typeOfA := reflect.TypeOf(name)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	var value uint32
	value =444
    typeOfA = reflect.TypeOf(value)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(unsafe.Sizeof(name))
	if value>2222222{
		fmt.Println("3333")
	}else{
		fmt.Println("111")
	}
	for i:=1;i<5;i+=1{
		fmt.Println("123456")
	}	
}