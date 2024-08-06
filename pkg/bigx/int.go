package bigx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
)

type Int struct {
	big.Int
}

func (i *Int) UnmarshalJSON(bytes []byte) error {
	var num interface{}
	if err := json.Unmarshal(bytes, &num); err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	switch val := num.(type) {
	case string:
		intVal := new(big.Int)
		intVal.SetString(val, 10)
		i.Int = *intVal
	case float64:
		intVal := new(big.Int)
		intVal.SetInt64(int64(val))
		i.Int = *intVal
	case bool:
		if val {
			i.Int = *big.NewInt(1)
		} else {
			i.Int = *big.NewInt(0)
		}
	default:
		return fmt.Errorf("unexpected type %T", val)
	}

	return nil
}

func (i *Int) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer([]byte{})
	b.WriteString(`"`)
	b.WriteString(i.String())
	b.WriteString(`"`)
	return b.Bytes(), nil
}
