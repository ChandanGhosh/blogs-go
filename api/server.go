package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandanghosh/blog/api/database"

	"github.com/chandanghosh/blog/config"

	"github.com/chandanghosh/blog/api/router"
)

func init() {
	config.Load()
	if err := database.Connect(); err != nil {
		log.Fatalln(err)
	}
	//seed.SeedDatabase()
}

// Run ..
func Run() {
	fmt.Printf("\t Listening [::]:%d\n", config.PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), router.New()))
}
