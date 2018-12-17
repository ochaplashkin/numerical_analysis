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
* [Метод квадратного корня](https://github.com/tritonsy/choleskyDecomposition)
* [Метод прогонки](https://github.com/tritonsy/thomasAlgorithm)
###### Итерационные методы решения систем линейных алгебраических уравнений
* [Метод Якоби](https://github.com/tritonsy/JacobiMethod)
* Метод Зейделя
* [Метод релаксации](https://github.com/tritonsy/relaxationMethod)
###### Приближение функций
* [Интерполяционный многочлен Ньютона](https://github.com/ochaplashkin/numerical_analysis/tree/master/interpolation_newton)
* [Интерполяционный кубический сплайн](https://github.com/ochaplashkin/numerical_analysis/tree/master/spline)
* [Cглаживание по методу наименьших квадратов](https://github.com/ochaplashkin/numerical_analysis/tree/master/ols)
###### Вычисление интегралов
* [Вычисление  простого интеграла и интеграла с параметром](https://github.com/ochaplashkin/numerical_analysis/blob/master/integral.go)
###### Решение нелинейных уравнений и систем нелинейных уравнений
* [Метод простой итерации(строки 9-19)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L9)
* [Метод Ньютона(строки 20-31)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L20)
* [Метод секущих(строки 32-44)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L32)
* [Метод вилки(строки 46-63)](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_equations.go#L46)
* [Решение системы уравнений методом Ньютона](https://github.com/ochaplashkin/numerical_analysis/blob/master/nonlinear_system(newton).go)
