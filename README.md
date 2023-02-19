## AXION

Axion is a simple auction application built with Go and React. It is a simple application that allows users to create auctions for products. Users can bid on auctions and the highest bidder wins the auction. The application also allows users to create products and auctions for those products.

## Getting Started

1. Clone the repository: `git clone https://github.com/codeyzx/axion.git`
2. Navigate to the project directory: `cd axion`
3. Install the dependencies: `go get`
4. Create mysql database: `axion`
5. Run the application: `go run main.go`

> notes: database will be migrated automatically

## API Documentation

| METHOD   | ROUTE                         | FUNCTIONALITY                                     | ACCESS           |
| -------- | ----------------------------- | ------------------------------------------------- | ---------------- |
| _GET_    | `/`                           | _Redirects to API documentation_                  | _All users_      |
| _GET_    | `/docs/*`                     | _Serves the API documentation page_               | _All users_      |
| _GET_    | `/public/*`                   | _Serves static files (images, CSS, etc.)_         | _All users_      |
| _POST_   | `/login`                      | _Logs in a user_                                  | _All users_      |
| _POST_   | `/register`                   | _Registers a new user_                            | _All users_      |
| _GET_    | `/users`                      | _Gets a list of all users_                        | _Admin_          |
| _GET_    | `/users/:id`                  | _Gets information on a single user_               | _All users_      |
| _PUT_    | `/users/:id`                  | _Updates a user's profile_                        | _Owner/Admin_    |
| _PUT_    | `/users/:id/update-email`     | _Updates a user's email address_                  | _Owner/Admin_    |
| _PUT_    | `/users/:id/update-role`      | _Updates a user's role (admin/operator/user)_     | _Admin_          |
| _DELETE_ | `/users/:id`                  | _Delete a user_                                   | _Admin_          |
| _GET_    | `/products`                   | _Gets a list of all products_                     | _All users_      |
| _GET_    | `/products/:id`               | _Gets information on a single product_            | _All users_      |
| _POST_   | `/products`                   | _Creates a new product_                           | _Operator_       |
| _PUT_    | `/products`                   | _Update a product_                                | _Owner/Operator_ |
| _GET_    | `/auctions`                   | _Gets a list of all auctions_                     | _All users_      |
| _GET_    | `/auctions/:id`               | _Gets information on a single auction_            | _All users_      |
| _POST_   | `/auctions`                   | _Creates a new auction_                           | _User/Operator_  |
| _PUT_    | `/auctions/:id`               | _Update an auction_                               | _User/Operator_  |
| _DELETE_ | `/auctions/:id`               | _Delete an auction_                               | _User/Operator_  |
| _GET_    | `/auctions-histories`         | _Gets a list of all auction histories_            | _Admin_          |
| _GET_    | `/auction-histories/:id`      | _Gets information on a single auction history_    | _All users_      |
| _GET_    | `/auction-histories/user/:id` | _Gets a list of all auction histories for a user_ | _Owner/Admin_    |
| _POST_   | `/auction-histories/`         | _Create a new auction history_                    | _User/Operator_  |
| _PUT_    | `/auction-histories/:id`      | _Update an auction history_                       | _Admin_          |
| _DELETE_ | `/auction-histories/:id`      | _Delete an auction history_                       | _Admin_          |
| _GET_    | `/history`                    | _Gets a list of all history log_                  | _Admin_          |
| _GET_    | `/history/:id`                | _Gets information on a single history log_        | _Admin_          |
| _POST_   | `/history`                    | _Creates a new history log_                       | _Admin_          |
| _PUT_    | `/history/:id`                | _Update an history log_                           | _Admin_          |
| _DELETE_ | `/history/:id`                | _Delete an history log_                           | _Admin_          |

## About the Project

Project ini bertujuan untuk memenuhi `UJI KOMPETENSI KEAHLIAN REKAYASA PERANGKAT LUNAK TAHUN PELAJARAN 2022/2023` dengan judul tugas `Sistem Lelang Online` pada `paket 4`.
