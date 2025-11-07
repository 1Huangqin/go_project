package main

import (
	"encoding/json"
	"errors"
	"fmt"
	T "project01/calc" //T是project01/calc的别名
	"sort"
	"strings"
	"time"
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

func func1(n int) int {
	if n > 1 {
		return n + func1(n-1)
	} else {
		return 1
	}
}

// 闭包：可以让一个变量常驻内存也不会污染全局 它的写法就是：函数里面嵌套一个函数，最后返回里面的函数
func func2() func(y int) int {
	i := 10
	return func(y int) int {
		i += y
		return i
	}
}

// defer
// (a int)，这是命名返回参数,命名返回参数会在函数启动时被初始化为对应类型的零值,
func func3() (a int) { //整个函数体内 a 都是同一个变量,所以defer 执行会直接修改返回值区域中的 a
	defer func() {
		a++
	}()
	return a
}

// 匿名返回值
func func4() int {
	var a int //a=0  局部变量，与返回值无关
	defer func() {
		a++
	}()
	return a
}

// 处理错误
// panic可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效 只用panic会使代码到此处之后下面不会在接着执行，所以与 recover 配套使用
// recover 可将程序恢复回来，继续往后执行
func func5(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var c int
	c = a / b
	return c
}

func Filename(filename string) error {
	if filename == "main.go" {
		return nil
	}
	return errors.New("文件名错误")
}
func func6() {
	defer func() {
		var err = recover() //这里会捕获到 panic(err) 抛出的错误
		if err != nil {
			fmt.Println("给管理员发邮件")
			fmt.Println("错误信息:", err)
		}
	}()
	var err = Filename("xxx.go") // 返回错误
	if err != nil {
		panic(err) // 触发恐慌，跳转到 defer 中的 recover
	}
	fmt.Println("继续执行") // 这行不会执行，因为 panic 中断了执行 即当 panic(err) 被调用时，当前函数的执行会立即停止
}

// 定义结构体
type Person struct {
	Name string
	Age  int
	Sex  string
}

// 结构体方法和接收者
func (q Person) PersonInformation() {
	fmt.Println(q.Name, q.Age)
}

// 修改结构体上的数据
func (w *Person) SetInfor(Name string, Age int) {
	w.Name = Name
	w.Age = Age
}

// 结构体字段类型
type Info struct {
	Name    string
	Age     int
	Hobby   []string
	map1    map[string]string
	address address
}

// 结构体的嵌套
type address struct {
	Street string
	City   string
}

// 结构体的继承 也是属于嵌套
type Animal struct {
	Name string
}

func (a *Animal) run() {
	fmt.Printf("%s会运动\n", a.Name) //%s 就是一个占位标记，告诉 Printf 函数："在这里插入后面提供的字符串参数"。
}

type Dog struct {
	Age int
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.Name)
}

// 结构体转Json序列化 & 结构体标签Tag
type Student struct {
	ID     int `json:"id"` //结构体标签Tag 作用：通过指定tag实现json序列化该字段时的value
	Name   string
	Gender string `json:"gender"`
	Age    int
	//私有属性不能被 json 包访问，公有属性才行 name首字母大写为公有属性 首字母小写为私有属性
}

// 嵌套结构体和 JSON 序列化反序列化
type Class struct {
	Title   string
	Student []Student
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
	//函数的递归 自己调用自己
	fmt.Println(func1(100))
	//闭包
	var fn2 = func2()
	fmt.Println(fn2(10))
	fmt.Println(fn2(20))
	//defer 遇到它从下往上执行
	fmt.Println(func3())
	fmt.Println(func4())
	//错误处理
	fmt.Println(func5(10, 0))
	fmt.Println("结束")
	func6()
	//Format方法格式化输出日期字符串 2006.1.2 15：04
	now := time.Now()
	fmt.Println(now)
	//24小时制
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//12小时制
	fmt.Println(now.Format("2006/01/02 03:04:05"))
	fmt.Println(now.Format("2006/01/02"))
	//获取当前时间戳  什么是时间戳：自1970年1月1日至当前时间的总毫秒数
	time1 := now.Unix() //获取当前时间戳
	fmt.Println(time1)
	time2 := now.UnixNano() //获取当前纳秒时间戳
	fmt.Println(time2)
	//时间戳转换日期 now.Forma()
	var time3 int64 = 1762157796 //秒时间戳
	t := time.Unix(time3, 0)     //秒
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	var time4 int64 = 1762157870458273500 //纳秒时间戳
	ti := time.Unix(0, time4)
	fmt.Println(ti.Format("2006-01-02 15:04:05"))
	//日期转化为时间戳   time.ParseInLocation函数将字符串解析为时间对象
	t1 := "2025-11-03 16:16:36"                                 //目前t1还是一个字符串类型 即 t1还是一个要解析的时间字符串
	timeStandar := "2006-01-02 15:04:05"                        //时间格式模板
	tim, _ := time.ParseInLocation(timeStandar, t1, time.Local) //time.Local使用本地时区
	fmt.Println(tim.Unix())
	//时间间隔  掌握time 包中定义的时间间隔类型 例如：time.Duration 表示 1 纳秒，time.Second 表示 1 秒
	//时间操作函数
	//注：
	//方法	参数类型	返回值类型	用途
	//Add()	Duration	Time	在时间点上加减时间段
	//Sub()	Time	Duration	计算两个时间点之间的时间差
	later := now.Add(time.Hour * 2).Format("15:04:05") //当前时间加2个小时后的时间
	fmt.Println(later)
	before := now.Add(-time.Minute * 3).Format("15:04:05") //当前时间减3分钟后的时间
	fmt.Println(before)
	//Before()方法 返回的是bool类型
	t2 := "2025-11-03 16:16:36"
	t3 := "2025-11-03 17:16:36"
	t4 := "2025-11-03 19:16:36"
	stamp1, _ := time.ParseInLocation(timeStandar, t2, time.Local)
	stamp2, _ := time.ParseInLocation(timeStandar, t3, time.Local)
	stamp3, _ := time.ParseInLocation(timeStandar, t4, time.Local)
	fmt.Println(stamp1.Before(stamp2)) //表示stamp1时间点在stamp2时间点之前 true
	fmt.Println(stamp2.After(stamp3))  //true
	fmt.Println(stamp3.Equal(stamp1))  //false
	//实际开发
	t5 := "2025-11-04 17:16:36"
	stamp4, _ := time.ParseInLocation(timeStandar, t5, time.Local)
	if now.Before(stamp4) {
		fmt.Println("还有时间工作")
	} else {
		fmt.Println("时间结束！")
	}
	//定时器
	h := time.NewTicker(time.Second * 2) //定义定时器并将定时器命名为h
	n := 0
	for i := range h.C {
		fmt.Printf("计时开始：%v\n", i.Format("15:04:05"))
		n++
		if n > 5 {
			h.Stop()
			fmt.Println("计时结束")
			break
		}
	}
	//new()方法给指针变量分配内存空间
	p := new(int)
	*p = 100
	fmt.Println(*p)
	//定义结构体
	var p1 Person
	p1.Name = "lihua"
	p1.Age = 20
	p1.Sex = "男"
	fmt.Println(p1)
	var p2 = Person{
		Name: "lihuahua",
		Age:  30,
		Sex:  "女",
	}
	fmt.Println(p2)
	//结构体方法和接收者
	var q = Person{
		Name: "wangwu",
		Age:  30,
		Sex:  "女",
	}
	q.PersonInformation()
	//修改结构体内信息
	q.SetInfor("李三", 6)
	q.PersonInformation()
	//结构体字段类型
	var q2 Info
	q2.Name = "李二"
	q2.Age = 37
	q2.Hobby = make([]string, 3, 6) //切片Slice 与 map 是引用类型需要分配空间后才可以复制
	q2.Hobby[0] = "写代码"
	q2.Hobby[1] = "追剧"
	q2.Hobby[2] = "睡觉"
	q2.map1 = make(map[string]string)
	q2.map1["address"] = "北京"
	q2.map1["phone"] = "15349756"
	//结构体的嵌套
	q2.address.Street = "一曼路"
	q2.address.City = "宜宾"
	fmt.Printf("%#v\n", q2) //%#v代表详细输出
	d1 := &Dog{
		Age: 3,
		Animal: &Animal{
			Name: "阿奇",
		},
	}
	d1.run()
	d1.wang()
	//结构体对象转Json序列化
	var S = Student{
		ID:     1,
		Name:   "琴子",
		Gender: "女",
		Age:    20,
	}
	fmt.Printf("%#v\n", S)
	var s1, _ = json.Marshal(S) //Marshal函数返回两个值：JSON数据字节和错误信息  _表示忽略错误返回值
	fmt.Printf("%#v\n", s1)
	var s2 = string(s1) //将字节数组转换为字符串
	fmt.Println(s2)     //打印JSON字符串
	//Json字符串转换成结构体对象
	var jsonStr = `{"ID":1,"Name":"琴子","Gender":"女","Age":20}` //jsonStr是Json字符串 数据引入不用"" 否则影响{} 这里我们用反引号``来引入数据
	var student Student
	var err = json.Unmarshal([]byte(jsonStr), &student) //[]byte(jsonStr)这里将jsonStr强制转换成byte类型的切片 &student 注意这里必须传入地址因为我们是要进行在结构体Student中修改数据
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", student)
	fmt.Println(student.Name, student.ID)
	//嵌套结构体和Json序列化
	s3 := Class{
		Title:   "001",
		Student: make([]Student, 0),
	}
	for i := 0; i <= 5; i++ {
		s4 := Student{
			ID:     i,
			Name:   fmt.Sprintf("stu_%d", i), //fmt.Sprintf字符串的拼接 将stu与i拼接成一个字符串
			Gender: "男",
			Age:    20,
		}
		s3.Student = append(s3.Student, s4)
	}
	strBytes, err := json.Marshal(s3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(strBytes))
	}
	//嵌套结构体和Json反序列化
	s5 := `{"Title":"001","Student":[{"id":0,"Name":"stu_0","gender":"男","Age":20},{"id":1,"Name":"stu_1","gender":"男","Age":20},{"id":2,"Name":"stu_2","gender":"男","Age":20},{"id":3,"Name":"stu_3","gender":"男","Age":20},{"id":4,"Name":"stu_4","gender":"男","Age":20},{"id":5,"Name":"stu_5","gender":"男","Age":20}]}`
	var s6 = &Class{}
	err = json.Unmarshal([]byte(s5), s6) //这里s6不用加& 因为var s6 = &Class{}已经s6已经代表了指针
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v", s6)
		fmt.Println(s6.Title)
	}
	sum1 := T.Add1(10, 22)
	fmt.Println(sum1)
	sum2 := T.Sub1(10, 2)
	fmt.Println(sum2)
	fmt.Println(T.Age)
	sum3 := T.Mul1(10, 2) //T.Mul1(10, 2)就等于calc.Mul1(10, 2)
	fmt.Println(sum3)
}

// main包中的init函数 优先于main函数 即init先执行输出之后main函数才执行输出
func init() {
	fmt.Println("main.init....")
}
