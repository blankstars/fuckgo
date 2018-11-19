package utils

func ConvertSliceString2SliceInterface(s []string) ([]interface{}){
	ss := make([]interface{}, len(s))
	for i, v := range s {
		ss[i] = v
	}
	return ss
}