## **Soal Latihan: 1**  
### **Program Pengurutan Daftar Rumah dengan Selection Sort**

**Penjelasan Program:**  
1. **Tujuan**: Program ini digunakan untuk mengurutkan daftar angka yang dimasukkan, yang mewakili nomor rumah, secara ascending menggunakan algoritma **Selection Sort**. Program ini akan mengurutkan data rumah berdasarkan urutan nomor rumah yang diterima.  
2. **Alur Program**:  
   - Program menerima input sejumlah data, yang masing-masing data berisi jumlah nomor rumah dan daftar nomor rumah itu sendiri.
   - Setiap daftar nomor rumah akan diurutkan secara ascending menggunakan algoritma **Selection Sort**.
   - Setelah diurutkan, program akan menampilkan hasil pengurutan nomor rumah.
3. **Output**:  

**Contoh Input:**
```
3
5 2 1 7 9 13
6 189 15 27 39 75 133
3 4 9 1
```

**Contoh Output:**
```
1 2 7 9 13
15 27 39 75 133 189
1 4 9
```

---

## **Soal Latihan: 2**  
### **Program Pemisahan dan Pengurutan Angka Ganjil dan Genap**

**Penjelasan Program:**  
1. **Tujuan**: Program digunakan untuk mengurutkan dua jenis angka (ganjil dan genap) secara terpisah. Angka ganjil diurutkan secara ascending dan angka genap diurutkan secara descending.  
2. **Alur Program**:  
   - Program menerima input sejumlah angka dan memisahkannya menjadi dua kategori: angka ganjil dan angka genap.
   - Program akan mengurutkan angka ganjil secara ascending dan angka genap secara descending menggunakan **Selection Sort**.
   - Setelah itu, program akan menampilkan angka ganjil terlebih dahulu, diikuti dengan angka genap.
3. **Output**:  

**Contoh Input:**
```
3
5 2 1 7 9 13
6 189 15 27 39 75 133
3 4 9 1
```

**Contoh Output:**
```
1 13 12 8 2
15 27 39 75 133 189
8 4 2
```

---

## **Soal Latihan: 3**  
### **Program Penghitungan Median dengan Insertion Sort**

**Penjelasan Program:**  
1. **Tujuan**: Program ini digunakan untuk menghitung median dari deretan angka yang dimasukkan oleh pengguna. Program akan mengurutkan data menggunakan algoritma **Insertion Sort** dan menghitung median.  
2. **Alur Program**:  
   - Input angka satu per satu. Program akan terus menerima input hingga angka `0` dimasukkan.
   - Jika angka `0` dimasukkan, program akan menghitung median dari angka yang telah dimasukkan dengan mengurutkan data terlebih dahulu.
   - Jika jumlah data ganjil, median adalah angka tengah. Jika jumlah data genap, median dihitung sebagai rata-rata dua angka tengah.
   - Program berhenti jika angka `-5313` dimasukkan.
3. **Output**:  

**Contoh Input:**
```
7 23 11 0 5 19 2 29 3 13 17 0 -5313
```

**Contoh Output:**
```
11
12
```
