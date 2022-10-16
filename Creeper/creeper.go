package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)
func CreateFile() {
var num int
	for i :=0; i<=10; i++{
num++
//you can also change the name of the file
FileName := fmt.Sprint("creeper", num, ".txt")
os.Create(FileName)
    }
}

func RandNum() int {
rand.Seed(time.Now().UnixNano())
min := 0
max := 10
rando := rand.Intn(max - min + 1) + min
return rando
}

func MoveFile(destination string) {
   //change the file path to whatever the fuck you want m8 
	oldpath := "creeper"
	newpath := fmt.Sprint("/home/zerocool/", destination, "/creeper")
	os.Rename(oldpath, newpath)
}

func main () {
r := RandNum()
	switch r {
	case 1:
CreateFile()
MoveFile("Home")
	case 2:
CreateFile()
MoveFile("Desktop")
	case 3:
CreateFile()
MoveFile("Documents")
	case 4:
CreateFile()
MoveFile("Music")
	case 5:
CreateFile()
MoveFile("Pictures")
	case 6:
CreateFile()
MoveFile("Videos")
	case 7:
CreateFile()
MoveFile("Downloads")
	case 8:
CreateFile()
MoveFile("Favorites")
	case 9:
CreateFile()
MoveFile("Recent")	
	}
}


