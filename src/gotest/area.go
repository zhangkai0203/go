package gotest

import "fmt"

var num int = 10;

func Area(w int, h int) int {
    return w * h;
}

func Add()  {
    num += 1
}

func Del()  {
    num -= 1
}

func GetNum() (num int)  {
    return num
}

func main() {

    fmt.Println("aaa")
}
