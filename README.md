# Rent Receipt Generator

A simple and efficient CLI tool to generate rent receipts for HRA claims or proof of investment. This tool supports creating receipts for a range of months and exports them as a PDF file, ensuring your sensitive information remains secure.

## Features

- **Generate Multiple Receipts**: Specify a date range to create receipts for each month.
- **PDF Export**: Automatically generates a PDF file containing all the receipts.
- **JSON Input Support**: Provide all required data through a JSON file.
- **User-Friendly CLI**: Easy-to-use commands and flags for customization.

---

## Requirements

- **Go**: Version 1.18 or higher.

---

## Installation

1. Clone this repository:

   ```bash
   git clone <repository-url>
   cd rent-receipt-generator
   ```

2. Build the application:

   ```bash
   go build -o rent-receipt
   ```

3. Run the tool:

   ```bash
   ./rent-receipt
   ```

---

## Usage

### Flags

- `-l, --landlord` : Landlord's name.
- `-t, --tenant` : Tenant's name.
- `-a, --address` : Property address.
- `-r, --rent` : Monthly rent amount.
- `-f, --from` : Start month (e.g., "January 2024").
- `-o, --to` : End month (e.g., "December 2024").
- `-j, --json` : Path to a JSON file containing input data.

### Examples

#### Using Flags

```bash
./rent-receipt -l "John Doe" -t "Jane Smith" -a "123 Elm Street, City" -r 15000 -f "January 2024" -o "March 2024"
```

#### Using JSON Input

1. Create a JSON file:

   ```json
   {
     "landlord": "John Doe",
     "tenant": "Jane Smith",
     "address": "123 Elm Street, City",
     "rent": 15000,
     "from": "January 2024",
     "to": "March 2024"
   }
   ```

2. Run the command:

   ```bash
   ./rent-receipt -j input.json
   ```

---

## Output

The tool generates a PDF file named `rent_receipts.pdf` in the current directory. Each receipt contains:

- Landlord's name
- Tenant's name
- Property address
- Rent amount
- Month of rent
- Date and signature placeholder

---

## License

This project is licensed under the MIT License.

---

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests to improve this tool.

---

## Author

Developed by [Your Name].

