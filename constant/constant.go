package constant

import "errors"

// Error
var (
	ERR_NO_ARGS           = errors.New("No Argument Present")
	ERR_INSUFFICIENT_ARGS = errors.New("Unsufficient Command Arguments")
	ERR_INVALID_COMMAND   = errors.New("Invalid Command")
	ERR_NO_INITIALIZATION = errors.New("First needs to be creating Parking lot")
	ERR_PARKING_FULL      = errors.New("Sorry, parking lot is full")

	ERR_INVALID_ARGUMENT = errors.New("Invalid Argument")
)

//Message

// Commands
const (
	CREATE_PARKING_LOT        = "create_parking_lot"
	PARK                      = "park"
	LEAVE                     = "leave"
	STATUS                    = "status"
	REG_NUM_FOR_CAR_WITH_COL  = "registration_numbers_for_cars_with_colour"
	SLOT_NUM_FOR_CAR_WITH_COL = "slot_numbers_for_cars_with_colour"
	SLOT_NUM_FOR_REG_NUM      = "slot_number_for_registration_number"
)
