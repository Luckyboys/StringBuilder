package StringBuilder

import (
	"strings"
)

type StringBuilder struct {
	data     []string
	length   uint32
	index    uint32
	maxIndex uint32
}

func GetStringBuilder() *StringBuilder {
	newOne := new(StringBuilder)
	newOne.length = 0
	newOne.index = 0
	newOne.maxIndex = 0
	newOne.alloc()
	return newOne
}

func (this *StringBuilder) Append(word string) *StringBuilder {
	this.alloc()
	this.data[this.index] = word
	this.index++
	this.length += uint32(len([]byte(word)))

	return this
}

func (this *StringBuilder) String() string {

	return strings.Join(this.data, "")
}

func (this *StringBuilder) alloc() {
	if this.index <= this.maxIndex {
		if this.maxIndex == 0 {
			this.maxIndex = 128
		}
		this.maxIndex *= 2
		newSpace := make([]string, this.maxIndex)

		for i := 0; i < len(this.data); i++ {
			newSpace[i] = this.data[i]
		}

		this.data = newSpace
	}
}
