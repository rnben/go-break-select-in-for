package main

func bad() {
	var ch chan string
	for {
		select {
		case <-ch:
			break
		}
	}
}
