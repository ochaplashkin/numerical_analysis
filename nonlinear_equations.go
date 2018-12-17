package main

import (
	"fmt"
	"math"
)
const eps float64 = 0.00001

func Iteration(x0 float64, f func(float64) float64, fp func(float64) float64)(float64,float64,uint){
	var n uint = 0
	var x float64 = x0
	var xk float64 = x0 - f(x0)/fp(x0)
	for (math.Abs(xk-x) > eps){
		n += 1
		x = xk
		xk = x - f(x)/fp(x0)
	}
	return xk,f(xk),n
}
func Newton(x0 float64, f func(float64) float64, fp func(float64) float64)(float64,float64,uint){
	var n uint = 0
	var x float64 = x0
	var xk float64 = x0 - f(x0)/fp(x0)
	for (math.Abs(xk-x) > eps){
		n += 1
		x = xk
		xk = x - f(x)/fp(x)
	}
	return xk,f(xk),n 
}

func Split(x0 float64, x1 float64, f func(float64) float64 )(float64,float64,uint){
	var n uint = 1
	var xk1 float64 = x0
	var x float64 = x1
	var xk float64 = x - f(x)*(x-xk1)/(f(x)-f(xk1))
	for (math.Abs(xk-x) > eps){
			n += 1
			xk1 = x
			x = xk
			xk = x - f(x)*(x-xk1)/(f(x)-f(xk1))
		}
	return xk,f(xk),n
}

func Plug(a float64,b float64,f func(float64) float64 )(float64,float64,uint){
	var c float64 = 0
	var n uint = 0
	var ak float64 = a
	var bk float64 = b
	var xk float64
	for (math.Abs(bk-ak) > 2*eps){
		n += 1
		c = ak + (bk - ak) / 2
		if ((f(ak)*f(c)) < 0){
			bk = c
		}else{
			ak = c
		}
	}
	xk = (ak + bk) / 2
	return xk,f(xk),n	
}

func f1(x float64)(float64){ //а) 6 – 5sh x = 0;
	return 6 - 5 * math.Sinh(x)
}
func df1(x float64)(float64){
	return -5 - (math.Pow(math.E,x*1.0) + math.Pow(math.E,-x*1.0))/2
}
func f2(x float64)(float64){ //b) 2 – xex = 0;
	return 2 - x * math.Pow(math.E,x*1.0)
}
func df2(x float64)(float64){
	return  - math.Pow(math.E,x*1.0) * (1 + x)
}
func f3(x float64)(float64){ //x^5 + 0,6x^4 + 0,7x^3 + 0,8x^2 + 0,9x + 1 =0.
	return math.Pow(x,5.0) + 0.6 * math.Pow(x,4.0) + 0.7 * math.Pow(x,3.0) + 0.8*x*x + 0.9*x + 1
}
func df3(x float64)(float64){
	return 5*math.Pow(x,4.0) + 2.4*x*x*x + 2.1*x*x + 1.6*x + 0.9
}

func main() {  
	fmt.Printf("         ____________________    __________\n")  
	fmt.Printf(" _______/ f(x) = 6 - 5*sh(x) \\__/ x0 = 0.5 \\_____\n")
	fmt.Printf("|                                               |\n")   
	fmt.Printf("|           |           |           |           |\n")
	fmt.Printf("|  Method   |  x*-koren |    f(x)   |      n    |\n")
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
	_xk, _f, _n := Plug(0,1.1,f1)
	fmt.Printf("|   Plug    | %4.7f | %4.7f |     %.d    |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
	
	_xk, _f, _n = Iteration(0.5,f1,df1)
	fmt.Printf("| Iteration | %4.7f | %4.6f |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
	
	_xk, _f, _n = Split(0.5,0.1,f1) 
	fmt.Printf("|  Split    | %4.7f | %4.6f  |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
	
	_xk, _f, _n = Newton(0.5,f1,df1)
	fmt.Printf("|  Newton   | %4.7f | %4.6f  |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|___________|___________|___________|___________|\n")
	
	//------------------------------------------------------------------- 
	
	fmt.Printf("                __________________\n")  
	fmt.Printf(" ______________/ f(x) = 2 – x*e^x \\_____________")
	fmt.Printf("|                                               |\n")   
	fmt.Printf("|           |           |           |           |\n")
	fmt.Printf("|  Method   |  x*-koren |    f(x)   |      n    |\n")
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
	_xk, _f, _n = Plug(0.75,1,f2)
	fmt.Printf("|   Plug    | %4.7f | %4.7f |     %.d    |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
		
	_xk, _f, _n = Iteration(0.9,f2,df2)
	fmt.Printf("| Iteration | %4.7f | %4.6f |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
		
	_xk, _f, _n = Split(0.75,0.76,f2) 
	fmt.Printf("|  Split    | %4.7f | %4.6f  |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
		
	_xk, _f, _n = Newton(0.75,f2,df2)
	fmt.Printf("|  Newton   | %4.7f | %4.6f  |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|___________|___________|___________|___________|\n")
	
	// ------------------------------------------------------------------
	
	fmt.Printf("  _______________________________________________\n")   
	fmt.Printf(" / f(x) = x^5 + 0,6x^4 + 0,7x^3 + 0,8x^2 + 0,9x + 1 \n") 
	fmt.Printf("|                                               |\n")   
	fmt.Printf("|           |           |           |           |\n")
	fmt.Printf("|  Method   |  x*-koren |    f(x)   |      n    |\n") 
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
	_xk, _f, _n = Plug(-1,-0.9,f3)
	fmt.Printf("|   Plug    | %4.7f | %4.7f |     %.d    |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
		
	_xk, _f, _n = Iteration(-0.9,f3,df3)
	fmt.Printf("| Iteration | %4.7f | %4.6f |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
		
	_xk, _f, _n = Split(-0.9,-0.91,f3) 
	fmt.Printf("|  Split    | %4.7f | %4.6f  |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|-----------|-----------|-----------|-----------|\n")
		
	_xk, _f, _n = Newton(-0.9,f3,df3)
	fmt.Printf("|  Newton   | %4.7f | %4.6f  |     %.d     |\n",_xk,_f,_n)
	fmt.Printf("|___________|___________|___________|___________|\n")
	
}