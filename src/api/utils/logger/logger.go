package logger

import "fmt"

func Info(data ...interface{}) {
    fmt.Println(append([]interface{}{"Info"}, data)...)
}

func Warn(data ...interface{}) {
    fmt.Println(append([]interface{}{"Warning"}, data)...)
}

func Error(data ...interface{}) {
    fmt.Println(append([]interface{}{"Error"}, data)...)
}
