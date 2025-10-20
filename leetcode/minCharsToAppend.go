pkg main





func minCharToAppend(source string, target string) int {
	//source: "cbacbb" , target: "abc"

	i, j := 0, 0

	for i < len(src) && j < len(target) {
		if source[i] == target[j] {
			j++
		} 
		i++
	}

	return len(target) - j

}

func main () {
	src := "cbacbb"
	target := "abc"
	fmt.Println(minCharToAppend(src, target))
}