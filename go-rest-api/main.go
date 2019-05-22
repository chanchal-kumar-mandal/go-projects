package main
 
import (
	"github.com/go-rest-api/app"
	"github.com/go-rest-api/config"
)
 
func main() {
	config := config.GetConfig()
 
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
