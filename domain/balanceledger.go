package domain

import "github.com/shopspring/decimal"

// DailyLoadLimit is the amount of times a customer can load per day
var DailyLoadLimit = 3

// DailyAmountLimit is the daily limit that can be loaded
var DailyAmountLimit = decimal.NewFromInt(5000)

// WeeklyAmountLimit is the daily limit that can be loaded
var WeeklyAmountLimit = decimal.NewFromInt(20000)
