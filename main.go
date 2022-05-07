package main

import (
	"github.com/joho/godotenv"
	"github.com/z9fr/greensforum-backend/cmd/server"
	_ "github.com/z9fr/greensforum-backend/docs"
)

// @title Green Forum Backend
// @version 1.0
// @description REST api documentation for green forum
// @termsOfService http://dasith.works

// @contact.name API support
// @contact.url http://greenforum.com/support
// @contact.email z9fr@proton.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.staging.green.dasith.works
// @BasePath /v2

//@securityDefinitions.apikey JWT
//@in header
//@name Authorization

func main() {
	godotenv.Load(".env")
	// start and run the server
	server.Start()
	//algorithm.Start()

}
