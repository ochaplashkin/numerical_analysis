package main

import(
	"fmt"
	"math"
)
const eps float64 = 1.e-8
const n int       = 2 

func f1(x [n]float64)(float64){
	return x[0]*x[0] + x[1]*x[1] - 1
	//return x[0]-x[1]  
	//return math.Exp(x[0]*x[0]-1-x[1]) - x[0]*x[0] + (x[1] + 0.5)*(x[1] + 0.5) - 0.4
}
func df1_dx(x [n]float64)(float64){
	return 2*x[0]
	//return 2*math.Exp(x[0]*x[0] - x[1] -1)*x[0]-2*x[0]
}
func df1_dy(x [n]float64)(float64){
	return 2*x[1]
	//return 2*(x[1]+0.5) - math.Exp(x[0]*x[0] - x[1] -1)
}
func f2(x [n]float64)(float64){
	return math.Exp(x[0]) - x[1]
	//return x[0]*x[0] - x[1]*x[1] - 0.5
}
func df2_dx(x [n]float64)(float64){
	return math.Exp(x[0])
	//return 2*x[0]
}
func df2_dy(x [n]float64)(float64){
	return -1 
	//return -2*x[1]
}
func Cramer(x [n]float64)([n]float64){
	a := [n][n+1]float64{}
	a[0][0] = df1_dx(x)
	a[0][1] = df1_dy(x)
	a[0][2] = -f1(x)
	a[1][0] = df2_dx(x)
	a[1][1] = df2_dy(x)
	a[1][2] = -f2(x)
	
	var def  = (a[0][0]*a[1][1] - a[0][1]*a[1][0])
	var def1 = (a[0][2]*a[1][1] - a[0][1]*a[1][2])
	var def2 = (a[0][0]*a[1][2] - a[0][2]*a[1][0])
	
	result := [n]float64{}
	result[0] = def1/def
	result[1] = def2/def
	return result
}

func main() {
	xk_new := [n]float64{}
	xk := [n]float64{}
	
	xk_new[0] = 1
	xk_new[1] = -1
	var iteration = 0
	for {  
		iteration += 1
		for i := 0; i < n; i++ { xk[i] = xk_new[i] }
		for i := 0; i < n; i++ { xk_new[i] = xk[i] + Cramer(xk)[i] }
		var norm = 0.0
		for i := 0; i < n; i++ {
			var value = math.Abs(xk_new[i] - xk[i])
			if value > norm{
				norm = value
			}
		}
	if (norm < eps){
		break
		}
	}
	fmt.Println("System of nonlinear equations:\n")
	fmt.Println("\t{ F(x,y) = 0")
	fmt.Println("\t{ G(x,y) = 0\n")
	fmt.Println("when\n\tF(x,y) = exp^(x^2-1-y) - x^2 + (y + 0.5)^2 - 0.4")
	fmt.Println("\tG(x,y) = x^2 - y^2 - 0.5\n")
	fmt.Printf("Solve:\n")
	fmt.Printf("\t|    x*     |    y*     |\n")
	fmt.Printf("\t|-----------|-----------|\n")
	if !((xk_new[0] > 0.0) || (xk_new[1] > 0.0)){
		fmt.Printf("\t|     -     |     -     |\n\n\n")
		fmt.Printf("System has not the positive roots: (x > 0) and (y > 0)\n\n")
		fmt.Printf("The current roots:\n")
		fmt.Printf("\t|    x*     |    y*     |\n")
		fmt.Printf("\t|-----------|-----------|\n")
		fmt.Printf("\t| %4.7f | %4.7f |\n\n\n",xk_new[0],xk_new[1])
	}else{
		fmt.Printf("\t| %4.7f | %4.7f |\n\n\n",xk_new[0],xk_new[1])
	}
	
	fmt.Printf("\t|  F(x*,y*)   |  G(x*,y*) |  Iterations  |\n")
	fmt.Printf("\t|-------------|-----------|--------------|\n")
	fmt.Printf("\t|  %4.7f  | %4.7f |     %.d       |\n",f1(xk_new),f2(xk_new),iteration)
	
}