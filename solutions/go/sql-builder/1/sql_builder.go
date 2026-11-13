package sqlbuilder

import "strings"

// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
	if table == "" {
        return ""
    }

	var sb strings.Builder
    sb.WriteString("SELECT ")
    
    if columns == nil || len(columns) == 0 {
        sb.WriteRune('*')
    } else {
        for id, col := range columns {
            sb.WriteString(col)
            
            if id < len(columns) - 1 {
            	sb.WriteString(", ")
            }
        }
    }

    sb.WriteString(" FROM ")
    sb.WriteString(table)
    sb.WriteRune(';')

    return sb.String()
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	if rows == nil || len(rows) == 0 {
        return []string{}
    }

	var sb strings.Builder
	var cols strings.Builder
	var vals strings.Builder
    ret := []string{}
    
    for _, batch := range rows {
        for id, colVal := range batch {
            if colVal.Column == "" || colVal.Value == "" {
                continue
            }

            if id == 0 {
                vals.WriteRune('\'')
            }

            cols.WriteString(colVal.Column)
            vals.WriteString(colVal.Value)

            if id < len(batch) - 1 {
                cols.WriteString(", ")
                vals.WriteString("', '")
            } else {
                vals.WriteRune('\'')
            }
        }

        sb.WriteString("INSERT INTO ")
        sb.WriteString(table)
        sb.WriteString(" (")
        sb.WriteString(cols.String())
        sb.WriteString(") VALUES (")
		sb.WriteString(vals.String())
        sb.WriteString(");")
        ret = append(ret, sb.String())
        sb.Reset()
        cols.Reset()
        vals.Reset()
    }

    return ret
}
