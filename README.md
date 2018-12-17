# Численные методы
Алгоритмы реализованы на следующих языках:

| Язык | Компилятор/Версия | Среда разработки | 
| ------ | ------ | ------ |
| C++ | MinGW / w64 6.0 | CLion 2018.1.5 |
| Go | go / 1.11 darwin_amd64 | CodeRunner 3.0 |
## Содержание
###### Прямые методы решения систем линейных алгебраических уравнений
* Метод Гаусса классический   
* Метод Гаусса с выбором главного элемента в строке 
* Метод Гаусса с выбором главного элемента в столбце  
* Метод Гаусса с выбором главного элемента в матрице
* [Метод квадратного корня](https://github.com/tritonsy/choleskyDecomposition)  `C++`
* [Метод прогонки](https://github.com/tritonsy/thomasAlgorithm) `C++`
###### Итерационные методы решения систем линейных алгебраических уравнений
* [Метод Якоби](https://github.com/tritonsy/JacobiMethod) `C++`
* Метод Зейделя
* [Метод релаксации](https://github.com/tritonsy/relaxationMethod) `C++`
###### Приближение функций
* [Интерполяционный многочлен Ньютона](https://github.com/ochaplashkin/numerical_analysis/tree/master/interpolation_newton) `Go`
* [Интерполяционный кубический сплайн](https://github.com/ochaplashkin/numerical_analysis/tree/master/spline) `Go`
* [Cглаживание по методу наименьших квадратов](https://github.com/ochaplashkin/numerical_analysis/tree/master/ols) `Go`
###### Вычисление интегралов
* [Вычисление  простого интеграла и интеграла с параметром](https://github.com/ochaplashkin/numerical_analysis/blob/master/integral.go) `Go`
###### Решение нелинейных уравнений и систем нелинейных уравнений
* [Метод простой итерации(строки 9-19)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L9) `Go`
* [Метод Ньютона(строки 20-31)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L20) `Go`
* [Метод секущих(строки 32-44)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L32) `Go`
* [Метод вилки(строки 46-63)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L46) `Go`
* [Решение системы уравнений методом Ньютона](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_system(newton).go) `Go`
