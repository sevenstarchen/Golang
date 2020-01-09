package main

func main() {

	c := make([][]int, 6)

	for i := 0; i < 6; i++ {
		c[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			c[i][j] = j
		}
	}

	for i := range c {
		println(c[i][1])
	}
}
