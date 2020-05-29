// Experimental
// try simple storage package
// the type have: KV, Table

package storage

//the any type of the data
type Any interface{}

// the Kv style data
type Kv map[Any]Any

// the list of table
type Table []Any

const (
	LiteralInt    = "int"   // golang type: int
	LiteralFloat  = "float" // golang type: float64
	LiteralNumber = "number"
	LiteralBool   = "bool"   // true/True; false/False
	LiteralString = "string" //'string' or "string" or string
	LiteralNull   = "null"
)

//the Literal variable. this is value from string
type Literal string
