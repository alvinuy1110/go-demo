package main
import "fmt"
import "time"
import "sync"
import "strconv"
import "math/rand"

//declare a mutex instance
var mu sync.Mutex

//declare count variable, which is accessed by all the routine instances
var count = 0

//copies count to temp, do some processing(increment) and store back to count
//random delay is added between reading and writing of count variable
func process(n int) {
	//loop incrementing the count by 10
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		//lock starts here
		mu.Lock()
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
		//lock ends here
		mu.Unlock()
	}
	fmt.Println("Count after i="+strconv.Itoa(n)+" Count:", strconv.Itoa(count))
}

func main() {
	//loop calling the process() 3 times
	for i := 1; i < 4; i++ {
		go process(i)
	}

	//delay to wait for the routines to complete
	time.Sleep(25 * time.Second)
	fmt.Println("Final Count:", count)
}

