//similar to which(1) in unix
//eg: $go run which_one.go which
package main
import (
"fmt"
"os"
"path/filepath"
)

func main(){
	arguments := os.Args
	if len(arguments) == 1{
		fmt.Println("Please provide an argument!")
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullpath := filepath.Join(directory,file)
		fileInfo, err := os.Stat(fullpath)  //returns the info of file
		if err != nil{
			continue
		}
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}
		if mode&0111 != 0{ //is it executable?
			fmt.Println(fullpath)
			return
		}
	}
}