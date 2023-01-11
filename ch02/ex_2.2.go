// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"ch02/lengthconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		i := lengthconv.Inch(t)
		c := lengthconv.Centimeter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			i, lengthconv.InToCM(i), c, lengthconv.CMToIN(c),
		)
	}
}
