package enum

import (
	"errors"
	//"bytes"
)

//remark the type of json obj int the returned filed
type DataType int
const (
	E_TYPE_USER DataType = iota + 1
	E_TYPE_RELATION

	ERROR_STR = "unknow enum val"
)

var dataTypeDesc = []string{
	"user",
	"relationship",
}

func (s *DataType) String() string {
	if *s >= E_TYPE_USER && *s <= E_TYPE_RELATION {
		return dataTypeDesc[*s-1]
	}
	return "unknown"
}

func (s *DataType)MarshalJSON() ([]byte, error) {
	sVal := s.String()
	if sVal == ERROR_STR {
		return nil, errors.New(ERROR_STR)
	}
	//var buf bytes.Buffer
	//buf.WriteString(sVal)
	//return buf.Bytes(), nil
	return []byte("\""+sVal+"\""), nil
}
