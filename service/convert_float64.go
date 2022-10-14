package service

import "strconv"

func convertFloat(num string) (float64, error) {
	resFloat, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}
	return resFloat, nil
}
