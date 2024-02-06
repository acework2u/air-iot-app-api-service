package smartapp

type CompressorInfo struct {
	Brand  string `bson:"brand"`
	Model  string `bson:"model"`
	ItemNo string `bson:"item_no"`
}
type OduInfo struct {
	Year       string         `bson:"year"`
	Item       string         `bson:"item"`
	Model      string         `bson:"model"`
	Compressor CompressorInfo `bson:"compressor"`
}
type Compressor struct {
	Brand  string `bson:"brand" json:"brand"`
	Model  string `bson:"model" json:"model"`
	ItemNo string `bson:"item_no",json:"item_no"`
}
type AcProduct struct {
	Year        int32      `bson:"year"`
	Btu         int64      `bson:"btu"`
	IndItem     string     `bson:"ind_model"`
	IndModel    string     `bson:"ind_item"`
	OduItem     string     `bson:"odu_item"`
	OduModel    string     `bson:"odu_model"`
	Compressors Compressor `bson:"compressor"`
}

type BomRepository interface {
	Compressor(indoor string) ([]*AcProduct, error)
	Compressors() ([]*AcProduct, error)
}
