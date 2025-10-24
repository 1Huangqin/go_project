package main

import "fmt"

func main() {
	//str1 := "this is str"
	//str2 := "t"
	//flag := strings.HasPrefix(str1, str2)
	//a := 2
	//length := len(2)
	//fmt.Println(flag)
	//a := 128
	//length := unsafe.Sizeof(a)
	//fmt.Println(length)
	//arr1 := [...]string{"php", "css", "html", "golang"}
	//for k, v := range arr1 {
	//	fmt.Printf("k的值：%v,v的值：%v\n", k, v)
	//}
	//求和与平均值
	//arr2 := [...]int{1, 5, 6, 2, 3, 8}
	//sum := 0
	//for i := 0; i < len(arr2); i++ {
	//	sum += arr2[i]
	//}
	//fmt.Printf("和为：%v;平均值为;%.2f", sum, float64(sum)/float64(len(arr2)))
	//数组
	//arr3 := [...]int{1, 2, 3}
	//arr4 := arr3
	//arr3[0] = 55
	//fmt.Println(arr3, arr4)
	////切片
	//arr5 := []int{1, 5, 6}
	//arr6 := arr5
	//arr5[0] = 55
	//fmt.Println(arr5, arr6)
	//多维数组
	var arr = [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//for _, v := range arr {
	//	for _, v2 := range v {
	//		fmt.Println(v2)
	//	}
	//}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
}
