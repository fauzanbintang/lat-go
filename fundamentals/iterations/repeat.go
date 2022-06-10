package iteration

// const repeatCount = 5

func Repeat(char string, num int) string {
	var repeated string

	for i := 0; i < num; i++ {
		repeated = repeated + char
	}

	return repeated
}
