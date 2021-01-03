# Implementasi

Aplikasi penghitung suara ini diimplementasikan berdasarkan rancangan yang telah dibuat.
Database yang digunakan adalah PostgreSQL.
Untuk pengelolaan database digunakan tool yaitu [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres).
Tool tersebut berguna untuk menghasilkan program yang akan membuat _table_ dan _view_ pada PostgreSQL, serta melakukan import data dari sumber data dalam format CSV ke dalam database.

## Struktur Direktori Program

Berikut ini adalah direktori-direktori yang mengelompokkan file-file dalam aplikasi ini:

- `data`: direktori ini berisi file sumber-sumber data dalam format CSV yang akan diimpor ke dalam database. Seluruh file yang dibutuhkan ada pada direktori ini harus disediakan sebelum proses impor data dilakukan (baca penjelasan dalam [Petunjuk Penggunaan | Setup | Production](../README.md#production)).
- `docs`: direktori ini berisi file dokumentasi dalam format Markdown.
- `helper`: direktori ini berisi kode program untuk membantu tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres) dalam melakukan konversi data sebelum data disimpan ke database. Kode program dalam direktori ini menggunakan bahasa pemrograman Go. Silahkan dibaca komentar-komentar yang ada dalam kode program tersebut untuk mendapatkan penjelasan lebih detail.
- `internal`: direktori ini berisi kode-kode program yang di-_generate_ oleh [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres) untuk melakukan pengelolaan database dan impor data. Direktori ini hanya akan ada setelah langkah inisiasi dilakukan (baca penjelasan dalam [Petunjuk Penggunaan | Inisiasi](../README.md#inisiasi)).
- `node_modules`: direktori ini berisi _library_/_package_ yang dibutuhkan oleh program generator data untuk testing. Direktori ini hanya akan ada setelah langkah testing/ujicoba dilakukan (baca penjelasan dalam [Petunjuk Penggunaan | Testing/Ujicoba](../README.md#testingujicoba)).
- `out`: direktori ini berisi data hasil penghitungan yang diekspor ke dalam format CSV. Direktori ini baru berisi file setelah langkah impor data dilakukan (baca penjelasan dalam [Petunjuk Penggunaan | Impor Data](../README.md#impor-data)).
- `tables`: direktori ini berisi konfigurasi yang dibutuhkan oleh tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres) untuk membuat _table_ pada database dan mengimpor data ke dalamnya.
- `test-data`: direktori ini berisi data-data yang digunakan untuk mengujicoba _table_ dan _view_ yang ada dalam database. Pada awalnya, direktori ini hanya akan berisi sebagian data saja. Untuk melengkapinya silahkan baca penjelasan dalam [Petunjuk Penggunaan | Testing/Ujicoba](../README.md#testingujicoba).
- `test-data-generator`: direktori ini berisi kode-kode program yang digunakan untuk meng-generate data ujicoba dalam direktori `test-data`. Kode program dalam direktori ini menggunakan bahasa pemrograman JavaScript (NodeJS).
- `views`: direktori ini berisi konfigurasi yang dibutuhkan oleh tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres) untuk membuat _view_ pada database.

## Implementasi Impor Data

Sumber data yang berupa file dalam format CSV akan diimpor ke database oleh tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres).
File sumber data disimpan dalam foler `data`.
Konfigurasi impor data ini ditentukan melalui file konfigurasi yang ada dalam folder `tables`.
Setiap file YAML yang ada dalam folder tersebut akan menghasilkan sebuah table pada database dan menentukan bagaimana data dalam format CSV akan diimpor ke dalam _table_ pada database.

Konfigurasi impor data ini memiliki peran penting dalam strategi menjaga kerahasiaan suara dan data pribadi pemilih.

## Implementasi Perlindungan Kerahasiaan Suara

Berdasarkan rancangan (baca penjelasan dalam [Rancangan | Perlindungan Kerahasiaan Suara](rancangan.md#perlindungan-kerahasiaan-suara)), usaha melindungi kerahasiaan data dilakukan melalui implementasi _salt_ dan _hash_.

Nilai _salt_ akan digunakan saat proses impor data ke dalam database.
Oleh karena itu, nilai _salt_ akan dibangkitkan oleh aplikasi sebelum tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres) membangkitkan kode program pengelolaan database.
Nilai _salt_ dibangkitkan oleh aplikasi dan disimpan dalam file `.salt`.
File `.salt` ini tidak boleh dilihat isinya dan harus segera dihapus setelah proses impor data dilakukan.

Penentuan data-data yang akan diberi _salt_ dan di-_hash_ ada pada konfigurasi impor data (dalam file YAML pada direktori `tables`).
Penggunaan _salt_ dan _hash_ memanfaatkan fitur `computeFn` yang disediakan oleh tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres).
Melalui fitur `computeFn` tersebut data yang berasal dari CSV akan dikonversikan ke dalam bentuk nilai _hash_-nya.
Kode program yang digunakan sebagai fungsi `computeFn` dapat dilihat pada file `helper/helper.go` beserta kode testingnya pada file `helper/helper_test.go`.

Berikut ini adalah daftar data/field yang akan di-_salt_ dan di-_hash_ (nama _table_ sama dengan nama file konfigurasi hanya saja tanpa ekstensi `.yaml`):

| Table | Field | Deskripsi |
|-      |-      |-          |
| daftar_pemilih | nim_email_hashed | Data NIM akan diberi akhiran @student.trunojoyo.ac.id kemudian di-_salt_ dan di-_hash_ |
| daftar_pilihan | nim_email_hashed | Data email akan di-_salt_ dan di-_hash_ |

## Implementasi Penentuan Suara Sah dan Penghitungan Suara

Penentuan suara sah dan penghitungan suara pada aplikasi ini diimplementasikan melalui _database view_.
_Database view_ dibangkitkan melalui tool [frm-adiputra/csv2postgres](https://github.com/frm-adiputra/csv2postgres) dengan menggunakan konfigurasi dalam file YAML pada direktori `views`.
Setiap file YAML dalam direktori `views` akan menjadi _database view_ dengan nama yang sama (tanpa ekstensi `.yaml`)

Penentuan suara sah dibuat dalam view yang konfigurasinya ada dalam file `views/vw_pilihan.yaml`.
Setiap suara yang tidak sah akan disertakan informasi alasannya berdasarkan urutan persyaratan yang tidak terpenuhi.

Penghitungan suara memanfaatkan fitur aggregasi pada SQL dan diimplementasikan pada view yang konfigurasinya ada pada file berikut ini:

- `views/vw_pilihan_gubernur.yaml` (untuk Gubernur)
- `views/vw_pilihan_dpm.yaml` (untuk DPM)
- `views/vw_pilihan_hmp.yaml` (untuk himpunan mahasiswa prodi)
