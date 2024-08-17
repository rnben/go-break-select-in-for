package p

func bad(ch <-chan bool) {
	for {
		select {
		case <-ch:
			break // want "break statement inside select statement inside for loop"
		}
	}
}

func good(ch <-chan bool) {
OUTER:
	for {
		select {
		case <-ch:
			break OUTER
		}
	}
}
