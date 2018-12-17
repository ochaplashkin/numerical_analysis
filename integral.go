package main

import (
	"fmt"
	"math"
)

const eps1 float64 = 0.00001
const eps2 float64 = 0.00001

const a1   float64 = 0 // TODO: test = 0
const b1   float64 = math.Pi / 2 // TODO: test = math.Pi / 2 

const a2   float64 = 0
const b2   float64 = math.Pi / 2

const nu   float64 = -0.01
const c    float64 = 0.5
const d    float64 = 1.5
const m    int 	   = 10
const H    float64 = (d-c)/float64(m)
 

func f1(x float64)(float64){
	//TODO: test = return math.Sin(x)
	//return math.Sqrt(math.Exp(x)-1)
	return math.Sin(x)
}
func phi(z float64)(float64){
	return math.Sin(z)
}
func psi(z float64)(float64){
	return (1 - math.Pow(z,2.0))/(1 + math.Pow(z,2.0))
}
func f2(t float64,x float64)(float64){ 
//	var arg = t / (1 + math.Pow(x,2.0)) + nu * x
//	return phi(arg)*psi(x)
	return math.Sin(x)
	//TODO: test = return math.Sin(x)
}
func u(n int)(float64){ // без параметра
	var h = (b1-a1)/float64(n)
	var x float64 = a1
	var sum float64 = 0
	for i := 0; i < n; i++{
		sum = sum + f1(x) 
		x = x + h
	}
	return sum * h 
}

func u_t(t float64, n int)(float64){ //с параметром
	var h float64 = (b2-a2)/float64(n)	
	var x float64 = a2
	var sum float64 = 0
	for i := 0; i < n; i++{ 
		sum = sum + f2(t,x)
		x = x + h 
	}
	return sum*h	
}

func main() { 
	
	//---- calc u = f1 dx intergal-----
	var n int = 1
	var u_1 float64 = u(n) // n = 1
	n = n * 2
	var u_2 float64 = u(n) // n = 2
	
	for math.Abs(u_1 - u_2) > eps1{
		u_1 = u_2
		n = n * 2
		u_2 = u(n)
	}
	fmt.Printf("u = %3.3f\nn = %d\neps = %f\n---------------------------\n",u_2,n,eps1)
	//TODO: result for test: 1 
	
	// ------- calc F(t) -----
	fmt.Printf(" t[i]\t   F(t)\t\tn\n")
	fmt.Printf("------|------------|-------\n")
	
	t := make([]float64,m+1) 
	for j := 0; j <= m; j++{
		t[j] = c + float64(j)*H
		
		var n int = 1
		var u_1 float64 = u_t(t[j],n) // n = 1
		n = n * 2
		var u_2 float64 = u_t(t[j],n) // n = 2

		
		for math.Abs(u_1 - u_2) > eps2{
			u_1 = u_2
			n = n * 2
			u_2 = u_t(t[j],n)
		}
		fmt.Printf("%2.3f | %8.8f | %d\n",t[j],u_2,n)		
	} 
	fmt.Printf("------|------------|-------\n")
	fmt.Printf("eps = %f\n",eps2)
	//TODO: result for test: 1 
}