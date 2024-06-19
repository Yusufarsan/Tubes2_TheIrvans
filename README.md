# Tugas Besar 2 IF2211 Strategi Algoritma

![screencapture-localhost-8000-2024-06-19-12_22_26](https://github.com/Yusufarsan/Tubes2_TheIrvans/assets/113454186/cf417d11-5179-4e50-ba48-7d7fc859e8ae)

## Deskripsi
WikiRace atau Wiki Game adalah permainan yang melibatkan Wikipedia, sebuah ensiklopedia online gratis yang dikelola oleh berbagai sukarelawan di seluruh dunia, di mana pemain memulai dari sebuah artikel Wikipedia dan harus menelusuri artikel lain di Wikipedia (dengan mengklik tautan di dalam setiap artikel) untuk mencapai tujuan yang telah ditentukan. Artikel dalam waktu sesingkat-singkatnya atau dengan klik (artikel) paling sedikit.
Website kami bertujuan untuk mencari rute artikel terpendek dari satu artikel ke artikel lainnya. Setelah User memasukkan judul artikel awal dan judul artikel tujuan dan menge-klik tombol “Find!”, akan ditampilkan:
1. Jumlah artikel yang diperiksa

2. Jumlah artikel yang lolos

3. Rute eksplorasi artikel

4. Waktu pencarian, dan

5. Visualisasi grafik eksplorasi rute

User dapat memilih metode pencarian dengan IDS (Iterative Deepening Search) atau BFS (Breadth First Search). IDS akan melakukan pencarian dengan meningkatkan nilai depth-cutoff menggunakan rangkaian DFS (Depth First Search) hingga ditemukan solusi. Dalam mencari suatu node dalam suatu graf, DFS akan melakukan pencarian dengan cara memperluas root child pertama dari pohon pencarian yang dipilih dan masuk lebih dalam lagi hingga node target ditemukan, atau hingga menemukan node yang tidak memiliki anak. Sementara itu, BFS akan memulai pencarian grafik dari node akar dan kemudian menjelajahi semua node tetangganya. Kemudian setiap node terdekat tersebut menelusuri node tetangga yang belum diperiksa, begitu seterusnya hingga node target ditemukan.
Selain itu, User juga dapat mengatur batasan dalam pencarian rute eksplorasi artikel sehingga rute terpendek yang ditampilkan tidak hanya satu, tetapi semua rute terpendek dari hasil pencarian juga akan ditampilkan.


## Cara Menjalankan

1. Pindah ke direktori src

    ```bash
    cd src
    ```

2. Jalankan perintah berikut, jika gagal, jalankan docker desktop terlebih dahulu

    ```bash
    docker-compose up
    ```

3. Kunjungi `localhost:8000`

4. Website sudah bisa digunakan

## Anggota

| NAMA ANGGOTA         | NIM      |
|----------------------|----------|
| Shafiq Irvansyah     | 13522003 |
| Ahmad Naufal Ramadan | 13522005 |
| Yusuf Ardian Sandi   | 13522015 |
