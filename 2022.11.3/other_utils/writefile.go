package other_utils

import (
    "fmt"
    "os"
)

func Utils() {
    //WriteToFile("./test.txt","12345")
    // trunc 直接覆盖原文件所有内容
    WriteToFile("./test.txt","89") // 89
}

func WriteToFile(fileName string, content string) error {
    f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("file create failed. err: " + err.Error())
    } else {
        // offset
        //os.Truncate(filename, 0) //clear
        n, _ := f.Seek(0, os.SEEK_END)
        _, err = f.WriteAt([]byte(content), n)
        fmt.Println("write succeed!")
        defer f.Close()
    }
    return err
}
