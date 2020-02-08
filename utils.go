package main

import (
	"math"
	"time"
)

//Функция для получения возраста пользователя
func getAgeUser(birthDate int) uint32 {
	ageT := time.Now().Unix() - int64(birthDate)
	ageUser := uint32(time.Unix(ageT, 0).Year() - 1970)
	return ageUser
}

//Функция для округления числа до 5-ти десятичных знаков
func round(x float64) float64 {
	x = x * 100000
	xr := math.Round(x)
	return xr / 100000
}
