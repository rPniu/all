package models

import "github.com/rPniu/all/pkg/checkin"

type UserCheckIn struct {
	ID uint
	*UserCore
	*checkin.CheckInManager
}
