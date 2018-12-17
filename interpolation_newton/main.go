package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция f(x) задана таблицой 1. Найти значения  функции для  для указанных значений аргумента
//для функции  h(x): x=0,47113,  x=0,47531,  x=0,48398, х=0,48675.
//x 1,0 1,1 1,2 1,3 1,4 1,5 1,6 1,7 1,8 1,9
//f 0,6827 0,7287 0,7699 0,8064 0,8385 0,8664 0,8904 0,9109 0,9281 0,9426
// n = количество точек - 1

var n int
var x []float64
var f []float64
var new_x []float64
var df [][]float64

func lineToArray(line string,n int)([]float64) {
	result := make([]float64,n)
	strArray := strings.Fields(line)
	for i := 0; i < len(strArray); i++ {
		result[i],_ = strconv.ParseFloat(strArray[i], 64)
	}
	return result
}

func loadData(path string)(int,[]float64,[]float64,[]float64) {
	file, _ := os.Open(path)
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	n, _ := strconv.Atoi(lines[0])
	x = lineToArray(lines[1],n+1)
	f = lineToArray(lines[2],n+1)
	new_x = lineToArray(lines[3], n+1)
	return n,x,f,new_x
}

func printDf(){
	var i int
	for i = 0; i <= n; i++{
		fmt.Println(df[i])
		fmt.Println("")
	}	
}
func main() {
	// вводим данные
	// n - размерность
	// массив х(точки)
	// массив f(x) - значения функций
	// массив new_x - значения, в которых нужно приближенно найти значение функции
	n,x,f,new_x = loadData("input.txt")
	fmt.Println("Input data")
	fmt.Println("n: ",n)
	fmt.Println("x(i):",x)
	fmt.Println("f(x[i]): ",f)
	fmt.Println("new_x: ",new_x)
	fmt.Println("--------") 

	df := make([][]float64, n+1)
	for i := range df {
		df[i] = make([]float64, n+1)
	}

	//задаём первый столбец матрицы df 
	var i int 
	var j int
	 for j = 0; j <= n; j++{
		df[j][0] = f[j] // df[0,j] = f(x[j]), j=0,n
	}
	// рассчитываем df[i,j]
	for i = 1; i <= n; i++{
		for j = 0; j <= n - i; j++{
			df[j][i] = (df[j][i-1] - df[j+1][i-1])/(x[j]-x[j+i]) 
		}
	}
	var gn float64
	gn = df[0][n]
	for k:=0; k <= n; k++{	
		for i = n; i >= 1; i--{
			gn = df[0][i-1] + (new_x[k] - x[i-1])*gn
		}
		fmt.Println("g(",new_x[k],"):",gn)
	}

}