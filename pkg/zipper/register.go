package zipper

import "fmt"

type ZipOps struct {
	Name    string
	Operate func()
}

var AllOperations []ZipOps

func AddOperations(operation ZipOps) {
	AllOperations = append(AllOperations, operation)
}

func Execute() {
	for _, operation := range AllOperations {
		fmt.Printf("\n- Executing operation: %s...\n", operation.Name)
		operation.Operate()
		fmt.Println("Operation executed successfully!")
	}
}
