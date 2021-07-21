package main

import (
	"fmt"
	"github.com/DABronskikh/bgo-3_02.2/pkg/credit"
)

func main() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
}
