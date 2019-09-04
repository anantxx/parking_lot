package handler

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/parking_lot/constant"
	logs "github.com/parking_lot/log"
	"github.com/parking_lot/repository"
	"github.com/parking_lot/service"
)

func HandleFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		if false == processCommand(input) {
			break
		}
	}
}

func HandleCommand() {
	reader := bufio.NewReader(os.Stdin)
	var (
		input string
		err   error
	)
	for {
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if false == processCommand(input) {
			break
		}
	}
}

func processCommand(input string) bool {
	input = strings.TrimRight(input, "\r\n")
	arguments := []string{}
	logger := logs.NewFmtIOLog()
	for _, s := range strings.Split(input, " ") {
		if s != "" {
			arguments = append(arguments, s)
		}
	}
	if len(arguments) < 1 {
		logger.Log("", ERR_NO_ARGS)
		return true
	}
	if strings.ToLower(arguments[0]) == "exit" {
		return false
	}
	parkingModel := repository.GetInstance()
	parkingRepo := repository.NewParkingRepository(parkingModel)
	slotRepo := repository.NewSlotRepository()
	parkingService := service.NewParkingService(parkingRepo, slotRepo)
	res, err := executeCommand(parkingService, arguments)
	logger.Log(res, err)
	return true
}

func executeCommand(parkingService service.ServicesInstance, arguments []string) (response string, err error) {

	switch arguments[0] {
	case CREATE_PARKING_LOT:
		if len(arguments) < 2 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		totalSlot, err := strconv.Atoi(arguments[1])
		if err != nil {
			return "", err
		}
		return parkingService.InitializeLot(totalSlot)
	case PARK:
		if len(arguments) < 3 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		return parkingService.AllocateSlot(arguments[1], arguments[2])

	case LEAVE:
		if len(arguments) < 2 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		position, _ := strconv.Atoi(arguments[1])
		return parkingService.ReleaseSlot(position)

	case STATUS:
		return parkingService.ShowStatus()
	case REG_NUM_FOR_CAR_WITH_COL:
		if len(arguments) < 2 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		return parkingService.FindRegistationNosByColour(arguments[1])
	case SLOT_NUM_FOR_CAR_WITH_COL:
		if len(arguments) < 2 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		return parkingService.FindAllocatedSlotByColour(arguments[1])
	case SLOT_NUM_FOR_REG_NUM:
		if len(arguments) < 2 {
			return "", ERR_INSUFFICIENT_ARGS
		}
		return parkingService.FindSlotByRegistationNo(arguments[1])
	default:
		return "", ERR_INVALID_COMMAND
	}
	return response, err
}
