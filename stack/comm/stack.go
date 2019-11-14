package comm

type Stack struct {
	Numbers []interface{}
}

func (this *Stack) Push(item interface{}) (string, bool) {
	this.Numbers = append(this.Numbers, item)
	return "ok", true
}

func (this *Stack) Pop() (interface{}, bool) {
	length := this.Length()
	var item interface{}
	if length == 0 {
		return item, false
	} else if length == 1 {
		item = this.Numbers[0]
		this.Numbers = []interface{}{}
		return item, true
	} else {
		item = this.Numbers[length-1]
		this.Numbers = this.Numbers[0 : length-2]
		return item, true
	}
}

func (this *Stack) Length() int {
	return len(this.Numbers)

}
