package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/cobra"
)

func main() {
	var landlordName, tenantName, address string
	var rentAmount float64
	var fromMonth, toMonth string
	var inputJSON string

	var rootCmd = &cobra.Command{
		Use:   "rent-receipt",
		Short: "Generate rent receipts for HRA claims",
		Run: func(cmd *cobra.Command, args []string) {
			if inputJSON != "" {
				parseJSONInput(inputJSON, &landlordName, &tenantName, &address, &rentAmount, &fromMonth, &toMonth)
			}
			generateReceiptsPDF(landlordName, tenantName, address, rentAmount, fromMonth, toMonth)
		},
	}

	rootCmd.Flags().StringVarP(&landlordName, "landlord", "l", "", "Landlord's name (optional if JSON is provided)")
	rootCmd.Flags().StringVarP(&tenantName, "tenant", "t", "", "Tenant's name (optional if JSON is provided)")
	rootCmd.Flags().StringVarP(&address, "address", "a", "", "Property address (optional if JSON is provided)")
	rootCmd.Flags().Float64VarP(&rentAmount, "rent", "r", 0, "Monthly rent amount (optional if JSON is provided)")
	rootCmd.Flags().StringVarP(&fromMonth, "from", "f", "", "Start month (e.g., January 2024) (optional if JSON is provided)")
	rootCmd.Flags().StringVarP(&toMonth, "to", "o", "", "End month (e.g., December 2024) (optional if JSON is provided)")
	rootCmd.Flags().StringVarP(&inputJSON, "json", "j", "", "Input data as JSON file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseJSONInput(filePath string, landlordName, tenantName, address *string, rentAmount *float64, fromMonth, toMonth *string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	data := struct {
		LandlordName string  `json:"landlord"`
		TenantName   string  `json:"tenant"`
		Address      string  `json:"address"`
		RentAmount   float64 `json:"rent"`
		FromMonth    string  `json:"from"`
		ToMonth      string  `json:"to"`
	}{}

	if err := decoder.Decode(&data); err != nil {
		fmt.Printf("Error parsing JSON file: %v\n", err)
		os.Exit(1)
	}

	*landlordName = data.LandlordName
	*tenantName = data.TenantName
	*address = data.Address
	*rentAmount = data.RentAmount
	*fromMonth = data.FromMonth
	*toMonth = data.ToMonth
}

func generateReceiptsPDF(landlordName, tenantName, address string, rentAmount float64, fromMonth, toMonth string) {
	start, err := time.Parse("January 2006", fromMonth)
	if err != nil {
		fmt.Printf("Invalid start month format: %v\n", err)
		os.Exit(1)
	}

	end, err := time.Parse("January 2006", toMonth)
	if err != nil {
		fmt.Printf("Invalid end month format: %v\n", err)
		os.Exit(1)
	}

	if start.After(end) {
		fmt.Println("Start month cannot be after end month")
		os.Exit(1)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	current := start

	for !current.After(end) {
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Rent Receipt")

		pdf.SetFont("Arial", "", 12)
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Landlord: %s", landlordName))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Tenant: %s", tenantName))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Address: %s", address))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Rent Amount: %.2f", rentAmount))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Month: %s", current.Format("January 2006")))
		pdf.Ln(20)
		pdf.Cell(0, 10, fmt.Sprintf("Date: %s", todayDate()))
		pdf.Ln(10)
		pdf.Cell(0, 10, "Signature: ________________")

		current = current.AddDate(0, 1, 0)
	}

	err = pdf.OutputFileAndClose("rent_receipts.pdf")
	if err != nil {
		fmt.Printf("Error generating PDF: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Rent receipts generated as rent_receipts.pdf")
}

func todayDate() string {
	return time.Now().Format("02/01/2006")
}
