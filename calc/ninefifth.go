// Calculate 95th
package calc

import (
	"github.com/itshosted/mcore/sort"
)

type NineFifth struct {
	Values []int64
}

// Calculate 95th on values (remember values for Total95th)
// please set sort to true unless you've already sorted from small
// too big.
func (n *NineFifth) Add(values []int64, autosort bool) int64 {
	if autosort {
		sort.Ints(n.Values)
	}
	length := len(values)
	fivePercent := int(float64(length) / 100 * 5) /* 5% to steps */

	n.Values = append(n.Values, values...)
	if len(n.Values) == 0 {
		return 0
	}
	nineFifth := values[length-fivePercent-1]

	return nineFifth
}

// Get 95th from summary
func (n *NineFifth) Total95th() int64 {
	length := len(n.Values)
	fivePercent := int(float64(length) / 100 * 5) /* 5% to steps */

	sort.Ints(n.Values)
	if len(n.Values) == 0 {
		return 0
	}
	return n.Values[length-fivePercent-1]
}

// Get new 95th calculator.
func NewNineFifth() *NineFifth {
	return &NineFifth{}
}
