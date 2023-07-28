package platform

import (
	"errors"
	"fmt"
	"math/big"
)

func ConvertToExponentDesired(unit interface{}, exponent float64) (*big.Float, error) {
	switch unit.(type) {
	case uint64:
		return new(big.Float).Quo(new(big.Float).SetUint64(unit.(uint64)), big.NewFloat(exponent)), nil
	case *big.Int:
		return new(big.Float).Quo(new(big.Float).SetInt(unit.(*big.Int)), big.NewFloat(exponent)), nil
	case float64:
		return new(big.Float).Quo(new(big.Float).SetFloat64(unit.(float64)), big.NewFloat(exponent)), nil
	case string:
		floatSetted, err := new(big.Float).SetString(unit.(string))
		if !err {
			return nil, errors.New(fmt.Sprintf("error converting to exponent %d", exponent))
		}
		return new(big.Float).Quo(floatSetted, big.NewFloat(exponent)), nil
	}

	return nil, errors.New(fmt.Sprintf("error converting to exponent %d", exponent))
}

func SetDecimalFormat(unit float64, numberOfDecimals string) string {
	format := "%." + numberOfDecimals + "f"
	return fmt.Sprintf(format, unit)
}
