package arrayutil

// Array type
type Array struct {
	Arr []interface{}
}

// Map returns element in a array
func (a *Array) Map(cb func(par interface{}, index int) interface{}) []interface{} {
	var newArr []interface{}
	for i := 0; i < len(a.Arr); i++ {
		val := cb(a.Arr[i], i)
		newArr = append(newArr, val)
	}
	return newArr
}

// Reduce array item
func Reduce(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	return result
}
