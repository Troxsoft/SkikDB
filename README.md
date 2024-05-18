# Key/Value Memory Database 
## Characteristics
- Easy 
- Simple
- All on `only one` executable
- Basic `autocomplete` in *CLI*

## Data types
- ### String
- ### Number(with decimal or without decimal)
- ### List
- ### JSON
## Examples:
- Set key:
    ```sql

    set "INeedCash" true
    ```
    ![set example](things/set.png)
- Get key:
    ```sql
    get "INeedCash"
    ```
    ![get example](things/get.png)
- Get all keys/values:
    ```sql
    get *
    ```
- Delete key:
    ```sql
    delete "INeedCash"
    ```
- Get type of key/value:
    ```sql
    type "INeedCash"
    ```
- Exists key/value:
    ```sql
    exists "INeedCash"
    ```
# Sixtax Error:
![sintax error](things/sintaxError.png)
# Features 😫 `(Expected to be implemented in the future)`:
- List `(add,delete,sort,is repeat,and more)`
- Limit
- Conditions `(likes sql)` 
- Data handling `(strings,numbers)`
- **DB Server**
- JSON handling
- Regex
- Exact type
- Update message error
- Add native **API** in golang
- Add **API** using DB server for others programing languages `(nodejs,c,c++,rust,python)`
#### This proyect use [c-bata/go-prompt](!https://github.com/c-bata/go-prompt) for autocomplete in the cli,[fatih/color](https://github.com/fatih/color) for terminal colors
#### [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)