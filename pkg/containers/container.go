package containers

import (
	"github.com/negrel/guds/pkg/elements"
)

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []elements.Element
}
