package lengthconv

func InToCM(i Inch) Centimeter { return Centimeter(i * 2.54) }

func CMToIN(c Centimeter) Inch { return Inch(c / 2.54) }
