package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
	"math"
)
// Функция F(x) задана таблицей своих значений. Построить кубический сплайн и вычислить значение функции в указанных точках.
// x    0,1 0,15 0,19 0,25 0,28 0,30
// F(x) 1,1052 1,1618 1,2092 1,2840 1,3231 0,3499
// 2M1+M2=3,3722,  0,5M5+2M6=3,3614.
//d)     x=0,27;

// Тестовая функция:
// n = 3
// x 0 1 2 3
// f(x)  0 1 8 27
// new_x 0.5

var n int
var x []float64
var y []float64
var new_x float64

func lineToArray(line string,n int)([]float64) {
	result := make([]float64,n+1)
	strArray := strings.Fields(line)
	for i := 0; i < len(strArray); i++ {
		result[i],_ = strconv.ParseFloat(strArray[i], 64)
	}
	return result
}

func loadData(path string)(int,[]float64,[]float64,float64) {
	file, _ := os.Open(path)
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	n, _ := strconv.Atoi(lines[0])
	fmt.Println(n)
	x = lineToArray(lines[1],n)
	y = lineToArray(lines[2],n)
	new_x, _ := strconv.ParseFloat(lines[3], 64)
	return n,x,y,new_x
}

func sqr(x float64)(float64){
	return math.Pow(x,2)
}
func cube(x float64)(float64){
	return math.Pow(x,3)
}

func main() {
	
	// 1.вводим данные
	// n - размерность
	// массив х(точки)
	// массив f(x) - значения функций
	// точка x - значение, в котором нужно приближенно найти значение функции
	n,x,y,new_x = loadData("input.txt")
	fmt.Println("Input data:")
	fmt.Println("n: ",n)
	fmt.Println("x(i):",x)
	fmt.Println("f(x[i]): ",y)
	fmt.Println("new_x: ",new_x)
	fmt.Println("-------")
	
	h  := make([]float64,n+1)
	mu := make([]float64,n+1)
	l  := make([]float64,n+1)
	d  := make([]float64,n+1)
		
	var Sx float64
	var dSx float64
		
	// 1. Вычисляем h[i], i=0_N-1
	var i int
	for i = 0; i <= n - 1; i++{
		h[i] = x[i+1] - x[i]
	}	
		
	// 2. Вычисляем mu, l, d i=1_N-1 
	l[0] = 0
	d[0] = 0 
	for i = 1; i <=  n-1; i++{
		mu[i] = h[i-1]/(h[i-1]+h[i])
		l[i]  = h[i]/(h[i-1]+h[i])
		d[i] = ( 6 * ( (y[i+1]-y[i])/h[i] - (y[i]-y[i-1])/h[i-1] ) ) / (h[i-1]+h[i])
	}
	mu[n] = 0
	d[n] = 36 
		

	fmt.Println("h: ",h)
	fmt.Println("mu: ",mu)
	fmt.Println("l: ",l)
	fmt.Println("d: ",d)
		
	// 4. Решить систему для моментов M, методом прогонки 
	a := make([]float64,n+1)
	b := make([]float64,n+1)
	c := make([]float64,n+1) 
	f := make([]float64,n+1)
	M := make([]float64,n+1) //решаем систему для M, а не для y
	alpha := make([]float64,n+1)
	beta := make([]float64,n+1)
		
	b[0] = 0
	c[0] = 2
	f[0] = d[0]
	for i := 1; i <= n; i++ {
		a[i] = -mu[i]
		c[i] = 2
		b[i] = -l[i]
		f[i] = d[i]
	}
	a[n] = 0
	c[n] = 2
	f[n] = d[n]
			
	alpha[0] = b[0]/c[0]
	beta[0]  = f[0]/c[0]
	for i := 1; i <= n - 1; i++ {
		alpha[i] = b[i] / (c[i] - a[i]*alpha[i-1])
		beta[i] = (f[i] + a[i]*beta[i-1]) / (c[i]-a[i]*alpha[i-1])
	}
		
	M[n] = (f[n] + a[n]*beta[n-1]) / (c[n] - a[n]*alpha[n])
	for i := n - 1; i >= 0; i-- {
		M[i] = alpha[i] * M[i+1] + beta[i]
	}
	fmt.Println("M: ",M)
	fmt.Println("--------\n")
		
		 
	// 5. Найти Ix: x принадлежит [ X[i] , X[i+1] ] 
	var Ix int
	for i = 0; i <= n; i++{
		if x[i] <= new_x && new_x <= x[i+1]{
			Ix = i
			break
		}
	}
		
	// 6. Вычислить S(x) и S'(x) - по формулам
	i = Ix 
	Sx = (M[i] * cube(x[i+1]-new_x))/6.0*h[i] + (M[i+1] * cube(new_x - x[i])/6.0*h[i]) + (y[i] - M[i]*sqr(h[i])/6.0) * (x[i+1]-new_x)/h[i] + (y[i+1] - M[i+1]*sqr(h[i])/6.0) * (new_x - x[i])/h[i]    
			
	dSx = (M[i+1]*sqr(new_x-x[i])/2*h[i]) - (M[i]*sqr(x[i+1]-new_x)/2*h[i]) + ((y[i+1]-y[i])/h[i]) - ((M[i+1] - M[i])/6)*h[i]
		
	fmt.Println("S (",new_x,") = ",Sx)
	fmt.Println("S'(",new_x,") = ",dSx)

}