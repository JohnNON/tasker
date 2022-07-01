/*
Задача о кенгуру
----------------

Есть два кенгуру на оси координат, готовые прыгать в одном направлении (например, в
положительном направлении). Первый кенгуру находится в положении x1 и прыгает на
расстояние v1 за прыжок. Анологично с первым кенгуру, второй находится изначально в
положении x2 и прыгает на v2 за прыжок. По заданным начальным положениям и
скоростям можете ли вы определить окажутся ли они в одном месте в одно и тоже время?

### Входные данные​

Stdin с четырьми целыми числами, разделенными пробелом формата: x1 v1 x2 v2

### Ограничения​

* − 10000 ≤ x1, x2 ≤ 10000
* − 10000 ≤ v1, v2 ≤ 10000

### Формат вывода​

В stdout YES, если кенгуру могут встретится в одном месте в одно и тоже время. И NO в
обратном случае.

### Примеры​

Вход: 0 3 4 2

Результат: YES

Вход: 0 2 5 3

Результат: NO
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal("not enough arguments")
	}

	x1, v1, x2, v2 := getParams(os.Args)

	if !validate(x1, v1, x2, v2) {
		log.Fatal("arguments must be between -10000 and 10000")
	}

	fmt.Println(solveTask(x1, v1, x2, v2))
}

func getParams(args []string) (int, int, int, int) {
	x1, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	v1, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}
	x2, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatal(err)
	}
	v2, err := strconv.Atoi(args[4])
	if err != nil {
		log.Fatal(err)
	}

	return x1, v1, x2, v2
}

func validate(x1, v1, x2, v2 int) bool {
	switch {
	case x1 < -10000 || x1 > 10000:
		return false
	case v1 < -10000 || v1 > 10000:
		return false
	case x2 < -10000 || x2 > 10000:
		return false
	case v2 < -10000 || v2 > 10000:
		return false
	}

	return true
}

func solveTask(x1, v1, x2, v2 int) string {
	if x1 > x2 {
		x1, x2 = x2, x1
		v1, v2 = v2, v1
	}

	if (x1 != x2 && v1 <= v2) ||
		(x2-x1)%(v1-v2) != 0 {
		return "NO"
	}

	return "YES"
}
