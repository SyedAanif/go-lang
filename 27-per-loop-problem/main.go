package main

func main() {
	strings := []string{"a", "b", "c"}
	var prints []func()
	// prior to go 1.22, the loop variable was re-used across iterations
	// and as this was a closure, the value was updated
	// so the print loop was always giving the last value(c)
	// from go-1.22, this was fixed by re-writing the per-loop iteration variables
	for _, v := range strings {
		// if we had to fix this in 1.21, re-assign to a local variable beyond closures
		// _v := v
		// prints = append(prints, func() { println(_v) })
		prints = append(prints, func() { println(v) })
	}

	for _, p := range prints { // used to print last value(c) always prior to go-1.22
		p()
	}
}
