package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregan"
	fmt.Println(states)

	california := states["CA"]

	if len(california) == 0 {
		fmt.Println("The value is not set")
	} else {
	  fmt.Println("The value is ", california)
	}

	delete(states,"OR")
	fmt.Println(states)

	states["CA"] = "California"
	states["NY"] = "New York"

	fmt.Println(states)

	for k,v := range states {
		fmt.Printf("%v: %v\n",k,v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k 
		i++
	}
	sort.Strings(keys)
	fmt.Println("Sorted Keys: \n",keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
