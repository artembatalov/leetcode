type ElementStatus int

const (
	Empty = iota
	Filled
	Deleted
)

type Element struct {
	value  int
	status ElementStatus
}

type MyHashSet struct {
	data []Element
	cap  int
	size int
}

func (this *MyHashSet) HashFunc(value int, i int) int {
	return (value%this.cap + i) % this.cap
}

func (this *MyHashSet) Reallocate(increase bool) {
	if increase {
		this.cap = this.cap * 2
	} else {
		if this.cap == 8 {
			return
		}
		this.cap = this.cap / 2
	}
	new_data := make([]Element, this.cap)
	fmt.Println(this.data)
	for k := 0; k < this.cap/2; k++ {
		if this.data[k].status == Filled {

			for i := 0; i < this.cap/2; i++ {
				if new_data[this.HashFunc(this.data[k].value, i)].status == Empty {
					new_data[this.HashFunc(this.data[k].value, i)].value = this.data[k].value
					new_data[this.HashFunc(this.data[k].value, i)].status = Filled
					break
				}
			}
		}
	}
	this.data = new_data
}

func Constructor() MyHashSet {
	return MyHashSet{make([]Element, 8), 8, 0}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	if this.size+this.size/2 > this.cap {

		this.Reallocate(true)
	}

	for i := 0; i < this.cap; i++ {
		if this.data[this.HashFunc(key, i)].status == Empty ||
			this.data[this.HashFunc(key, i)].status == Deleted {
			this.data[this.HashFunc(key, i)].value = key
			this.data[this.HashFunc(key, i)].status = Filled
			this.size += 1

			return
		}
	}

}

func (this *MyHashSet) Remove(key int) {
	if !this.Contains(key) {
		return
	}
	for i := 0; i < this.cap; i++ {
		switch this.data[this.HashFunc(key, i)].status {
		case Filled:
			if this.data[this.HashFunc(key, i)].value == key {
				this.data[this.HashFunc(key, i)].status = Deleted
				if this.size*2+this.size <= this.cap {
					this.Reallocate(false)
				}
			}
		default:
			continue
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for i := 0; i < this.cap; i++ {
		switch this.data[this.HashFunc(key, i)].status {
		case Empty:
			return false
		case Filled:
			if this.data[this.HashFunc(key, i)].value == key {
				return true
			}
		case Deleted:
			continue
		}
	}
	return false
}
