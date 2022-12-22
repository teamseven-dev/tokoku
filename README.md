# Tokoku App

Tokoku is a CLI based point-of-sale app built with Golang and MySQL.

## Table Of Content

- [Features of the App](#features-of-the-app)
- [Programming Languages](#programming-languages)
- [Entity Relationship Diagram (ERD)](#entity-relationship-diagram-erd)
- [Folder Structure Pattern](#folder-structure-pattern)
- [How to Run](#how-to-run)
- [Credit](#credit)

## Features of the App
- [x] Login as staff or admin
- [x] Different menu option for staff or admin
- [x] Register a new staff account
- [x] Insert a new product
- [x] Insert a new customer
- [x] Insert a new transaction
- [x] Show and delete a product
- [x] Show and delete a customer
- [x] Show and delete a staff
- [x] Show history transaction and delete transaction
- [x] Update product information and stock
- [x] Update customer information
- [x] Update staff account information

## Programming Languages
- Go v1.19
- MySQL v8.x

## Entity Relationship Diagram (ERD)
![run](./ERD.png)

## Folder Structure Pattern
```
├── config
│   └── config.go
└── customer
│   └── customer.go
└── product
│   └── product.go
└── staff
│   └── staff.go
└── transaction
│   └── transaction.go
├── .gitignore
├── ERD.png
├── go.mod
├── go.sum
├── LICENSE
├── local.env.example
├── main.go
├── README.md
└── tokoku-script.sql
```

## How to Run

- Clone it

```
$ git clone [https://github.com/teamseven-dev/tokoku.git]
```

- Go to directory

```
$ cd tokoku
```

- Run the project

```
$ go run .
```

## Credit
[Indra Darmawan](https://github.com/e1more)

[Kharisma Januar Muhammad JN](https://github.com/kharismajanuar)

[Muhammad Habibullah](https://github.com/hebobibun)



<p align="right" style="padding: 5px; border-radius: 100%; background-color: red; font-size: 2rem;">
  <b><a href="#tokoku-app">BACK TO TOP</a></b>
</p>
