# CMS WardFunding Backend

### Technology
- Go v1.16.5

### Modules
- [Go Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)
- [Go JWT](https://github.com/dgrijalva/jwt-go)
- [Go Multitemplate](https://github.com/gin-contrib/multitemplate)
- [Go Accounting](https://github.com/leekchan/accounting)
- [Go Cors](https://github.com/gin-contrib/cors)
- [Go Session](https://github.com/gin-contrib/sessions)

### Payment API Gateway
- [Midtrans](https://midtrans.com/)

### Dasboard Admin
- [BootAdmin](https://web.archive.org/web/20210301183117/https://bootadmin.net/demo/docs)

### Install Project
```
git clone https://github.com/abiwardanii/wardfunding-backend.git
cd wardfunding-backend
```

### Run Project
```
go run main.go
```

### APIs User 
| Method | Endpoint | Auth Middleware | Description |
| --- | --- | --- | --- | 
| POST | /users| :x: | Register User |
| POST | /sessions | :x: | Login |
| POST | /email_checkers | :x: | Check Email Avaibility |
| POST | /avatars | :white_check_mark: | Upload Avatar |
| GET | /users/fetch | :white_check_mark: | Fetch User |

### APIs Campaigns 
| Method | Endpoint | Auth Middleware | Description |
| --- | --- | --- | --- |
| GET | /campaigns| ::x:: | List of Campaign |
| GET | /campaigns/:id | :x: | Get Detail Campaign |
| POST | /campaigns | :x: | Create Campaign |
| PUT | /campaigns/:id | :white_check_mark: | Update Campaign |
| POST | /campaign-images | :white_check_mark: | Upload Image  |

### APIs Transactions 
| Method | Endpoint | Auth Middleware | Description |
| --- | --- | --- | --- |
| GET | /campaigns/:id/transactions| :white_check_mark: | Campaign Transaction |
| GET | /transactions | :white_check_mark: | Get User Transaction |
| POST | /transactions | :white_check_mark: |  Create Transaction |
| POST | /transactions/notification | :x: | Get Notification |

### CMS APIs User 
| Method | Endpoint | Auth Admin Middleware | Description |
| --- | --- | --- | --- | 
| GET | /users | :white_check_mark: | List of user|
| GET | /users/new | :x: | Form Create User |
| POST | /users | :x: |  Create New User |
| GET | /users/edit/:id | :x: |  Edit a User  |
| POST | /users/update/:id | :white_check_mark: | Update User |
| GET | /users/avatar/:id | :white_check_mark: | New Avatar  |
| POST | /users/avatar/:id | :white_check_mark: | Create Avatar  |

### CMS APIs Campaign
| Method | Endpoint | Auth Admin Middleware | Description |
| --- | --- | --- | --- | 
| GET | /campaigns | :white_check_mark: | List of Campaign |
| GET | /campaigns/new | :white_check_mark: | Form Create Campaign  |
| POST | /campaigns | :white_check_mark: | Create campaign |
| GET | /campaigns/image/:id | :white_check_mark: | New Image |
| POST | /campaigns/image/:id | :white_check_mark: | Create Image |
| GET | /campaigns/edit/:id | :white_check_mark: | Form Edit|
| POST | /campaigns/update/:id | :white_check_mark: | Update Campaign |
| GET | /campaigns/show/:id | :white_check_mark: | Detail Campaign |

### CMS APIs Transaction 
| Method | Endpoint | Auth Admin Middleware | Description |
| --- | --- | --- | --- | 
| GET | /transactions | :white_check_mark: | List of Transaction |

### CMS APIs Session 
| Method | Endpoint | Description |
| --- | --- | --- |
| GET | /login | List of Transaction |
| POST | /session | List of Transaction |
| GET | /logout | List of Transaction |
