package main

import "fmt"

func makeProgression(start, step int) func() int {

	res := start

	return func() int {

		res += step
		return res

	}
}


func makeFiba() func() int {
  prev, next := 0, 1

  return func() int {
      result := prev      // сначала сохраняем текущее значение
      prev, next = next, prev + next
      return result       // возвращаем старое prev (включая 0)
  }
}

func main() {

	progression := makeFiba()

	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	fmt.Println(progression())
	
}