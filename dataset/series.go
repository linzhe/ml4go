package dataset

type Type string
type ElementValue interface{}

type Series struct {
	Name     string
	elements Elements
	t        Type
	Err      error
}

type Elements interface {
	Elem(int) Element
	Len() int
}

type Element interface {
	// Setter method
	Set(interface{})

	// Comparation methods
	Eq(Element) bool
	Neq(Element) bool
	Less(Element) bool
	LessEq(Element) bool
	Greater(Element) bool
	GreaterEq(Element) bool

	// Accessor/conversion methods
	Copy() Element     // FIXME: Returning interface is a recipe for pain
	Val() ElementValue // FIXME: Returning interface is a recipe for pain
	String() string
	Int() (int, error)
	Float() float64
	Bool() (bool, error)

	// Information methods
	IsNA() bool
	Type() Type
}
