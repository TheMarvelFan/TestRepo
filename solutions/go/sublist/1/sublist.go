package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) {
        if checkEqual(l1, l2) {
            return RelationEqual
        }
    } else if len(l1) < len(l2) {
        if checkSubOrSuperlist(false, l1, l2) {
            return RelationSublist
        }
    } else {
        if checkSubOrSuperlist(true, l1, l2) {
            return RelationSuperlist
        } 
    }

    return RelationUnequal
}

func checkEqual(l1, l2 []int) bool {
    for i := 0; i < len(l1); i ++ {
        if l1[i] != l2[i] {
            return false;
        }
    }

    return true;
}

func checkSubOrSuperlist(subOrSuper bool, l1, l2 []int) bool { // subOrSuper = false -> check Sub, subOrSuper= true -> check Super
    if subOrSuper { // super check
        if len(l2) == 0 {
            return true;
        }

        return contains(l2, l1)
    } else { // sub check
        if len(l1) == 0 {
            return true;
        }

        return contains(l1, l2)
    }
}

func contains(subList, superList []int) bool {
    foundFirst := false
    j := 0

    for i := 0; i < len(superList); i ++ {
		if j == len(subList) {
            return true;
        }
        
		if superList[i] == subList[j] {
            if j == 0 {
                foundFirst = true
            }
        } else {
            if foundFirst {
                foundFirst = false;
                i -= j
                j = 0;
            }
        }
        
        if foundFirst {
            j ++
        }
    }

    return j == len(subList);
}
