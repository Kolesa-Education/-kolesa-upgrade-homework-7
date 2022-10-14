# Usage:
```bash
go run cmd/main.go
```

# kolesa-upgrade-homework-7
Дедлайн: 14.10.22 пятница 23:59  
Домашнее задание по Go

### Реализуйте калькулятор на Go

В домашнем задании предлагаем вам реализовать собственный консольный калькулятор на go.

Мы будем использовать базовый синтаксис Go.  
Достаточно сделать с помощью fmt пакета. Необходимо сделать подобный калькулятор тому, который был реализован на PHP
(в этом задании нет необходимости в веб сервере, достаточно сделать консольный калькулятор)

Требования к основной части задания:
* Программа должна принимать операнды и оператор(-ы) с помощью fmt.
* Выполнять простые арифметические операции (сложение, вычитание, умножение и деление)
* Валидировать данные перед вычислениями
   * Проверка оператора (пример: выражение "5 q 2" не должно выполняться, так как "q" не является оператором)
   * Проверка переданных переменных (являются ли они численными (int, float))
* Добавление дополительной функциональности (степень числа, сравнение чисел, процент числа)

Требования к дополнительной части задания:
Необходимо добавить в реализацию калькулятора, возможность понимать приоритеты операций и вложенные скобки. Пример: 1+3*(1+2/1-(2-1))  должно вывести 7

* За основную часть: максимальная оценка 80%.
* За дополнительную часть: максимальная оценка 20%.

UPD: ссылочка на установку go https://go.dev/dl/