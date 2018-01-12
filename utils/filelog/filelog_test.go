package filelog

import (
    "testing"
    "fmt"
    "io/ioutil"
    "os"
    "log"
)

func TestFileLog(t *testing.T) {
    fpath := "./log.txt"

    fl, err := NewDefault(fpath)
    if err != nil {
        t.Error("can't open file, error: ", err)
    }
    defer fl.Close()

    logger := fl.logger

    logger.Println("line 1")
    logger.Println("line 2")
    logger.SetFlags(log.Ldate|log.Lmicroseconds)
    logger.Println(2, "line 3")

    buf, err := ioutil.ReadFile(fpath)
    if err != nil {
        t.Error("can't read the file ", err)
    }
    fmt.Printf("%s", buf)

    os.Remove(fpath)
}
