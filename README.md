# api-library

This is an API that is used to manage libraries such as borrowing books, returning books and others

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

**!!! HANYA ADMIN YANG DAPAT MENGGUNAKAN FITUR INI !!!**

Digunakan untuk mendapatkan data pengguna dengan menggunakan parameter **page** dan **limit** untuk mengatur halaman dan batas hasil.

#### Endpoint

```http
GET http://localhost:3000/api/user?page=1&limit=20
Access-Token: <YourAccessToken>
```

#### Query Parameter

- **page** (number,optional): Halaman ke berapa
- **limit** (number,optional): Jumlah data perhalaman

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
			"id": "8eb09df1-56da-4927-82df-8d7a4681266b",
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

**!!! HANYA ADMIN YANG DAPAT MENGGUNAKAN FITUR INI !!!**

Digunakan untuk mendapatkan data pengguna berdasarkan **ID** pengguna yang ditentukan.

#### Endpoint

```http
GET http://localhost:3000/api/user/8eb09df1-56da-4927-82df-8d7a4681266b
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **id** (uuid, required): user id

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
		"id": "8eb09df1-56da-4927-82df-8d7a4681266b",
		"username": "admin",
		"email": "admin@gmail.com",
		"address": "",
		"phone_number": "",
		"image_url": "",
		"created_at": "2023-09-18T20:21:06.783696+07:00",
		"updated_at": "2023-09-18T22:01:38.383622+07:00",
		"list_of_books": []
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
PUT http://localhost:3000/api/user/8eb09df1-56da-4927-82df-8d7a4681266b
Access-Token: <YourAccessToken>
```

#### Url Parameter

- **id** (uuid, required): user id

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
