package main

import (
	"github.com/watshim-b/aws-sdk-wrapper-go-ddd/handler"
)

func main() {
	r := handler.NewRouter()
	r.SetConfig()
	r.SetRouting()
	r.Run(":8080")
}
