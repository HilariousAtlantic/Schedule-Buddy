package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "import":
			importDatabase()
		case "createdb":
			if docker {
				createTables()
			} else {
				createDatabase()
			}
		case "dropdb":
			if docker {
				fmt.Println("Not available in docker, please delete database manually")
			} else {
				deleteDatabase()
			}
		case "resetdb":
			if docker {
				fmt.Println(`Not available in docker, please delete and create the database manually`)
			} else {
				deleteDatabase()
				createDatabase()
			}
		case "gpas":
			importGPAs()
		default:
			startServer()
		}
	} else {
		startServer()
	}
}
