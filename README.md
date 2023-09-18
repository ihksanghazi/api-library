# api-library

This is an API that is used to manage libraries such as borrowing books, returning books and others

## Feature

### Register User

digunakan untuk mendaftarkan pengguna baru.

#### Endpoint

```bash
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
  Contoh Response Body:

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

- Peran (**role**) pengguna dalam contoh di atas diatur sebagai "user" atau "password". Anda dapat menggantinya sesuai dengan kebutuhan aplikasi Anda.
- Kata sandi (**password**) yang digunakan dalam contoh di atas telah di-hash dengan algoritma bcrypt.
- Beberapa bidang seperti **role**, **address**, **phone_number**, dan **image_url** dapat dikosongkan sesuai dengan kebutuhan aplikasi Anda.

### Login User

Digunakan untuk mengotentikasi pengguna dan menghasilkan AccessToken.

#### Endpoint

```bash
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
- **Set-Cookie**: AccessToken=<YourToken>
  Contoh Response Body:

  ```json
  {
  	"code": 200,
  	"status": "Your Access Token",
  	"data": "<YourToken>"
  }
  ```

#### Catatan

- **Token** yang diatur dalam cookie dan **Token** yang diberikan melalui respons dalam bentuk JSON pada body memiliki nilai yang berbeda.
