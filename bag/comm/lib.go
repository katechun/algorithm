package comm

type Bag struct {
	Numbers []int
}

func (this *Bag) Add(item int) {
	this.Numbers = append(this.Numbers, item)
}

func (this *Bag) Size(items Bag) int {
	return len(items.Numbers)
}

func (this *Bag) IsEmpty(items Bag) bool {
	if len(items.Numbers) == 0 {
		return true
	}

	return false
}
