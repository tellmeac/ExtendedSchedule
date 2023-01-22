// Code generated by ent, DO NOT EDIT.

package studygroup

const (
	// Label holds the string label denoting the studygroup type in the database.
	Label = "study_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFacultyName holds the string denoting the facultyname field in the database.
	FieldFacultyName = "faculty_name"
	// Table holds the table name of the studygroup in the database.
	Table = "study_groups"
)

// Columns holds all SQL columns for studygroup fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFacultyName,
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
