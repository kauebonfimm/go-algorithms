package inverted_list

type Users struct {
	Id int
}

func Invert(arr []Users) []Users {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}
