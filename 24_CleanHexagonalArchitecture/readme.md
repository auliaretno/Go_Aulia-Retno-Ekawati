# SUMMARY Clean and Hexagonal Architecture

- Aulia Retno Ekawati
- Universitas Sahid Surakarta

## Konsep Clean Architecture
Setiap komponen yang ada bersifat independen dan tidak bergantung pada library external yang spesifik. Seperti tidak tergantung pada spesifik framework atau tidak bergantung pada spesifik database yang dipakai. 

## Context dalam clean and hexagonal architecture
Suatu package yang membawa deadline, concellation signal, atau value lain pada request/permintaan API. Implementasian context kita dapat membuat root context dengan fungsi background. Fungsi background sendiri akan membalikan root context dimana kita bisa memakainya untuk operasi lain.  

## Layer pada Clean Architecture
- Use Case/Domain Layer : Mengandung file-file logika bisnis.
- Controller/Handler/Presentation Layer : Berisi handler atau controller yang menangani user interaction yang menyediakan data ke layar
- Drivers/Data Layer : Berisi file yang mengatur data aplikasi, mengambil data dan mengatur data cache.
