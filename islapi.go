package main

// "fmt"

func islapi(s string) bool {
	ch := []rune(s)
	hash := make(map[rune]int)
	for i := 0; i < len(ch)/2; i++ {
		hash[ch[i]]++
	}
	for i := len(ch)/2 + 1; i < len(ch); i++ {
		hash[ch[i]]--
	}
	for _, val := range hash {
		if val != 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	// s := "hello"
// 	// t := "el"
// 	str := "radar"
// 	var res bool = islapi(str)
// 	fmt.Println(res)
// }
