package produk

type Produk struct {
	ID          int
	NmProduk    string
	Harga       int
	Berat       int
	Deskripsi   string
	Merek       string
	Jenisproduk string
	Tags        string
	Color       string
	Size        int
	Bahan       string
	Stok        int
	Gambar      string
	UserID      int
}

type ProdukImage struct {
	ID        int
	ProdukID  int
	FileName  string
	IsPrimary int
}
