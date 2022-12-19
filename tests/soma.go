package soma

import "fmt"

func main()  {
	x := soma(1, 2, 3)
	fmt.Println(x)

	y := multiplica(1, 2, 3)
	fmt.Println(y)
}

func soma(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}

	return total
}

func multiplica(i ...int) int  {
	total := 1

	for _, v := range i {
		total *= v
	}

	return total
}