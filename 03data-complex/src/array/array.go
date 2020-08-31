package main
import "fmt"

func main()  {
	fmt.Println("hello-world")
	var array1 [10] uint32
	array1[0]=221
	array1[1]=2
	var array2 = [5] uint32{2,3,4,5,6}
	fmt.Println(array1[0])
	fmt.Println(array2[0])
	for index,value := range array2{
		fmt.Printf("%d %d\n",index,value)
	}
	

}