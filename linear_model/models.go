package linear_model

import (
	"ml4go/dataset"
)

type DataPoint struct {
	label    float64
	features []float64
}

type Model struct {
	data DataPoint
	ceof []float64
	bias []float64
}

type LinearModel struct {
	fitted bool
}

func NewLinearModel() *LinearModel {
	return &LinearModel{fitted: false}
}

func (linearModel *LinearModel) Fit(dataFrame dataset.DataFrame) error {
	// TODO
	return nil
}

func (linearModel *LinearModel) Predict(dataFrame dataset.DataFrame) (dataset.DataFrame, error) {
	// TODO
	return dataFrame, nil
}

func (model *Model) Fit() error {
	/*
		observations := len(model.data)
		numOfvars := len(model.data[0].Variables)
		// Create some blank variable space
		observed := mat.NewDense(observations, 1, nil)
		variables := mat.NewDense(observations, numOfvars+1, nil)
	*/
	return nil
}
