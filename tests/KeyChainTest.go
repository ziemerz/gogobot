package tests
import (
	"github.com/ziemerz/gogobot/utility"
	"fmt"
)

func TestKeyChain() {
	fmt.Println(utility.GetKey("cat"))
	fmt.Println(utility.GetKey("dog"))
}
