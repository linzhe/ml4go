package algo

import (
	"ml4go/core"
)

type Clustering interface {
	Init(params map[string]string)
	Cluster(dataset core.DataSet)
}
