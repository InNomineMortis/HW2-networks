package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"math"
)

/*
		Создаем матрицу преобразования вида:
		{1,0,1,0,1,0,1}
		{0,1,1,0,0,1,1}
		{0,0,0,1,1,1,1}
*/
func matrixCount() [][]int {
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 7)
	}
	row := []int{1,0,1,0,1,0,1}
	matrix[0] = row
	row = []int{0,1,1,0,0,1,1}
	matrix[1] = row
	row = []int{0,0,0,1,1,1,1}
	matrix[2] = row
	return matrix
}

/*
Кодирование сигнала, путем перемножения матрицы преобразования на код Хэминга
и дальнейшего использования кода ошибок для кодирования
*/
func encoding(hammingCode [7]int) [7]int {
	code := make([][]int, 3)
	for i := range code {
		code[i] = make([]int, 7)
	}
	matrix := matrixCount()
	errors := make([]int, 3)
	for i, v := range hammingCode{
		for x := range []int{0,1,2}{
			code[x][i] = v * matrix[x][i]
			errors[x] = int(math.Mod(funk.Sum(code[x]), 2))
		}
	}
	for i,x := range []int{6,5,3}{
		hammingCode[x] = errors[i]
	}
	return hammingCode
}
/*
Вычисление кода Хэмминга
*/
func calcHamming(code [4]int) [7]int {
	var hammingCode = [7]int{code[3],code[2],code[1],0,code[0],0,0}
	for _,v := range []int{6,5,3}{
		hammingCode[v] = int(math.Mod(funk.Sum(hammingCode), 2))
	}
	return hammingCode
}

func decoding(decoded [7]int) [7]int {
	code := make([][]int, 3)
	for i := range code {
		code[i] = make([]int, 7)
	}
	matrix := matrixCount()
	errors := make([]int, 3)
	for i, v := range decoded{
		for _,x := range []int{0,1,2}{
			code[x][i] = v * matrix[x][i]
			fmt.Println(matrix[x])
		}
		fmt.Println(code[1])
	}
	for x := range []int{2,1,0} {
		errors[x] = int(math.Mod(funk.Sum(code[x]), 2))
		fmt.Println(errors[x], code[x])
	}
	for i,x := range []int{6,5,3}{
		decoded[x] = errors[i]
	}
	return decoded
}
func main() {
	code := [4]int{1,1,0,1} //1011
	fmt.Println(encoding(calcHamming(code)))
	cde := encoding(calcHamming(code))
	fmt.Println(decoding(cde))

}