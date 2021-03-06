

#include "stdafx.h"
#include <fstream>
#include <iostream>
//#include <ifstream>
const int n = 6;

int main()
//var fail : textfile;
//f, Al, Bet, x, h, d, M, Mu, L  : array[1..n] of real;
{
	float f[n], Al[n], Bet[n], x[n], h[n], d[n], M[n], Mu[n], L[n];
	int w;
	int i;
	float z, p, S;

	std::ifstream file;
	file.open("input.txt");
	for (int i = 0; i < n; i++) {
		file >> x[i];
		file >> f[i];
	}
	file >> z;
	file.close();

	for (int i = 0; i < n; i++) {
		h[i]  = x[i + 1] - x[i];
	}
	for (int i = 1; i < n; i++) {
		d[i]  = 6 / (h[i - 1] + h[i]) * ((f[i + 1] - f[i]) / h[i] - (f[i] - f[i - 1]) / h[i - 1]);
	}

	L[0] = -1;
	d[0] = 3.3722;

	for (int i = 1; i < n; i++) {
		
	    Mu[i] = -h[i - 1] / (h[i - 1] + h[i]);
		L[i] = -h[i] / (h[i - 1] + h[i]);
	}

	Mu[n] = -0.5;
	d[n] = 3.3614;

	// 3. Решаем систему относительно M[i] методом прогонки

	// Вычисляем коэффициенты Al и Bet
	Al[0] = L[0] / 2;
	Bet[0] = d[0] / 2;

	for (int i = 1; i < n; i++) {

	   p = 2 - Mu[i] * Al[i - 1];
		Al[i] = L[i] / p;
		Bet[i] = (d[i] + Mu[i] * Bet[i - 1]) / p;
	}

	// Вычисляем х
	M[n]  = (d[n] + Mu[n] * Bet[n - 1]) / (2 - Mu[n] * Al[n - 1]);
	for (int i = n-1; i >= 0; i--) {
		M[i] = Al[i] * M[i + 1] + Bet[i];
	}
	// 4. Найдем в каком интервале лежит Z
	for (int i = 0; i < n; i++) {
		if ((x[i] < z) and (z < x[i + 1])) {
		w = i;
		}
	}
	//std::cout << w;

	// 5. Вычислить S(x)

S = M[w] * (x[w + 1] - z)*(x[w + 1] - z)*(x[w + 1] - z) / (6 * h[w]);
S = S + M[w + 1] * (z - x[w])*(z - x[w])*(z - x[w]) / (6 * h[w]);
S = S + (f[w] - M[w] * h[w] * h[w] / 6)*((x[w + 1] - z) / h[w]);
S = S + (f[w + 1] - M[w + 1] * h[w] * h[w] / 6)*((z - x[w]) / h[w]);

	
	std::cout << "z = "<<z <<std::endl<< "   S(z) = "<<S<<std::endl;
	system("pause");

	

    return 0;
}

