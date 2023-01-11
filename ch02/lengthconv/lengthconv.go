package lengthconv

import "fmt"

type Inch float64
type Centimeter float64

func (i Inch) String() string       { return fmt.Sprintf("%ginch", i) }
func (c Centimeter) String() string { return fmt.Sprintf("%gcm", c) }
