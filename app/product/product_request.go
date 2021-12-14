package product

type RequestProduct struct {
	ID          int    `json:"id" validator:"required"`
	NmProduk    string `json:"nmproduk" validator:"required"`
	Harga       int    `json:"harga" validator:"required"`
	Berat       int    `json:"berat" validator:"required"`
	Deskripsi   string `json:"deskripsi" validator:"required"`
	Merek       string `json:"merek" validator:"required"`
	Jenisproduk string `json:"jenisproduk" validator:"required"`
	Tags        string `json:"tags" validator:"required"`
	Color       string `json:"color" validator:"required"`
	Size        int    `json:"size" validator:"required"`
	Bahan       string `json:"bahan" validator:"required"`
	Stok        int    `json:"stok" validator:"required"`
	Gambar      string `json:"gambar" validator:"required"`
	UserID      int    `json:"userid" validator:"required"`
}
