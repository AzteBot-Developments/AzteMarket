package utils

import "strconv"

func StringToInt64(src string) (*int64, error) {
	i, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func StringToInt(src string) (*int, error) {
	i, err := strconv.Atoi(src)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func StringToFloat64(src string) (*float64, error) {
	i, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
