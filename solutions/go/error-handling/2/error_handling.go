package erratum

const TRANSIENT_ERROR string = "transient error"
const FROB_ERROR string = "frob error"
const UNKNOWN_ERROR string = "unknown error"

func Use(opener ResourceOpener, input string) (retErr error) {
	res, resErr := opener()
    
    if resErr != nil {
        if getErrorType(resErr) == TRANSIENT_ERROR {
            return Use(opener, input)
        } else {
            return resErr
        }
    }

    defer func() {
        frobbingErr := recover()
        
        if frobbingErr != nil {
            if getErrorType(frobbingErr) == FROB_ERROR {
                frobbingError := frobbingErr.(FrobError)
                res.Defrob(frobbingError.defrobTag)
                retErr = frobbingError.inner
            } else {
            	retErr = frobbingErr.(error)
            }
        }
        
        res.Close()
    }()

    res.Frob(input)
    return
}

func getErrorType(errorInst interface{}) string {
    switch errorInst.(type) {
    case TransientError:
        return TRANSIENT_ERROR
    case FrobError:
        return FROB_ERROR
    default:
        return UNKNOWN_ERROR
    }
}
