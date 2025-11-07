package calc //package一定要放在最上面

var Age = 20

func Add1(a, b int) int { //Add首字母大写表示公有方法 即可在其他地方进行访问
	return a + b
}

func Sub1(a, b int) int {
	return a - b
}

//故：以后要在其他地方它，就需要将命名的变量或者函数名首字母大写，变成共有变量或者共有方法
