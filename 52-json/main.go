package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

// struct with tags
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func p(v []byte) {
	fmt.Println(string(v))
}

var sp = fmt.Println

func main() {
	bolB, _ := json.Marshal(true)
	p(bolB)
	intB, _ := json.Marshal(1)
	p(intB)

	fltB, _ := json.Marshal(2.34)
	p(fltB)

	strB, _ := json.Marshal("gopher")
	p(strB)

	slcD := []string{"apple", "pench", "pear"}
	slcB, _ := json.Marshal(slcD)
	p(slcB)

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	p(mapB)

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	p(res1B)

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	p(res2B)

	byt := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)
	var dat map[string]interface{}

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
	sp(dat)

	strs := dat["strs"].([]interface{})
	sp(strs)
	str1 := strs[0].(string)
	sp(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	sp(res)
	sp(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	res1 := response2{}
	json.NewDecoder(strings.NewReader(str)).Decode(&res1)
	fmt.Printf("res1: %+v\n", res1)

}
