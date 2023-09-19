# api-library

## Description

Ini adalah sebuah API yang digunakan untuk mengelola berbagai aktivitas terkait perpustakaan, termasuk peminjaman dan pengembalian buku, dan lain sebagainya.

## Feature

### Register User

digunakan untuk mendaftarkan pengguna baru.

#### Endpoint

```http
POST http://localhost:3000/api/auth/register
```

#### Request Body

- **username** (string, required): Nama pengguna.
- **email** (string, required): Alamat email pengguna.
- **role** (string, optional): Peran pengguna (dapat dikosongkan).
- **password** (string, required): Kata sandi pengguna.
- **address** (string, optional): Alamat pengguna (dapat dikosongkan).
- **phone_number** (string, optional): Nomor telepon pengguna (dapat dikosongkan).
- **image_url** (string, optional): URL gambar profil pengguna (dapat dikosongkan).

Contoh Request Body:

```json
{
	"username": "admin",
	"email": "admin@gmail.com",
	"role": "",
	"password": "admin123",
	"address": "",
	"phone_number": "",
	"image_url": ""
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"username": "admin",
		"email": "admin@gmail.com",
		"role": "user",
		"password": "$2a$10$jWefb0SQvC7Xo3tibOyEueCyMwXLckbuXYUpG1H9VUehdci0wKaoG",
		"address": "",
		"phone_number": "",
		"image_url": ""
	}
}
```

#### Catatan

- Peran (**role**) pengguna dalam contoh di atas diatur sebagai "user" atau "admin". Anda dapat menggantinya sesuai dengan kebutuhan aplikasi Anda.
- Kata sandi (**password**) yang digunakan dalam contoh di atas telah di-hash dengan algoritma bcrypt.
- Beberapa bidang seperti **role**, **address**, **phone_number**, dan **image_url** dapat dikosongkan sesuai dengan kebutuhan aplikasi Anda.

##

### Login User

Digunakan untuk mengotentikasi pengguna dan menghasilkan AccessToken.

#### Endpoint

```http
POST http://localhost:3000/api/auth/login
```

#### Request Body

- **email** (string, required): Alamat email pengguna.
- **password** (string, required): Kata sandi pengguna.

Contoh Request Body:

```json
{
	"email": "admin@gmail.com",
	"password": "admin123"
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json
- **Set-Cookie**:AccessToken=**YourRefreshToken**

```json
{
	"code": 200,
	"status": "Your Access Token",
	"data": "<YourAccessToken>"
}
```

#### Catatan

- **Token** yang diatur dalam cookie dan **Token** yang diberikan melalui response dalam bentuk JSON pada body memiliki nilai yang berbeda.

##

### Get Token

Digunakan untuk mendapatkan AccessToken yang telah diatur dalam cookie.

#### Endpoint

```http
GET http://localhost:3000/api/auth/token
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

```json
{
	"code": 200,
	"status": "Your Access Token",
	"data": "<YourAccessToken>"
}
```

#### Catatan

Hanya dapat digunakan setelah pengguna berhasil login. Jika pengguna belum melakukan login, permintaan akan menghasilkan respons berikut:

- **HTTP Status**: 401 Unauthorized
- **Content-Type**: application/json

```json
{
	"message": "Unauthorized"
}
```

##

### Logout User

Digunakan untuk keluar dan menghapus AccessToken dari cookie.

#### Endpoint

```http
DELETE http://localhost:3000/api/auth/logout
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

```json
{
	"code": 200,
	"status": "You Have Logged Out",
	"data": null
}
```

##

### Get All Users

> [!IMPORTANT]
> Hanya role admin yang dapat menggunakan fitur ini !!!

Digunakan untuk mendapatkan data pengguna dengan menggunakan parameter **page** dan **limit** untuk mengatur halaman dan batas hasil.

#### Endpoint

```http
GET http://localhost:3000/api/user?page=1&limit=20
Access-Token: <YourAccessToken>
```

#### Query Parameters

- **page** (integer, optional): Nomor halaman yang diinginkan (default: 1).
- **limit** (integer, optional): Jumlah maksimal entri per halaman (default: 5).

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"current_page": 1,
	"total_page": 1,
	"data": [
		{
			"id": "<UserId>",
			"username": "admin",
			"email": "admin@gmail.com",
			"address": "",
			"phone_number": "",
			"image_url": ""
		}
	]
}
```

##

### Get User By ID

> [!IMPORTANT]
> Hanya role admin yang dapat menggunakan fitur ini !!!

Digunakan untuk mendapatkan data pengguna berdasarkan **ID** pengguna yang ditentukan.

#### Endpoint

```http
GET http://localhost:3000/api/user/<UserId>
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **UserId** (uuid, required): user id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"id": "<UserId>",
		"username": "admin",
		"email": "admin@gmail.com",
		"address": "",
		"phone_number": "",
		"image_url": "",
		"created_at": "2023-09-18T20:21:06.783696+07:00",
		"updated_at": "2023-09-18T22:01:38.383622+07:00",
		"list_of_books": [
			{
				"id": "d4f79719-e5aa-4f9b-b9eb-67ee328aa6bb",
				"title": "Praktek Pemrograman",
				"author": "Admin Keren",
				"publication_year": 2020,
				"image_url": "",
				"borrowing_status": {
					"borrowing_date": "2023-09-19T00:52:10.019807+07:00",
					"return_date": "2023-09-26T00:52:10.019807+07:00",
					"status": "borrowed"
				}
			}
		]
	}
}
```

#### Catatan

- **list_of_books** akan berisi daftar buku yang dipinjam oleh User tersebut.

##

### Update User By ID

Digunakan untuk memperbarui data pengguna berdasarkan **ID** pengguna yang ditentukan.

#### Endpoint

```http
PUT http://localhost:3000/api/user/<UserId>
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **UserId** (uuid, required): user id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Request Body

- **username** (string, optional): Nama pengguna baru.
- **email** (string, optional): Alamat email pengguna baru.
- **password** (string, optional): Kata sandi pengguna baru.
- **address** (string, optional): Alamat pengguna baru.
- **phone_number** (string, optional): Nomor telepon pengguna baru.
- **image_url** (string, optional): URL gambar profil pengguna baru.

Contoh Request Body:

```json
{
	"username": "Admin Keren",
	"email": "",
	"password": "",
	"address": "",
	"phone_number": "",
	"image_url": ""
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"username": "Admin Keren",
		"email": "",
		"password": "",
		"address": "",
		"phone_number": "",
		"image_url": ""
	}
}
```

##

### Delete User By ID

Digunakan untuk menghapus data pengguna berdasarkan **ID** pengguna yang ditentukan.

#### Endpoint

```http
DELETE http://localhost:3000/api/user/<UserId>
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **UserId** (uuid, required): user id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": "Success Deleted User With Id '<UserId>'"
}
```

##

### Create New Book

> [!IMPORTANT]
> Hanya role admin yang dapat menggunakan fitur ini !!!

Digunakan untuk menambahkan buku baru ke dalam sistem.

#### Endpoint

```http
POST http://localhost:3000/api/book
Access-Token: <YourAccessToken>
```

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Request Body

- **title** (string, required): Judul buku baru.
- **author** (string, required): Penulis buku baru.
- **publication_year** (integer, required): Tahun publikasi buku baru.
- **image_url** (string, optional): URL gambar sampul buku.
- **total** (integer, required): Jumlah total salinan buku yang tersedia.

Contoh Request Body:

```json
{
	"title": "Dasar Pemrograman",
	"author": "Admin Keren",
	"publication_year": 2020,
	"image_url": "",
	"total": 50
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"title": "Dasar Pemrograman",
		"author": "Admin Keren",
		"publication_year": 2020,
		"image_url": "",
		"total": 50
	}
}
```

##

### Get All Books

Digunakan untuk mendapatkan data buku dengan menggunakan parameter **page** dan **limit** untuk mengatur halaman dan batas hasil.

#### Endpoint

```http
GET http://localhost:3000/api/book?page=1&limit=3
Access-Token: <YourAccessToken>
```

#### Query Parameters

- **page** (integer, optional): Nomor halaman yang diinginkan (default: 1).
- **limit** (integer, optional): Jumlah maksimal entri per halaman (default: 5).

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"current_page": 1,
	"total_page": 1,
	"data": [
		{
			"id": "<BookId>",
			"title": "Dasar Pemrograman",
			"author": "Admin Keren",
			"publication_year": 2020,
			"image_url": "",
			"total": 50
		}
	]
}
```

##

### Update Book By Id

> [!IMPORTANT]
> Hanya role admin yang dapat menggunakan fitur ini !!!

Digunakan untuk memperbarui data buku berdasarkan **ID** buku yang ditentukan.

#### Endpoint

```http
PUT http://localhost:3000/api/book/<BookId>
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **BookId** (uuid, required): book id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Request Body

- **title** (string, optional): Judul buku yang baru.
- **author** (string, optional): Penulis buku yang baru.
- **publication_year** (integer, optional): Tahun publikasi buku yang baru.
- **image_url** (string, optional): URL gambar sampul buku yang baru.
- **total** (integer, optional): Jumlah total salinan buku yang baru.

Contoh Request Body:

```json
{
	"title": "Praktek Pemrograman",
	"author": "",
	"publication_year": 0,
	"image_url": "",
	"total": 0
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"title": "Praktek Pemrograman",
		"author": "",
		"publicaion_year": 0,
		"image_url": "",
		"total": 0
	}
}
```

##

### Delete Book By Id

> [!IMPORTANT]
> Hanya role admin yang dapat menggunakan fitur ini !!!

Digunakan untuk menghapus buku berdasarkan **ID** buku yang ditentukan.

#### Endpoint

```http
DELETE  http://localhost:3000/api/book/<BookId>
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **BookId** (uuid, required): book id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": "Succes Delete Book With Id '<BookId>'"
}
```

##

### Get Book By ID

Digunakan untuk mendapatkan detail buku berdasarkan **ID** buku yang ditentukan.

#### Endpoint

```http
GET  http://localhost:3000/api/book/<BookId>
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **BookId** (uuid, required): book id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"id": "<BookId>",
		"title": "Praktek Pemrograman",
		"author": "Admin Keren",
		"publicaion_year": 2020,
		"image_url": "",
		"total": 50,
		"created_at": "2023-09-18T23:38:04.845859+07:00",
		"updated_at": "2023-09-18T23:57:05.956312+07:00",
		"list_of_users": [
			{
				"id": "8eb09df1-56da-4927-82df-8d7a4681266b",
				"username": "Admin Keren",
				"email": "admin@gmail.com",
				"address": "",
				"phone_number": "",
				"image_url": "",
				"borrowing_status": {
					"borrowing_date": "2023-09-19T00:52:10.019807+07:00",
					"return_date": "2023-09-26T00:52:10.019807+07:00",
					"status": "borrowed"
				}
			}
		]
	}
}
```

#### Catatan

- **list_of_users** adalah data seluruh users yang meminjam buku tersebut.

##

### Borrowing Book By ID

Digunakan untuk meminjam buku berdasarkan **ID** buku yang ditentukan.

### Endpoint

```http
GET  http://localhost:3000/api/book/borrow/<BookId>
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **BookId** (uuid, required): book id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": "Success To Borrow A Book with ID '<BookId>'"
}
```

##

### Return Book By ID

Digunakan untuk mengembalikan buku berdasarkan **ID** buku yang ditentukan.

#### Endpoint

```http
GET  http://localhost:3000/api/book/return/<BookId>
Access-Token: <YourAccessToken>

```

#### Url Parameter

- **BookId** (uuid, required): book id

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"data": "Success To Return A Book with ID '<BookId>'"
}
```

##

### Get All Expired Books

> [!IMPORTANT]
> Hanya role admin yang dapat menggunakan fitur ini !!!

Digunakan untuk mendapatkan daftar buku yang telah kedaluwarsa.

#### Endpoint

```http
GET  http://localhost:3000/api/book/expired?page=1&limit=5
Access-Token: <YourAccessToken>
```

#### Query Parameters

- **page** (integer, optional): Nomor halaman yang diinginkan (default: 1).
- **limit** (integer, optional): Jumlah maksimal entri per halaman (default: 5).

#### Request Headers

- **Access-Token** (string, required): AccessToken yang diperoleh dari proses otentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json

Contoh Response Body:

```json
{
	"code": 200,
	"status": "OK",
	"current_page": 1,
	"total_page": 1,
	"data": [
		{
			"user": {
				"id": "<UserId>",
				"username": "Admin Keren",
				"email": "admin@gmail.com",
				"address": "",
				"phone_number": "",
				"image_url": ""
			},
			"book": {
				"id": "<BookId>",
				"title": "Praktek Pemrograman",
				"author": "Admin Keren",
				"publication_year": 2020,
				"image_url": "",
				"total": 49
			},
			"borrowing_date": "2023-09-19T00:52:10.019807+07:00",
			"return_date": "2023-09-19T00:56:59.627946+07:00",
			"status": "expired"
		}
	]
}
```

## MIT License

MIT License

Copyright (c) 2023 NURSANDY IHKSAN

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
