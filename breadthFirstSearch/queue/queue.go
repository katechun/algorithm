package queue

type Queue struct {
	queue []string
}

func (this *Queue) Put(name string) {
	this.queue = append(this.queue, name)
}

func (this *Queue) Pop() string {
	if this.queue == nil {
		return ""
	}
	this.queue = this.queue[1:]
	res := this.queue[0]
	return res
}

func (this *Queue) Size() int {
	sz := len(this.queue)
	return sz
}
