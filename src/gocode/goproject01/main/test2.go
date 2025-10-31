package main

import (
	"fmt"
	"sort"
	"strings"
)

func intSum(x ...int) int {
	sum := 0
	fmt.Printf("x的类型为：%T\n", x)
	for _, v := range x {
		sum += v
	}
	return sum
}

func floatSum(x float32, y ...float32) float32 {
	sum := float32(0)
	sum = sum + x
	fmt.Printf("y的类型为%T\n", y)
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func mapSort(x map[string]string) string {
	var keys []string
	for k := range x {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var str string
	for _, v := range keys {
		str += fmt.Sprintf("%v=>%v  ", v, x[v])
	}
	return str
}

// 定义函数类型
type clac func(int, int) int //凡是满足func(int, int) int这个条件的函数都是calc类型，之后函数名都能赋值给clac类型的变量

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 高阶函数
func hh(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 返回函数
type clacType func(int, int) int

func do(o string) clacType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}

	default:
		return nil
	}
}
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
	//切片选择排序
	var sliceA = []int{8, 4, 6, 7, 1, 3, 9}
	for i := 0; i < len(sliceA); i++ {
		for j := i + 1; j < len(sliceA); j++ {
			if sliceA[i] > sliceA[j] {
				temp := sliceA[i]
				sliceA[i] = sliceA[j]
				sliceA[j] = temp
			}
		}
	}
	fmt.Println(sliceA)
	//切片冒泡排序
	var sliceB = []int{6, 1, 3, 8, 7, 9, 5}
	for i := 0; i < len(sliceB); i++ {
		for j := i + 1; j < len(sliceB); j++ {
			if sliceB[i] > sliceB[j] {
				temp := sliceB[i]
				sliceB[i] = sliceB[j]
				sliceB[j] = temp
			}
		}
	}
	fmt.Println(sliceB)

	//sort包排序
	intlist := []int{2, 9, 6, 4, 7, 8, 3, 4}
	float64list := []float64{3.12, 5.36, 5.12, 4.29, 6.17}
	stringlist := []string{"a", "d", "v", "t", "w", "c", "h", "z"}
	sort.Ints(intlist)
	sort.Float64s(float64list)
	sort.Strings(stringlist)
	fmt.Println(intlist)
	fmt.Println(float64list)
	fmt.Println(stringlist)

	//map的使用
	ScoreMap := make(map[string]int, 8)
	ScoreMap["张三"] = 90
	ScoreMap["王五"] = 100
	ScoreMap["王二麻子"] = 92
	fmt.Println(ScoreMap)
	fmt.Println(len(ScoreMap)) //查看长度

	//判断某个健是否存在
	v, ok := ScoreMap["张二"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("此用户不存在")
	}
	//map遍历与删除键值对
	for k1, v := range ScoreMap {
		fmt.Println(k1, v)
	}
	delete(ScoreMap, "张三")
	for k2 := range ScoreMap {
		fmt.Println(k2)
	}
	//map类型的切片，作用允许放多个数据
	var mapSlice = make([]map[string]string, 3)
	for k, v := range mapSlice {
		fmt.Println(k, v)
	}
	if mapSlice[0] == nil {
		mapSlice[0] = make(map[string]string)
		mapSlice[0]["username"] = "张三三"
		mapSlice[0]["password"] = "123458"
		mapSlice[0]["age"] = "56"
	}
	if mapSlice[1] == nil {
		mapSlice[1] = make(map[string]string)
		mapSlice[1]["username"] = "李四四"
		mapSlice[1]["password"] = "12345889"
		mapSlice[1]["age"] = "86"
	}
	for _, v1 := range mapSlice {
		//fmt.Println(k3, v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//如果我们想在map对象中存放一系列的属性的时候，我们就可以把map类型的值定义为切片
	var userHobby = make(map[string][]string)
	userHobby["hobby"] = []string{
		"吃饭",
		"睡觉",
		"追剧",
	}
	userHobby["work"] = []string{
		"php",
		"python",
		"java",
	}
	fmt.Println(userHobby)
	for _, v2 := range userHobby {
		//fmt.Println(k4, v2)
		for k4, v3 := range v2 {
			fmt.Println(k4, v3)
		}
	}
	//map排序
	map1 := make(map[int]int)
	map1[0] = 1
	map1[8] = 27
	map1[3] = 38
	map1[7] = 40
	map1[6] = 523
	fmt.Println(map1)
	for k5, v4 := range map1 {
		fmt.Println(k5, v4)
	}

	//按照map的key升序输出key=>value
	//先把map的key放在切片里
	var keySlice []int
	for k5 := range map1 {
		keySlice = append(keySlice, k5)
	}
	fmt.Println(keySlice)
	sort.Ints(keySlice)
	for _, v5 := range keySlice {
		fmt.Println(v5, map1[v5])
	}
	//统计一个字符串中每个单词出现的次数
	var str = "how do you do"
	var strSlice = strings.Split(str, " ") //转化为切片类型
	fmt.Println(strSlice)
	var strMap = make(map[string]int)
	for _, v6 := range strSlice {
		strMap[v6]++
	}
	fmt.Println(strMap)
	//函数
	var a int
	a = intSum(100, 56, 35, 89, 46, 45, 666)
	fmt.Println(a)

	var b float32
	b = floatSum(3.14, 5.69, 3.28, 6.23)
	fmt.Println(b)
	//函数案例 对k进行升序排序  方法先将k转化为切片
	var B = map[string]string{
		"userName": "lihua",
		"age":      "20",
		"sex":      "man",
		"class":    "five",
	}
	var A string
	A = mapSort(B)
	fmt.Println(A)

	//定义函数类型
	var c clac
	c = add
	fmt.Printf("%T\n", c)
	fmt.Println(c(15, 5))

	d := sub
	fmt.Printf("%T\n", d)
	fmt.Println(d(10, 3))
	//高阶函数的调用
	e := hh(10, 30, add)
	fmt.Println(e)
	//匿名函数
	f := hh(31, 20, func(x, y int) int {
		return x * y
	})
	fmt.Println(f)
	//返回函数的调用
	g := do("+")
	fmt.Println(g(30, 56))
	//匿名函数：就是没有函数名的函数，它可以不同与普通函数，它可以在普通函数中定义运行
	func() {
		fmt.Println("hello world")
	}()
	//加括号是为执行此函数
	var fn = func(x, y int) int {
		return x + y
	}
	fmt.Println(fn(1, 2))
	func(x, y int) {
		fmt.Println(x * y)
	}(10, 20)
}
