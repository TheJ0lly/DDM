package datapack

import "encoding/json"

// DataPack is a struct that holds the data as an int that is being prepared to be sent over the network.
type DataPack struct {
	Size int   `json:"Size"`
	Data []int `json:"Data"`
}

// Create creates a new DataPack. This method is recommended over manual instantiation.
func Create() *DataPack {
	return &DataPack{}
}

// AddData adds data of a certain type to the DataPack.
func (dp *DataPack) AddData(data ...int) {
	dp.Data = append(dp.Data, data...)
	dp.Size = len(dp.Data)
}

func (dp *DataPack) ToJson() ([]byte, error) {
	return json.Marshal(dp)
}

func FromJson(b []byte) (*DataPack, error) {
	var dp DataPack
	err := json.Unmarshal(b, &dp)

	return &dp, err

}
