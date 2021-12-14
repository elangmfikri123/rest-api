package produk

type Produk struct {
	ID          int    //`json:"id" form:"id" query:"id"`
	NmProduk    string //`json:"nmproduk" form:"nmproduk" query:"nmproduk"`
	Harga       int    //`json:"harga" form:"harga" query:"harga"`
	Berat       int    //`json:"berat" form:"berat" query:"berat"`
	Deskripsi   string //json:"deskripsi" form:"deskripsi" query:"deskripsi"`
	Merek       string //`json:"merek" form:"merek" query:"merek"`
	Jenisproduk string //`json:"jenisproduk" form:"jenisproduk" query:"jenisproduk"`
	Tags        string //`json:"tags" form:"tags" query:"tags"`
	Color       string //`json:"color" form:"color" query:"color"`
	Size        int    //`json:"size" form:"size" query:"size"`
	Bahan       string //`json:"bahan" form:"bahan" query:"bahan"`
	Stok        int    //`json:"stok" form:"stok" query:"stok"`
	Gambar      string //`json:"gambar" form:"gambar" query:"gambar"`
	UserID      int    //`json:"userid" form:"userid" query:"userid"`
}

type ProdukImage struct {
	ID        int    //`json:"id" form:"id" query:"id"`
	ProdukID  int    //`json:"produkid" form:"produkid" query:"produkid"`
	FileName  string //`json:"filename" form:"filename" query:"filename"`
	IsPrimary int    //`json:"isprimary" form:"isprimary" query:"isprimary"`
}

// func FetchAllProduk() (Response, error) {
// 	var obj Produk
// 	var arrobj []Produk
// 	var res Response
// }
