package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/mauleyzaola/issue-tracker/server/application"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
)

//configuration files are parsed and application is initialized
func main() {
	jsonFile := flag.String("config", "config.json", "Config file. Make a copy of config.json.sample")
	flag.Parse()

	log.Printf("Configuration file:%s\n", *jsonFile)

	rootChDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if err = os.Chdir("../"); err != nil {
		log.Fatal(err)
	}

	if rootChDir, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	app := application.ParseConfiguration(filepath.Join(rootChDir, "server", *jsonFile), rootChDir)

	//get the base root path of the application
	app.Setup.RootChDir = rootChDir

	err = app.BootstrapApplication()
	if err != nil {
		log.Fatal(err)
	}

	//register all api routes
	app.Router()

	graceful.PostHook(func() {
		log.Println("Application is closing now... releasing resources")
		log.Println(app.Db == nil)
	})
	goji.Serve()
}
