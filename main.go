package main

import (
	"fmt"
	"gostl/hashmap"
)

func main() {
	fmt.Println("coucou")

	fmt.Printf("Hash of Bastien : %d\n", hashmap.Hash("Bastien"))
	fmt.Printf("Hash of Misao : %d\n", hashmap.Hash("Misao"))
	fmt.Printf("Hash of Emma : %d\n", hashmap.Hash("Emma"))
	fmt.Printf("Hash of Luna : %d\n", hashmap.Hash("Bastien"))

	hm := hashmap.NewHashMap[string, string](50)

	hm.Put("bastien", "top")

	b, err1 := hm.Get("bastien")
	if err1 != nil {
		fmt.Printf("error 1 : %s\n", err1)
	}

	fmt.Println(b)

	/*
		m, err2 := hm.Get("misao")
		if err2 != nil {
			fmt.Printf("error 1 : %s\n", err1)
		}

		fmt.Println(m)
	*/
}
