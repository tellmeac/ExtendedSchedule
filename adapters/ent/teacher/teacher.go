// Code generated by ent, DO NOT EDIT.

package teacher

const (
	// Label holds the string label denoting the teacher type in the database.
	Label = "teacher"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the teacher in the database.
	Table = "teachers"
)

// Columns holds all SQL columns for teacher fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
