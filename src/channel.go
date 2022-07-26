package main
import "fmt"
import "time"

func display(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Inside display()")
	ch <- 1234
}


//This subroutine pushes numbers 0 to 9 to the channel and closes the channel
func add_to_channel(ch chan int) {
	fmt.Println("Send data")
	for i:=0; i<10; i++ {
	//	time.Sleep(1 * time.Second)
		ch <- i //pushing data to channel
	}
	close(ch) //closing the channel

}

//This subroutine fetches data from the channel and prints it.
func fetch_from_channel(ch chan int) {
	fmt.Println("Read data")
	for {
		//fetch data from channel
		x, flag := <- ch

		//flag is true if data is received from the channel
		//flag is false when the channel is closed
		if flag == true {
			fmt.Println(x)
		}else{
			fmt.Println("Empty channel")
			break
		}
	}
}

func main() {
	ch := make(chan int)
	go display(ch)
	x := <-ch
	fmt.Println("Inside main()")
	fmt.Println("Printing x in main() after taking from channel:",x)


	//invoking the subroutines to add and fetch from the channel
	//These routines execute simultaneously
	go add_to_channel(ch)
	go fetch_from_channel(ch)

	//delay is to prevent the exiting of main() before goroutines finish
	time.Sleep(2 * time.Second)
	fmt.Println("Inside main()")
}