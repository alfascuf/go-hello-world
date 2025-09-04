package main

<<<<<<< HEAD
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	results := make([]string, t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)

		n := len(s)
		idx := n - 1
		for j := 0; j < n-1; j++ {
			if s[j] < s[j+1] {
				idx = j
				break
			}
		}

		results[i] = s[:idx] + s[idx+1:]
	}

	for _, res := range results {
		fmt.Fprintln(out, res)
	}
=======
//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	print("Hello World")
>>>>>>> 3d3d44b45693ae873794566897ac0a57eee51f05
}
