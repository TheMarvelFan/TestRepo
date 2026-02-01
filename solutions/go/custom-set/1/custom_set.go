package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

import (
    "fmt"
    "strings"
)

// Define the Set type here.
type Set map[string]struct{}

func New() Set {
	return Set(map[string]struct{}{})
}

func NewFromSlice(l []string) Set {
	under := map[string]struct{}{}
    
	for _, key := range l {
        if _, present := under[key]; !present {
            under[key] = struct{}{}
        }
    }

    return Set(under)
}

func (s Set) String() string {
	var sb strings.Builder

    sb.WriteString("{")

    more := false
    
	for key := range s {
        if more {
            sb.WriteString(", ")
        }

        sb.WriteString(fmt.Sprintf("\"%s\"", key))
        
        more = true
    }

    sb.WriteString("}")

    return sb.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, present := s[elem]

    return present
}

func (s Set) Add(elem string) {
    if _, present := s[elem]; !present {
        s[elem] = struct{}{}
    }
}

func Subset(s1, s2 Set) bool {
	if len(s1) == 0 {
        return true
    }

    for key := range s1 {
        if _, present := s2[key]; !present {
            return false
        }
    }

    return true
}

func Disjoint(s1, s2 Set) bool {
	if len(s1) == 0 || len(s2) == 0 {
        return true
    }

    for key := range s1 {
        if _, present := s2[key]; present {
            return false
        }
    }

    return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) == len(s2) {
        for key := range s1 {
            if _, present := s2[key]; !present {
                return false
            }
        }
    } else {
        return false
    }

    return true
}

func Intersection(s1, s2 Set) Set {
	under := map[string]struct{}{}

    for key := range s1 {
        if _, present := s2[key]; present {
            under[key] = struct{}{}
        }
    }

    return Set(under)
}

func Difference(s1, s2 Set) Set {
	under := map[string]struct{}{}
    
	for keyOne := range s1 {
        if _, present := s2[keyOne]; !present {
        	under[keyOne] = struct{}{}
        }
    }

    return Set(under)
}

func Union(s1, s2 Set) Set {
	under := map[string]struct{}{}

    for keyOne := range s1 {
        under[keyOne] = struct{}{}
    }

    for keyTwo := range s2 {
        if _, present := under[keyTwo]; !present {
        	under[keyTwo] = struct{}{}
        }
    }

    return Set(under)
}
