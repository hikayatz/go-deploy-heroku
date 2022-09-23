# Tugas Deployment Golang di Heroku
deployment  menggunakan metode [buildpack](https://devcenter.heroku.com/articles/buildpacks)
## Cara setup
- Buat file `Procfile` di root project kemudian lakukan setting `web: <nama-project>`
- tambahkan `// +heroku goVersion go<versi-golang>` di dalam `go.mod`  untuk versi disesuaikan dengan versi golang yang dibuat.
- kemudian push project ke github repository.
- setelah itu buat project app di heroku
- setelah dibuat pindah ke tab `resource` pilih addon database yang sesuai
- setelah addon database terbuat pilih add-on kemudian simapan konfigurasi seperti nama host user dan password yang telah di generate
- di tab `settings` cari form `Config Vars` kemudian isi environment yang tadi disimpan
- pada bagian buildpacks pilih bahasa go
- pindah bagian tab `deploy` pilih method connect to github lalu cari repo  yang sebelumnya dipush kemudian pilih dan koneksikan
- terus centang `automatic deploy`
- setelah itu kembali lagi ke tab `deploy` lalu klik `Deploy Branch`.
- Selesai
