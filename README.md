# payment

## Installation

Install the dependencies and devDependencies and start the server.

```sh
go mod tidy
```

## Migration 
```sh
make migrate-new // create new script for migrate 
make migration-up //up script 
```

## Run 
```sh
make run
```

## List API
```sh
/v1/payment [GET]
```
| Query Param | Value |
| ------ | ------ |
| account_number | string |
| bank_id | string |

```sh
sample response

{
    "code": 200,
    "message": "success",
    "data": {
        "beneficiary_name": "Joh Khannedy",
        "beneficiary_account_number": "1234567890",
        "beneficiary_banks": {
            "id": "12538030-4e59-11ed-8b9d-0242ac120005",
            "beneficiary_bank_name": "Bank Republik Indonesia",
            "created_at": "2024-04-06T10:08:01.237887Z",
            "updated_at": "2024-04-06T10:08:01.237887Z"
        },
        "beneficiary_amount": 990000,
        "status": "active",
        "created_at": "2024-04-06T07:37:44.438613Z",
        "updated_at": "2024-04-06T07:37:44.438613Z"
    }
}
```

```sh
/v1/payment/transfer-disbursement [POST]
```
| Body Request | Value | Information |
| ------ | ------ | ------ |
| sender_account_number | string | mandatory |
| recipient_account_number | string | mandatory |
| beneficiary_bank_id_sender | string | mandatory |
| beneficiary_bank_id_recipient | string | mandatory |
| amount | number | mandatory |

```sh
sample response

{
    "code": 200,
    "message": "success",
    "data": "47df451e-0dd0-41a2-81dc-f88f3734df4f" // transfer_id
}
```

```sh
/v1/payment/callback-transfer-disbursement [PUT]
```
| Body Request | Value | Information |
| ------ | ------ | ------ |
| sender_account_number | string | mandatory |
| recipient_account_number | string | mandatory |
| beneficiary_bank_id_sender | string | mandatory |
| beneficiary_bank_id_recipient | string | mandatory |
| transfer_id | string | mandatory |

```sh
sample response

{
    "code": 200,
    "message": "success",
    "data": null
}
```