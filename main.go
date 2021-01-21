package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Get Commandline input and determine project name
	args := os.Args
	var name string
	var module bool = false
	if len(args) == 1 {
		name = "NewGoProject"
	} else {
		name = args[1]
	}
	// Check for additional flags
	if len(args) > 2 {
		for i := 2; i < len(args); i++ {
			// TODO: add switch case if more flags added
			if args[i] == "-m" {
				module = true
			}
		}

	}

	// Check if project folder already exists
	infoDir, _ := os.Stat(name)

	// Create project folder if not existing
	if infoDir == nil {
		err := os.Mkdir(name, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}

		mainPath := name + "/main.go"

		// Check if main.go already exists -> this should never happen
		info, _ := os.Stat(mainPath)

		// Create main.go if not existing
		if info == nil {
			err2 := ioutil.WriteFile(mainPath, []byte("package main\n\nfunc main(){\n\n}"), 0755)
			if err2 != nil {
				log.Fatal(err2)
				return
			}
		}

		// If module flag is set
		if module {

			os.Chdir("./" + name)

			cmd := exec.Command("go", "mod", "init")

			if _, err := cmd.Output(); err != nil {
				fmt.Println("Error: ", err)
				return
			}

		}

		fmt.Printf("Project %s was created.\n", name)

	} else {
		fmt.Printf("Project/Folder %s already exists.\n", name)
	}
}
