package tmp

import "fmt"

var title = "tmp"

func init() {
	title = "changed"
	fmt.Println("init tmp", title)
}
