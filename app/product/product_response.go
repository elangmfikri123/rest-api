package product

type ResponseProduct struct {
	ID          int    `json:"id"`
	NmProduk    string `json:"nmproduk"`
	Harga       int    `json:"harga"`
	Berat       int    `json:"berat"`
	Deskripsi   string `json:"deskripsi" `
	Merek       string `json:"merek" `
	Jenisproduk string `json:"jenisproduk"`
	Tags        string `json:"tags" `
	Color       string `json:"color" `
	Size        int    `json:"size" `
	Bahan       string `json:"bahan" `
	Stok        int    `json:"stok" `
	Gambar      string `json:"gambar" `
	UserID      int    `json:"userid"`
}

func ProductResponseFormatter(product Product) ResponseProduct {
	formatter := ResponseProduct{
		ID:          product.ID,
		NmProduk:    product.NmProduk,
		Harga:       product.Harga,
		Berat:       product.Berat,
		Deskripsi:   product.Deskripsi,
		Merek:       product.Merek,
		Jenisproduk: product.Jenisproduk,
		Tags:        product.Tags,
		Color:       product.Color,
		Size:        product.Size,
		Bahan:       product.Bahan,
		Stok:        product.Stok,
		Gambar:      product.Gambar,
		UserID:      product.UserID,
	}
	return formatter
}
