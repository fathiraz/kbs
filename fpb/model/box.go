package model

type (

	// Box struct to hold box data
	Box struct {
		Cakes  uint64
		Apples uint64
	}
)

// GetCakes function to get cakes data
func (b Box) GetCakes() uint64 {
	return b.Cakes
}

// GetApples function to get apples data
func (b Box) GetApples() uint64 {
	return b.Apples
}

// GetBoxEvenly function to get evenly cakes and apples each box
func (b Box) GetBoxEvenly() uint64 {

	var (
		num1 = b.GetApples()
		num2 = b.GetCakes()
	)

	// get greatest common divisor by cakes and apples, using euclidean algorithm
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}

	return num1
}

// GetBoxCakes function to get box cakes data
func (b Box) GetBoxCakes() uint64 {
	return b.Cakes / b.GetBoxEvenly()
}

// GetBoxApples function to get box apples data
func (b Box) GetBoxApples() uint64 {
	return b.Apples / b.GetBoxEvenly()
}
