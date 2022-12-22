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
 :white_check_mark: Login as staff or admin
 :white_check_mark: Different menu option for staff or admin
 :white_check_mark: Register a new staff account
 :white_check_mark: Insert a new product
 :white_check_mark: Insert a new customer
 :white_check_mark: Insert a new transaction
 :white_check_mark: Show and delete a product
 :white_check_mark: Show and delete a customer
 :white_check_mark: Show and delete a staff
 :white_check_mark: Show history transaction and delete transaction
 :white_check_mark: Update product information and stock
 :white_check_mark: Update customer information
 :white_check_mark: Update staff account information

## Requirements
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
