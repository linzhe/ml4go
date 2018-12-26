package dataset

type DataFrame struct {
	data  []Series
	ncols int
	nrows int
	Err   error
}
