package waveshape

import (
	"errors"
	"fmt"
)

type lookupShape struct {
	table *[]float64
}

func NewLookupShape(table *[]float64) (*lookupShape, error) {
	if nil == table {
		return nil, errors.New(
			"waveshaper.NewLookupShape called with nil table",
		)
	}
	n := len(*table)
	if n != nSamples {
		return nil, fmt.Errorf(
			"waveshaper.NewLookupShape: table must have length `%d`",
			n,
		)
	}
	return &lookupShape{table}, nil
}

func (ls lookupShape) Sample(pos float64) float64 {
	i := int(nSamplesF * pos)
	return (*ls.table)[i]
}
