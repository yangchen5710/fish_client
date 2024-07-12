package command

import (
	"bufio"
	"fmt"
	"os"
)

var Input = bufio.NewScanner(os.Stdin)

func Write(message string) string {
	fmt.Print("[poker@" + message + "]$ ")
	Input.Scan()
	return Input.Text()
}

func PrintNotice(msg string) {
	fmt.Println(msg)
}

func Write1(message string) {
	fmt.Print("[poker@" + message + "]$ ")
}

func DeletePreAndSufSpace(str string) string {
	strList := []byte(str)
	spaceCount, count := 0, len(strList)
	for i := 0; i <= len(strList)-1; i++ {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	strList = strList[spaceCount:]
	spaceCount, count = 0, len(strList)
	for i := count - 1; i >= 0; i-- {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	return string(strList[:count-spaceCount])
}
