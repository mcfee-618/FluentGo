package main
import "fmt"

type Person struct{
	name string
	age  int
}

func main()  {
	fmt.Println("222")
	person:= new (Person)
	person.age=22
	person.name="fei"
	fmt.Println(person.age)
	person1:= &Person {name:"shishi",age:22}
	fmt.Println(person1.name)	
}