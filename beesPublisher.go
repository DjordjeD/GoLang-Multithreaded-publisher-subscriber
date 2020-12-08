package main

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
			n := 0 // A local variable is only visible to one goroutine.
			n++
			ch <- n // The data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another goroutine.
	n++
	fmt.Println(n) // Output: 2
}()

func main()
{
    
}