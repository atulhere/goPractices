package constant

import (
	"fmt"
)

func PlayWithConstant() {

	const a = 50
	const b int = 40

	fmt.Printf("The type of constants are %T and %T\n", a, b)
}
