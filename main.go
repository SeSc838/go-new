package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	var name string
	var module bool = false
	if len(args) == 1 {
		name = "NewGoProject"
	} else {
		name = args[1]
	}
	if len(args) > 2 {
		for i := 2; i < len(args); i++ {
			// TODO: add switch case if more flags added
			if args[i] == "-m" {
				module = true
			}
		}

	}

	infoDir, _ := os.Stat(name)

	if infoDir == nil {
		err := os.Mkdir(name, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}

		mainPath := name + "/main.go"

		info, _ := os.Stat(mainPath)

		if info == nil {
			err2 := ioutil.WriteFile(mainPath, []byte("package main\n\nfunc main(){\n\n}"), 0755)
			if err2 != nil {
				log.Fatal(err2)
				return
			}
		}

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
