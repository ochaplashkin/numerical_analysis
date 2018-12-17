package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)
//
var n int
var x []float64
var f []float64
var new_x float64
var m int
var alpha [][]float64

func lineToArray(line string,n int)([]float64) {
	result := make([]float64,n)
	strArray := strings.Fields(line)
	for i := 0; i < len(strArray); i++ {
		result[i],_ = strconv.ParseFloat(strArray[i], 64)
	}
	return result
}

func loadData(path string)(int,[]float64,[]float64,float64,int) {
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
	new_x, _ := strconv.ParseFloat(lines[3], 64)
	m, _ := strconv.Atoi(lines[4])
	return n,x,f,new_x,m
}

func Gauss(a [][]float64, n int)([]float64){
	var _summ float64
	
	x := make([]float64, n+1)
	
	for m := 0; m <= n - 1; m++{
		for i := m + 1; i <= n; i++{
			for j := m + 1; j <= n + 1; j++{
				a[i][j] = a[i][j] - a[i][m]/a[m][m]*a[m][j] 
			} 
		}
	}
	x[n] = a[n][n+1]/a[n][n]
	
	for i:= n - 1; i >= 0; i--{
		_summ = 0
		for j:= i + 1; j <= n; j++{
			_summ = _summ +a[i][j]*x[j]
		}
		x[i] = (a[i][n+1] - _summ)/a[i][i]
	}
	return x 
	
}
func main() {
	// вводим данные 
	// n - размерность ( количество точек - 1 )
	// массив х(точки)
	// массив f(x) - значения функций
	// точка new_x - значения, в которых нужно приближенно найти значение функции
	// m - степень
	n,x,f,new_x,m = loadData("input.txt")
	fmt.Println("Input data")
	fmt.Println("n: ",n)
	fmt.Println("x(i):",x)
	fmt.Println("f(x[i]): ",f)
	fmt.Println("new_x: ",new_x)
	fmt.Println("m:",m)
	fmt.Println("--------")
	
	alpha := make([][]float64, m+1)
	for i := range alpha {
		alpha[i] = make([]float64, m+2) // n + 1 - обычная, n + 2 - расширенная, совмещенная с вектором B
	}
	fmt.Println("test",alpha)
	//2. Вычислить значение alpha[j][k], B[j]
	// alpha[j][k] = sum(from i=0 to n)x[i]^(j+k) , j,k = 0,m
	// beta[j] = sum(from i=0 to n)x[i]^j*f(x[i]), j = 0,m
	var _summ float64

	for j := 0; j <= m; j++ {
		for k := 0; k <= m+1; k++{
			_summ = 0
			if (k == m+1){
				for i := 0; i <= n; i++{
					_summ += math.Pow(x[i], float64(j))*f[i]
				}
			}else{
				for i := 0; i <= n; i++ {
					_summ += math.Pow(x[i], float64(j+k))
				}
			}
			alpha[j][k] = _summ	 
		}
	}
	fmt.Println("result",alpha)
	//3. Решить СЛАУ методов Гаусса/Квадратного корня ( ответ - вектор a[i])
	a := Gauss(alpha,m)
	fmt.Println(a)
	
	//4. Находим приближенную функцию gm(x) = sum(from k=0 to m)(a[k] * x^k)
	_summ = 0
	for k := 0; k <= m; k++ {
		_summ += a[k] * math.Pow(new_x, float64(k))
	}
	fmt.Println("g(",new_x,"): ",_summ)
}