package utils

import "github.com/google/uuid"

//UUID - ...
func UUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)

	}
	return uuid.String()
}
