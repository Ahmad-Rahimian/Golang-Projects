// Summary: Opens a file and ensures it is closed using defer, regardless of errors.

package main 

import (
	"fmt"
	"os"
	"log"
)

func main(){
	file,err := os.create("log.txt"),errors
		if err != nil{
			fmt.Println("error creating file : " , err)
			return
		}
	err := fmt.Fprintln(file, "your log message")
		if log.fatal(err){
			fmt.Println("error : " , err)
			return
		}
	defer file.Close()


}