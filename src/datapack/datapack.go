package datapack

// DataPack is a struct that holds the data that is being prepared to be sent over the network.
type DataPack[T any] struct {
	Size int `json:"Size"`
	Data []T `json:"Data"`
}

// CreateTyped creates a new DataPack with the type in brackets. This method is recommended over manual instantiation.
//
// It is equivalent to &DataPack[T]{}.
func CreateTyped[T any]() *DataPack[T] {
	return &DataPack[T]{}
}

// CreateGeneric creates a new DataPack with no specific type. This method is recommended over manual instantiation.
//
// It is equivalent to &DataPack[any]{}.
func CreateGeneric() *DataPack[any] {
	return &DataPack[any]{}
}

// AddData adds data of a certain type to the DataPack.
func (dp *DataPack[T]) AddData(data ...T) {
	dp.Data = append(dp.Data, data...)
	dp.Size = len(dp.Data)
}
