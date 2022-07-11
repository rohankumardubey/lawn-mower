# Booking Service

## Use Cases

`Booking`:

- `CreateBooking`: create a new booking
- `CancelBooking`: cancel a booking

## Entites

`Booking`:
| Field          | Type                | Description |
| -------------- | ------------------- | ----------- |
| id             | uuid                |             |
| createdAt      | timestampz          |             |
| updatedAt      | timestampz          |             |
| deletedAt      | timestampz          |             |
| ref            | string              |             |
| bookedMinAt    | timestampz          |             |
| bookedMaxAt    | timestampz          |             |
| storeInventory | StoreInventory.[id] |             |
| user           | User.[id]           |             |
