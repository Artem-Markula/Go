package main

import "fmt"


func twoSum(nums []int, target int, j int)(a int, b int) {
	sum:= [100]int{}
	for i:=1; i<len(nums); i++{
	sum[i] = nums[i] + nums[j]
	
	if sum[i] == target{
	a = i
	b = j
	
	fmt.Println("Это ", j, "Элемент масива и", i , "Они дают в сумме", target)
	return b, a	
	
	}	
	
	}
	fmt.Println(j, " Элемент не дает в сумме ", target, " Программа печатает 2 нуля")
	return 0, 0

}

func title (){
	fmt.Println("Программа на 1 задание, которая возвращает индексы таких двух чисел")
	fmt.Println("Сумма которых равна заданному значению target")
	fmt.Println("Массив nums может содержать в себе 100 чисел, добавляются через запятую")
	fmt.Println("Программа на каждый элемент складывает другой элемент и сравнивает сумму target")
	fmt.Println("target ему так же можно изменять значения")
	fmt.Println("Если сумма 2 элементов не такая как в target, программа напишет 2 нуля")
	fmt.Println("Программа начинает свою работу\n")
	
	
}

func end (){
	fmt.Println("\nИндексы найдены")
	fmt.Println("Программа заканчивает свою работу")
}

func main() {
	title()
	
	nums:= []int{2, 5, 4, 7}
	target := 9
	
	for j:= 0; j < len(nums); j++{
	
	fmt.Println(twoSum(nums, target, j))
	}
	end()
}
