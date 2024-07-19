package datapack

type DataPack struct {
	Size int   `json:"Size"`
	Data []any `json:"Data"`
}

// CreateNew has the same effect as: DataPack{}.
// It is a more fool-proof way to create a DataPack.
func CreateNew() *DataPack {
	return &DataPack{}
}

// AddData adds data to the TypeDataPack Data slice.
func (dp *DataPack) AddData(data ...any) {
	dp.Data = append(dp.Data, data...)
	dp.Size = len(dp.Data)
}
