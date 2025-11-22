package utils

import (
	"database/sql"
)

// ScanRowsToMaps converts *sql.Rows to []map[string]interface{}
func ScanRowsToMaps(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for rows.Next() {
		// Create a slice of interface{} to hold values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		// Scan the row
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// Build the map
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]

			// Convert byte arrays to string
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}

			entry[col] = v
		}

		results = append(results, entry)
	}

	return results, rows.Err()
}
