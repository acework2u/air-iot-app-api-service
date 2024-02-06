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
	Brand  string `bson:"brand"`
	Model  string `bson:"model"`
	ItemNo string `bson:"item_no",json:"item_no"`
}
type AcBomResponse struct {
	Year        string     `bson:"year"`
	Btu         int64      `bson:"btu"`
	IndItem     string     `bson:"ind_model"`
	IndModel    string     `bson:"ind_item"`
	OduItem     string     `bson:"odu_item"`
	OduModel    string     `bson:"odu_model"`
	Compressors Compressor `bson:"compressors"`
}

type BomService interface {
	CheckCompressor(indoor string) ([]*AcBomResponse, error)
	CompressorList() ([]*AcBomResponse, error)
}
