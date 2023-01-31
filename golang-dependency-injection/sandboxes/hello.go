package sandboxes

type SayHelloContract interface {
	Hello(name string) string
}

type SayHello struct {
}

type HelloService struct {
	SayHello SayHelloContract
}

func (self *SayHello) Hello(name string) string {
	return "Hello " + name
}

func NewSayHello() *SayHello {
	return &SayHello{}
}

func NewHelloService(sayHello SayHelloContract) *HelloService {
	return &HelloService{SayHello: sayHello}
}
