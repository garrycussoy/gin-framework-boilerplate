package helpers

import (
	"fmt"
	"strings"
)

// Function to convert a "filter" map into a string which will be fed as a query to database
// Rule of the map is like following:
// 1. If a field contains a list of single string, this means the field name in the map and database are the same
// and compared using "="
// 2. If a field contains a list of two strings, this means the field in the map and database are the same
// and compared using the operator ()
// 3. If a field contains a list of three strings, the format will be: value, operator, field name in database
// The resulted query will be in this order: field name in database - operator - value.
func FilterQueryGenerator(tableName string, filter map[string][]*string) (query string) {
	// Setup base query
	query = fmt.Sprintf(`SELECT * FROM "%s"`, tableName)

	// Loop through each field to build the where clause
	var whereClause []string
	for k, v := range filter {
		// We will add the field to where clause only if it isn't empty
		val := ExtractNullableString(v[0])
		if val != "" {
			if len(v) == 1 {
				// Case list contains one element
				whereClause = append(whereClause, fmt.Sprintf(`"%s" = '%s'`, k, val))
			} else if len(v) == 2 {
				// Case list contains two elements
				whereClause = append(whereClause, fmt.Sprintf(`"%s" %s '%s'`, k, *v[1], val))
			} else if len(v) == 3 {
				// Case list contains three elements
				whereClause = append(whereClause, fmt.Sprintf(`"%s" %s '%s'`, *v[2], *v[1], val))
			}
		}
	}

	// Build the query
	if len(whereClause) == 0 {
		return query // This means the statement will query all records in the table
	}
	query = fmt.Sprintf("%s WHERE %s", query, strings.Join(whereClause, " AND "))

	return query
}
