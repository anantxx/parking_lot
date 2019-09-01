package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/parking_lot/constant"
	"github.com/parking_lot/repository"
	"github.com/parking_lot/service"
)

func HandleFile(fileName string) {

}

func HandleCommand() {
	fmt.Println("\n Enter Command:")
	reader := bufio.NewReader(os.Stdin)
	var command string
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		arguments := []string{}
		for _, s := range strings.Split(input, " ") {
			if s != "" {
				arguments = append(arguments, s)
			}
		}
		if len(arguments) < 1 {
			fmt.Println(ERR_NO_ARGS)
			continue
		}
		command = arguments[0]
		if strings.ToLower(command) == "exit" {
			break
		}
		parkingModel := repository.GetInstance()
		parkingRepo := repository.NewParkingRepository(parkingModel)
		slotRepo := repository.NewSlotRepository()
		parkingService := service.NewParkingService(parkingRepo, slotRepo)
		res, err := ProcessCommand(parkingService, arguments)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
}

func ProcessCommand(parkingService service.ServicesInstance, arguments []string) (response string, err error) {

	switch arguments[0] {
	case CREATE_PARKING_LOT:
		if len(arguments) < 1 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		totalSlot, err := strconv.Atoi(arguments[1])
		if err != nil {
			return "", err
		}
		return parkingService.InitializeLot(totalSlot)
	case PARK:
		if len(arguments) < 2 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		return parkingService.AllocateSlot(arguments[1], arguments[2])

	case LEAVE:
		if len(arguments) < 1 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		position, _ := strconv.Atoi(arguments[1])
		return parkingService.ReleaseSlot(position)

		/*case STATUS:
			model.ShowStatus()
		case REG_NUM_FOR_CAR_WITH_COL:
			model.FindRegistationNosByColour()
		case SLOT_NUM_FOR_CAR_WITH_COL:
			model.FindAllocatedSlotByColour()
		case SLOT_NUM_FOR_REG_NUM:
			model.FindSlotByRegistationNos()
		default:
			return "", ERR_INVALID_COMMAND
		*/
	}
	return response, err
}
