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
	row = []int{1,1,0,0,1,1,0}
	matrix[1] = row
	row = []int{1,1,1,1,0,0,0}
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
		for x := range []int{2,1,0}{
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

func decoding(encoded [7]int) ([7]int, []int) {
	code := make([][]int, 3)
	for i := range code {
		code[i] = make([]int, 7)
	}
	matrix := matrixCount()
	errors := make([]int, 3)
	for i, v := range encoded{
		for _,x := range []int{2,1,0}{
			code[x][i] = v * matrix[x][i]
		}
	}
	for i := range []int{2,1,0} {
		errors[i] = int(math.Mod(funk.Sum(code[i]), 2))

	}
	for i,v := range []int{6,5,3}{
		encoded[v] = errors[i]
	}
	return encoded, errors
}

/*
Вычисление количества ошибок и места ошибок
*/
func calcErr(dec [7]int, enc [7]int, err []int) (int, [7]int){
	if funk.Sum(dec) - funk.Sum(enc) >= 2 {
		//fmt.Println("More than 1 err!")
	}else if funk.Sum(dec) - funk.Sum(enc) == 1{
		plErr := err[0] * 1 + err[1] * 2 + err[2] * 8
		//fmt.Println("Error is located in: ", plErr ," symbol")
		if dec[plErr] == 0{
			dec[plErr] = 1
		} else {
			dec[plErr] = 0
		}
		fixErr := 1
		return fixErr, dec
	} else {
		return 0, dec
	}
	return 0, dec
}
/*
Факториал
 */
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
/*
Вывод форматированной таблицы
*/
func tableOut(fixErr int){
	var comb float64
	var ck float64
	fmt.Printf("|%1s|%5s|%2s|%5s|\n", "i", "C(i)k", "Nk", "Ck")
	for i:=1; i < 8; i++ {

		comb = float64(fact(7) / ( fact(7 - i) * fact(i)))
		ck = float64(fixErr) / comb
		fmt.Printf("|%1d|%5.2f|%2d|%5.2f|\n", i, comb, fixErr, ck)
	}

}

func main() {
	code := [4]int{1,1,0,1} //1011
	//fmt.Println(calcHamming(code)) 1010101
	cde := calcHamming(code)
	enc := encoding(cde)
	dec, err := decoding(enc)
	fixErr, _ := calcErr(dec, cde, err) // Проверяем на совпадение раскодированного и исходного сообщений
	fmt.Println("encoded: ", enc)
	fmt.Println("decoded: ", dec)
	tableOut(fixErr)
}