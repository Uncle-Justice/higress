package main

import "fmt"

type Calculator struct{}

func (c Calculator) Add(x, y int) int {
    return x + y
}

func (c Calculator) Subtract(x, y int) int {
    return x - y
}

func (c Calculator) Multiply(x, y int) int {
    return x * y
}

func (c Calculator) Divide(x float64, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("Cannot divide by 0")
    }
    return x / y, nil
}

// 示例用法：
func main() {
    calc := Calculator{}
    fmt.Println(calc.Add(3, 5)) // 输出: 8
    fmt.Println(calc.Subtract(10, 7)) // 输出: 3
    fmt.Println(calc.Multiply(4, 6)) // 输出: 24
    result, err := calc.Divide(10.0, 2.0)
    if err != nil {
        fmt.Println(err.Error()) // 在y为0时输出错误信息
    } else {
        fmt.Println(result) // 在正常情况下输出除法结果
    }
}