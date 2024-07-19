package datapack

type DataPack[T any] struct {
	Size int `json:"Size"`
	Data []T `json:"Data"`
}

func CreateTyped[T any]() *DataPack[T] {
	return &DataPack[T]{}
}

func CreateGeneric() *DataPack[any] {
	return &DataPack[any]{}
}

func (dp *DataPack[T]) AddData(data ...T) {
	dp.Data = append(dp.Data, data...)
	dp.Size = len(dp.Data)
}
