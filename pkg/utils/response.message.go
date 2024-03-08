package utils

import (
	"log"

	pb "go.onRunCRM/pkg/api"
)

const ERROR_MESSAGE = "Something went wrong: %v"
const SUCCESS_MESSAGE = "SUCCESS"

func HandleOperationResult(err error) (*pb.Status, error) {
	if err != nil {
		log.Printf(ERROR_MESSAGE, err.Error())
		return &pb.Status{
			Code:    500,
			Message: err.Error(),
		}, err
	}
	return &pb.Status{
		Code:    200,
		Message: SUCCESS_MESSAGE,
	}, nil
}
