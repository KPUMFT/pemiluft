# Aplikasi Penghitung Suara | Pemilu Mahasiswa UTM

Aplikasi ini adalah aplikasi penghitung suara yang digunakan dalam Pemilu mahasiswa UTM 2020.

Aplikasi ini dikembangkan dengan menggunakan tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres).
Sehingga pada dasarnya, aplikasi ini adalah aplikasi untuk membuat _table_ dan _view_ pada PostgreSQL serta mengimpor data ke dalamnya.

Untuk melakukan penghitungan suara, aplikasi ini akan melakukan hal-hal sebagai berikut:

1. Inisiasi: generate _salt_, generate kode program untuk mengolah database.
1. Membuat berbagai _table_ dan _view_ pada database.
2. Mengimpor berbagai data pemilu ke dalam database.

Hasil penghitungan akan dapat dilihat melalui berbagai _view_ yang ada pada
database.

## Dokumen-Dokumen Pendukung

Dokumen ini adalah petunjuk penggunaan dari aplikasi penghitung suara.
Dokumen-dokumen lainnya terkait aplikasi ini dapat dilihat melalui link berikut:

- [Analisis](docs/analisis.md)
- [Rancangan](docs/rancangan.md)
- [Implementasi](docs/implementasi.md)

## Requirements

- Go 1.15+
- Node 14+
- PostgreSQL 9+

## Petunjuk Penggunaan

### Setup

Clone repository ini.
Buat file `db.yaml` pada project root.
File `db.yaml` berisi konfigurasi koneksi database, dengan contoh pengisiannya dapat dilihat pada file `db.example.yaml`.

Secara default, aplikasi ini berjalan dalam mode testing.
Untuk menjalankan dalam mode production, baca penjelasan pada bab [Production](#production).

Buat _schema_ pada database dengan nama `evote`.

#### Testing/Ujicoba

Selama berjalan dalam mode testing, aplikasi ini akan menggunakan data-data yang ada pada folder `test-data`.
Folder `test-data` pada awalnya tidak akan berisi seluruh data yang dibutuhkan untuk testing.
Untuk meng-generate data testing pada folder tersebut, jalankan perintah berikut:

```bash
# instalasi dependency
# (membutuhkan koneksi internet untuk mengunduh dependency)
npm install

# generate data ujicoba
npm run gen
```

#### Production

Untuk menjalankan dalam mode production ada beberapa hal yang harus dilakukan.

Dalam folder `data` sediakan file-file berikut ini:

- `daftar-pemilih.csv`: berisi daftar pemilih
- `PEMILIHAN UMUM MAHASISWA ELECTRONIC VOTE FAKULTAS TEKNIK UNIVERSITAS TRUNOJOYO MADURA .csv`: berisi data respon yang diunduh dari Google Forms
- `ref-calon-gubernur.csv`: berisi daftar calon gubernur
- `ref-calon-dpm.csv`: berisi daftar calon DPM
- `ref-calon-himatif.csv`: berisi daftar calon HIMATIF
- `ref-calon-himatro.csv`: berisi daftar calon HIMATRO
- `ref-calon-hmti.csv`: berisi daftar calon HMTI

Edit file-file berikut ini dengan cara mengganti isian pada field `csv` yang awalnya merujuk ke file yang ada dalam folder `test-data` dengan merujuk ke file yang ada dalam folder `data`:

- `tables/daftar_pemilih.yaml`
- `tables/daftar_pilihan.yaml`
- `tables/ref_calon_gubernur.yaml`
- `tables/ref_calon_dpm.yaml`
- `tables/ref_calon_himatif.yaml`
- `tables/ref_calon_himatro.yaml`
- `tables/ref_calon_hmti.yaml`

Setelah melakukan hal-hal di atas, silahkan melanjutkan ke proses inisiasi.

### Inisiasi

Sebelum melakukan langkah inisiasi, ikuti terlebih dahulu langkah-langkah yang ada pada bab [Setup](#setup) sesuai dengan mode eksekusi yang diinginkan (testing atau production).

```bash
# Kedua perintah berikut ini membutuhkan koneksi internet
# untuk mengunduh dependency

# Meng-generate kode program untuk mengolah database
go generate .

# Perintah ini akan mengeksekusi program tanpa benar-benar
# melakukan impor data ke database.
# Pastikan perintah ini berjalan tanpa ada error
go run . --dry-run allUp
```

### Impor Data

Jalankan perintah berikut ini untuk mengimpor data ke dalam database.
_Database schema_ yang menjadi target dari impor data ini adalah _schema_ dengan
nama `evote` yang ada pada database yang ditentukan dalam file `db.yaml`.
Apabila database telah berisi data maka, data yang ada akan dihapus terlebih
dahulu.

```bash
go run . allUp
```

Setelah menjalankan perintah di atas, semua data akan berada dalam database dan
hasil penghitungan dapat dilihat melalui berbagai _view_ yang ada.

### Hasil Penghitungan

Setelah impor data dilakukan, maka hasil penghitungan akan dapat dilihat pada
_view_ dalam database atau pada hasil ekspor dari _view_ tersebut yang tersimpan
dalam folder `out`.

Berikut ini adalah _database view_ hasil penghitungan (dalam _schema_ `evote`):

#### `vw_pilihan`

`vw_pilihan` adalah _view_ yang mempresentasikan data pilihan setiap pemilih.
Dalam _view_ ini juga ditampilkan status dari setiap persyaratan sahnya suara.
_View_ ini juga diekspor ke dalam file dalam format CSV dengan nama `out/vw_pilihan.csv`.

Berikut ini adalah penjelasan setiap _field_ yang ada di dalamnya:

- `suara_non_dapil_sah`: akan bernilai `true` bila semua persyaratan suara sah
untuk pilihan Gubernur dan DPM terpenuhi.
- `alasan_suara_non_dapil_tidak_sah`: berisi alasan suara pilihan Gubernur dan DPM ini
dianggap tidak sah.
- `suara_dapil_sah`: akan bernilai `true` bila semua persyaratan suara sah
untuk pilihan pada dapil prodi terpenuhi.
- `alasan_suara_dapil_tidak_sah`: berisi alasan suara pilihan pada dapil prodi ini dianggap
tidak sah.
- `waktu`: waktu pemilih men-submit pilihannya.
- `dalam_waktu_pemungutan`: akan bernilai `true` jika pemilih men-submit
pilihannya dalam rentang waktu yang ditentukan.
- `nim_email_hashed`: nilai _hash_ dari alamat email pemilih (email mahasiswa adalah NIM dengan domain @student.trunojoyo.ac.id)
- `jenis_email`: berisi nilai `student` apabila domain email yang digunakan
pemilih adalah domain email untuk student. Selain itu akan bernilai `bukan student`
- `input_dapil`: nama dapil (prodi) yang diinputkan oleh pemilih.
- `dapil_tercatat`: nama dapil (prodi) yang tercatat pada daftar pemilih.
- `terdaftar`: akan bernilai `true` jika dan hanya jika NIM tercatat pada daftar pemilih dan email yang digunakan adalah email student.
- `dapil_benar`: akan bernilai `true` jika dan hanya jika dapil yang diinputkan oleh pemilih sama dengan dapil yang tercatat pada daftar pemilih.
- `dapil_sesuai`: akan bernilai `true` jika dan hanya jika pemilih memberikan suara pilihan pada dapil yang sesuai dengan dapil yang tercatat pada daftar pemilih.
- `pilihan_gubernur`: berisi paslon Presma yang dipilih.
- `pilihan_dpm`: berisi calon DPM yang dipilih.
- `pilihan_himatif`: berisi calon HIMATIF yang dipilih.
- `pilihan_himatro`: berisi calon HIMATRO yang dipilih.
- `pilihan_hmti`: berisi calon HMTI yang dipilih.

#### `vw_hasil_gubernur` dan `vw_hasil_dpm`

Kedua _view_ ini, masing-masing mempresentasikan data hasil penghitungan suara untuk Gubernur dan DPM.
Dalam _view_ ini akan ditampilkan dua kategori suara, yaitu suara sah dan suara
tidak sah.
Pada kategori suara sah, akan ditunjukkan daftar nama paslon serta jumlah dan persentase suaranya.
Pada kategori suara tidak sah, akan ditunjukkan daftar alasan serta jumlah dan persentase suaranya.
_View_ ini juga diekspor ke dalam file dalam format CSV dengan nama `out/vw_hasil_gubernur.csv` dan `out/vw_hasil_dpm.csv`.

Berikut ini adalah penjelasan setiap _field_ yang ada di dalamnya:

- `sah`: bernilai `true` untuk kategori suara sah, dan `false` untuk kategori suara tidak sah.
- `suara`: untuk kategori suara sah akan berisi nama-nama paslon dan untuk kategori suara tidak sah berisi alasan suara tidak sah.
- `jumlah`: jumlah suara yang masuk dalam kategori
- `persentase_dari_suara_sah`: persentase jumlah suara bila dihitung berdasarkan jumlah suara sah.
- `persentase_dari_seluruh_suara`: persentase jumlah suara bila dihitung berdasarkan jumlah total suara (baik sah maupun tidak).

#### `vw_hasil_hmp`

_View_ ini mempresentasikan data hasil penghitungan suara untuk himpunan mahasiswa prodi.
Dalam _view_ ini, untuk setiap dapil akan ditampilkan dua kategori suara, yaitu suara sah dan suara
tidak sah.
Pada kategori suara sah, akan ditunjukkan daftar nama calon serta jumlah dan persentase suaranya.
Pada kategori suara tidak sah, akan ditunjukkan daftar alasan serta jumlah dan persentase suaranya.

Khusus untuk kategori suara tidak sah dengan alasan "(a) Tidak terdaftar dalam daftar pemilih", tidak akan ditampilkan per dapil.
Karena pemilih yang tidak terdaftar tidak dapat diverifikasi kebenaran dapilnya.

_View_ ini juga diekspor ke dalam file dalam format CSV dengan nama `out/vw_hasil_hmp.csv`.

Berikut ini adalah penjelasan setiap _field_ yang ada di dalamnya:

- `dapil`: dapil prodi
- `sah`: bernilai `true` untuk kategori suara sah, dan `false` untuk kategori suara tidak sah.
- `suara`: untuk kategori suara sah akan berisi nama-nama calon dan untuk kategori suara tidak sah berisi alasan suara tidak sah.
- `jumlah`: jumlah suara yang masuk dalam kategori
- `persentase_dari_suara_sah_dalam_dapil`: persentase jumlah suara bila dihitung berdasarkan jumlah suara sah dalam dapil.
- `persentase_dari_seluruh_suara_dalam_dapil`: persentase jumlah suara bila dihitung berdasarkan jumlah total suara dalam dapil (baik sah maupun tidak).
