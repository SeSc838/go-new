package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func getProjectName(args []string) string {
	if len(args) == 1 {
		return "NewGoProject"
	}
	return args[1]
}

func checkFlags(args []string) (module bool) {
	module = false

	// Check for additional flags
	if len(args) > 2 {
		for i := 2; i < len(args); i++ {
			// TODO: add switch case if more flags added
			if args[i] == "-m" {
				module = true
			}
		}
	}
	return module
}

func createMainGo(pathName string) {
	// Check if main.go already exists -> this should never happen
	info, _ := os.Stat(pathName)

	// Create main.go if not existing
	if info == nil {
		errMain := ioutil.WriteFile(pathName, []byte("package main\n\nfunc main(){\n\n}"), 0755)
		if errMain != nil {
			log.Fatal(errMain)
			return
		}
	}
}

func generateModule(module bool, name string) {
	// If module flag is set
	if module {

		os.Chdir("./" + name)

		cmd := exec.Command("go", "mod", "init")

		if _, err := cmd.Output(); err != nil {
			fmt.Println("Error: ", err)
			return
		}

	}
}

func main() {
	// Get Commandline input and determine project name
	args := os.Args
	name := getProjectName(args)
	module := checkFlags(args)

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

		createMainGo(mainPath)
		generateModule(module, name)

		fmt.Printf("Project %s was created.\n", name)
	} else {
		fmt.Printf("Project/Folder %s already exists.\n", name)
	}
}
