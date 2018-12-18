package main

import(
	"fmt"
	"math"
)
const e float64 = 0.000001

func f(x float64,y float64)(float64){
	return math.Pow(math.Sin(0.6*x),2)*math.Exp(-x*x + 2.5*x - 1.5) - 2*(x-2)*y; 
}
func calculation(h float64, yn float64, tn float64)(float64){
	k1 := f(tn,yn)
	k2 := f(tn+h/2, yn + (h/2)*k1)
	k3 := f(tn+h/2, yn+(h/2)*k2);
	k4 := f(tn+h,yn+h*k3)
	return (k1 + k2 + 2*k2 + 2*k3 + k4)*h/6 + yn
}
func runge_kutta(_t0 float64, T float64, _y0 float64, _h0 float64,N int){
	var y1 float64 = 8.0
	var y2 float64
	var y3 float64 = 0.0
	var h float64  = _h0
	var yn float64 = _y0
	var t float64  = _t0
	
	for{
		if (!(t <= (T-e))){
			break
		}
		for{
			if (!(math.Abs((y1-y3)/15.0) > e)){
				break
			}
			y1 = calculation(h,yn,t)
			h = h/2.0
			y2 = calculation(h,yn,t)
			y3 = calculation(h,y2,t+h)
		}
		t = t + 2.0*h
		yn = y3
		N = N+1
		fmt.Printf("n=\t%f\tt=\t%f\ty=\t%f\n",N,t,yn)
		h = 4*h
	} 
}
func main() {
	fmt.Println("hello world")
	var n int = 0
	runge_kutta(1,6,10,0.5,n)
	fmt.Printf("\n--\n",n)
	
	
}