# User Service

## Use Cases

`User`:

- `RegisterNewUser`: create a new user
- `UpdateUser`: update a user

## Entites

`Booking`:
| Field     | Type       | Description |
| --------- | ---------- | ----------- |
| id        | uuid       |             |
| createdAt | timestampz |             |
| updatedAt | timestampz |             |
| deletedAt | timestampz |             |
| firstName | string     |             |
| lastName  | string     |             |

> Add address information to geolocate the nearby store
