package registry

import "github.com/gerajuarez/wize-academy-go/controller"

func (r *registry) RegisterHello() controller.HelloController {
	return controller.NewHelloController()
}
