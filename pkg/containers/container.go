package containers

import (
	"github.com/negrel/guds/pkg/elements"
)

// Container define the common methods shared by all container data structure.
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []elements.Element
}
