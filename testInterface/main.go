package main
import "fmt"

// 定义Phone的接口
type Phone interface{
	call()
}
// 定义NokiaPhone类
type NokiaPhone struct{

}
// 定义苹果类
type Iphone struct{

}

func (nokiaPhone NokiaPhone)call()  {
	fmt.Println("我是诺基亚")
}

func (iphone Iphone)call()  {
	fmt.Println("我是苹果")
}

func main()  {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(Iphone)
	phone.call()
}