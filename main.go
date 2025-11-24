package main

import "brsonparser/brsoncli"

func main(){
	if err := brsoncli.Run(); err != nil {
		panic(err)
	}
}
