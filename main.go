package main

import (
	"fmt"
)

// Vehicle represents a vehicle in the transport plan
type Vehicle struct {
	VehicleType       string
	Manufacture       string
	RouteNumber       string
	RegistrationNumber string
	OwnerName         string
	OwnerPhoneNumber  string
	OwnerAddress      string
	FuelType          string
	SeatingCapacity   int
	InsuranceDetails  Insurance
	AdditionalFeatures AdditionalFeatures
}

// Insurance holds insurance details for a vehicle
type Insurance struct {
	CompanyName       string
	PolicyNumber      string
	Covered           string // e.g., Third-Party, Zero Depreciation
	AmountPaid        float64
	IssueDate         string
	ExpiryDate        string
}

// AdditionalFeatures represents additional features of a vehicle
type AdditionalFeatures struct {
	AC              bool
	Radio           bool
	Camera          bool
	FireExtinguisher bool
	FirstAidKit     bool
}

// Driver represents a driver in the transport plan
type Driver struct {
	Name            string
	MobileNumber    string
	LicenseNumber   string
	IssueDate       string
	ExpiryDate      string
	AadharNumber    string
	BloodGroup      string
	PoliceVerified  bool
	Address         string
}

// Route represents a route in the transport plan
type Route struct {
	RouteNumber     string
	StartLocation   string
	EndLocation     string
	PickupPoints    []PickupPoint
}

// PickupPoint represents a pickup/drop point in a route
type PickupPoint struct {
	Name      string
	PickupTime string
	DropTime   string
	StopNumber int
}

// TransportPlan represents the overall transport plan
type TransportPlan struct {
	Vehicles []Vehicle
	Drivers  []Driver
	Routes   []Route
}

// AddVehicle adds a vehicle to the transport plan
func (tp *TransportPlan) AddVehicle(v Vehicle) {
	tp.Vehicles = append(tp.Vehicles, v)
}

// AddDriver adds a driver to the transport plan
func (tp *TransportPlan) AddDriver(d Driver) {
	tp.Drivers = append(tp.Drivers, d)
}

// AddRoute adds a route to the transport plan
func (tp *TransportPlan) AddRoute(r Route) {
	tp.Routes = append(tp.Routes, r)
}

// DisplayTransportPlan displays the transport plan details
func (tp *TransportPlan) DisplayTransportPlan() {
	fmt.Println("Transport Plan:")
	fmt.Println("Vehicles:")
	for _, v := range tp.Vehicles {
		fmt.Printf("  Vehicle Type: %s, Owner: %s, Route: %s\n", v.VehicleType, v.OwnerName, v.RouteNumber)
	}

	fmt.Println("Drivers:")
	for _, d := range tp.Drivers {
		fmt.Printf("  Driver Name: %s, Mobile: %s, License: %s\n", d.Name, d.MobileNumber, d.LicenseNumber)
	}

	fmt.Println("Routes:")
	for _, r := range tp.Routes {
		fmt.Printf("  Route Number: %s, Start: %s, End: %s\n", r.RouteNumber, r.StartLocation, r.EndLocation)
		for _, pp := range r.PickupPoints {
			fmt.Printf("    Pickup Point: %s, Stop: %d, Pickup Time: %s, Drop Time: %s\n", pp.Name, pp.StopNumber, pp.PickupTime, pp.DropTime)
		}
	}
}

func main() {
	transportPlan := TransportPlan{}

	// Adding a vehicle
	vehicle := Vehicle{
		VehicleType:       "Bus",
		Manufacture:       "XYZ Motors",
		RouteNumber:       "R1",
		RegistrationNumber: "AB1234CD",
		OwnerName:         "John Doe",
		OwnerPhoneNumber:  "1234567890",
		OwnerAddress:      "123 Main St",
		FuelType:          "Diesel",
		SeatingCapacity:   50,
		InsuranceDetails: Insurance{
			CompanyName:  "Insurance Co.",
			PolicyNumber: "POL123456",
			Covered:      "Third-Party",
			AmountPaid:   10000.00,
			IssueDate:    "2023-01-01",
			ExpiryDate:   "2024-01-01",
		},
		AdditionalFeatures: AdditionalFeatures{
			AC:              true,
			Radio:           true,
			Camera:          false,
			FireExtinguisher: true,
			FirstAidKit:     true,
		},
	}

	transportPlan.AddVehicle(vehicle)
}
