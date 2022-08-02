package utils

import "github.com/bxcodec/faker/v3"

func CreateFakerData[T any]() (T, error) {
	var fakerData *T = new(T)
	err := faker.FakeData(fakerData)
	if err != nil {
		return *fakerData, err
	}

	return *fakerData, nil
}
