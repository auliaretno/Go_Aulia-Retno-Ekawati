**API DOCUMENTATION** 

**REGISTER**

Method: POST

Endpoint: users/register

Request (Body):

{

`    `"name": "aulia2",

`    `"email": "aulia2@gmail.com",

`    `"role": "User",

`    `"password": "aulia"

}

Response:

{

`    `"message": "success create new user",

`    `"user": {

`        `"ID": 3,

`        `"CreatedAt": "2023-05-26T17:24:57.357+07:00",

`        `"UpdatedAt": "2023-05-26T17:24:57.357+07:00",

`        `"DeletedAt": null,

`        `"name": "aulia2",

`        `"email": "aulia2@gmail.com",

`        `"password": "aulia"

`    `}

}

**LOGIN**

Method: POST

Endpoint: users/login

Request:

{

`    `"email": "aulia@gmail.com",

`    `"password": "aulia"

}

response:

{

`    `"message": "success login user",

`    `"user": {

`        `"id": 1,

`        `"name": "aulia",

`        `"email": "aulia@gmail.com",

`        `"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUxMDAzNjgsIm5hbWUiOiJhdWxpYSIsInVzZXJJRCI6MX0.kEb-zAUi6EyiGZ3s1Od0XmMk50NVzGMMFEiwTSHwlUY"

`    `}

}

**CREATE ITEM**

Method: POST

Endpoint: /items

Request:

{

`    `"category\_id": 3,

`    `"name": "toshiba",

`    `"description": "available",

`    `"stock": 7,

`    `"price": 3000000

}

Response:

{

`    `"item": {

`        `"ID": 12,

`        `"CreatedAt": "2023-05-26T17:31:41.75+07:00",

`        `"UpdatedAt": "2023-05-26T17:31:41.75+07:00",

`        `"DeletedAt": null,

`        `"category\_id": 3,

`        `"Category": {

`            `"ID": 3,

`            `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`            `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`            `"DeletedAt": null,

`            `"category\_name": "TV"

`        `},

`        `"name": "toshiba",

`        `"description": "available",

`        `"stock": 7,

`        `"price": 3000000

`    `},

`    `"message": "success create new item"

}

**GET ALL ITEM**

Method: GET

Endpoint: /items

Response:

{

`    `"items": [

`        `{

`            `"ID": 1,

`            `"CreatedAt": "2023-05-25T11:30:37.761+07:00",

`            `"UpdatedAt": "2023-05-25T11:30:37.761+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-05-23T12:21:43.307+07:00",

`                `"UpdatedAt": "2023-05-26T17:18:44.388+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "Mesin Cucii"

`            `},

`            `"name": "electrolux",

`            `"description": "available",

`            `"stock": 10,

`            `"price": 5000000

`        `},

`        `{

`            `"ID": 2,

`            `"CreatedAt": "2023-05-25T12:10:03.581+07:00",

`            `"UpdatedAt": "2023-05-25T12:10:03.581+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-05-23T12:21:43.307+07:00",

`                `"UpdatedAt": "2023-05-26T17:18:44.388+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "Mesin Cucii"

`            `},

`            `"name": "Toshiba",

`            `"description": "available",

`            `"stock": 8,

`            `"price": 3000000

`        `},

`        `{

`            `"ID": 3,

`            `"CreatedAt": "2023-05-25T12:12:45.018+07:00",

`            `"UpdatedAt": "2023-05-26T00:37:48.973+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-05-23T12:21:43.307+07:00",

`                `"UpdatedAt": "2023-05-26T17:18:44.388+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "Mesin Cucii"

`            `},

`            `"name": "Sharp",

`            `"description": "available",

`            `"stock": 7,

`            `"price": 3000000

`        `},

`        `{

`            `"ID": 5,

`            `"CreatedAt": "2023-05-26T16:55:39.82+07:00",

`            `"UpdatedAt": "2023-05-26T17:22:09.589+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 2,

`            `"Category": {

`                `"ID": 2,

`                `"CreatedAt": "2023-05-23T12:21:58.346+07:00",

`                `"UpdatedAt": "2023-05-23T12:21:58.346+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "Setrika"

`            `},

`            `"name": "maspion",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 3000000

`        `},

`        `{

`            `"ID": 6,

`            `"CreatedAt": "2023-05-26T16:55:57.247+07:00",

`            `"UpdatedAt": "2023-05-26T16:55:57.247+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 2,

`            `"Category": {

`                `"ID": 2,

`                `"CreatedAt": "2023-05-23T12:21:58.346+07:00",

`                `"UpdatedAt": "2023-05-23T12:21:58.346+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "Setrika"

`            `},

`            `"name": "sanken",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 200000

`        `},

`        `{

`            `"ID": 7,

`            `"CreatedAt": "2023-05-26T16:56:18.745+07:00",

`            `"UpdatedAt": "2023-05-26T16:56:18.745+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 2,

`            `"Category": {

`                `"ID": 2,

`                `"CreatedAt": "2023-05-23T12:21:58.346+07:00",

`                `"UpdatedAt": "2023-05-23T12:21:58.346+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "Setrika"

`            `},

`            `"name": "kirin",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 200000

`        `},

`        `{

`            `"ID": 8,

`            `"CreatedAt": "2023-05-26T16:57:04.755+07:00",

`            `"UpdatedAt": "2023-05-26T16:57:04.755+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "panasonic",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 2000000

`        `},

`        `{

`            `"ID": 9,

`            `"CreatedAt": "2023-05-26T16:57:29.271+07:00",

`            `"UpdatedAt": "2023-05-26T16:57:29.271+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "LG",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 3000000

`        `},

`        `{

`            `"ID": 10,

`            `"CreatedAt": "2023-05-26T16:58:11.316+07:00",

`            `"UpdatedAt": "2023-05-26T16:58:11.316+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "samsung",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 3000000

`        `},

`	`{

`            `"ID": 12,

`            `"CreatedAt": "2023-05-26T17:31:41.75+07:00",

`            `"UpdatedAt": "2023-05-26T17:31:41.75+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "toshiba",

`            `"description": "available",

`            `"stock": 7,

`            `"price": 3000000

`        `}

`    `],

`    `"messages": "success get all products"

}

**GET ITEM BY ID**

Method: GET

Endpoint: /items/:id

Response:

{

`    `"items": {

`        `"ID": 8,

`        `"CreatedAt": "2023-05-26T16:57:04.755+07:00",

`        `"UpdatedAt": "2023-05-26T16:57:04.755+07:00",

`        `"DeletedAt": null,

`        `"category\_id": 3,

`        `"Category": {

`            `"ID": 3,

`            `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`            `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`            `"DeletedAt": null,

`            `"category\_name": "TV"

`        `},

`        `"name": "panasonic",

`        `"description": "available",

`        `"stock": 6,

`        `"price": 2000000

`    `},

`    `"messages": "success get item"

}

**GET ITEM BY ID CATEGORY**

Method: GET

Endpoint: /items/category/:id

Response:

{

`    `"items": [

`        `{

`            `"ID": 8,

`            `"CreatedAt": "2023-05-26T16:57:04.755+07:00",

`            `"UpdatedAt": "2023-05-26T16:57:04.755+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "panasonic",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 2000000

`        `},

`        `{

`            `"ID": 9,

`            `"CreatedAt": "2023-05-26T16:57:29.271+07:00",

`            `"UpdatedAt": "2023-05-26T16:57:29.271+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "LG",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 3000000

`        `},

`        `{

`            `"ID": 10,

`            `"CreatedAt": "2023-05-26T16:58:11.316+07:00",

`            `"UpdatedAt": "2023-05-26T16:58:11.316+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "samsung",

`            `"description": "available",

`            `"stock": 6,

`            `"price": 3000000

`        `},

`        `{

`            `"ID": 12,

`            `"CreatedAt": "2023-05-26T17:31:41.75+07:00",

`            `"UpdatedAt": "2023-05-26T17:31:41.75+07:00",

`            `"DeletedAt": null,

`            `"category\_id": 3,

`            `"Category": {

`                `"ID": 3,

`                `"CreatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"UpdatedAt": "2023-05-23T12:22:11.417+07:00",

`                `"DeletedAt": null,

`                `"category\_name": "TV"

`            `},

`            `"name": "toshiba",

`            `"description": "available",

`            `"stock": 7,

`            `"price": 3000000

`        `}

`    `],

`    `"messages": "success get item"

}

**UPDATE ITEM**

Method: PUT

Endpoint: /items/:id

Request:

{

`    `"category\_id": 2,

`    `"name": "maspion",

`    `"description": "available",

`    `"stock": 9,

`    `"price": 3000000

}

Response:

{

`    `"message": "success update item data",

`    `"user": {

`        `"ID": 5,

`        `"CreatedAt": "2023-05-26T16:55:39.82+07:00",

`        `"UpdatedAt": "2023-05-26T17:37:49.344+07:00",

`        `"DeletedAt": null,

`        `"category\_id": 2,

`        `"Category": {

`            `"ID": 2,

`            `"CreatedAt": "2023-05-23T12:21:58.346+07:00",

`            `"UpdatedAt": "2023-05-23T12:21:58.346+07:00",

`            `"DeletedAt": null,

`            `"category\_name": "Setrika"

`        `},

`        `"name": "maspion",

`        `"description": "available",

`        `"stock": 9,

`        `"price": 3000000

`    `}

}

**DELETE ITEM**

Method: DELETE

Endpoint: /items/:id

Response:

{

`    `"id": 11,

`    `"messages": "success delete product"

}

