package flatten

func Flatten(nested any) []any {
	ret := []any{}

    items, ok := nested.([]any)

    if !ok {
        return []any{nested}
    }

    for _, inner := range items {
        switch inner.(type) {
        	case []any:
            	elems := Flatten(inner)
            	ret = append(ret, elems...)
    
            default:
            	if inner != nil {
            		ret = append(ret, inner)
                }
        }
    }

    return ret
}
