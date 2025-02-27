package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan(){
		msg := scanner.Text()
		handleMessage()
	}
}

fun handleMessage(_,any){

}
