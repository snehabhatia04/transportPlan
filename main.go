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

// Student represents a student in the transport plan
type Student struct {
	Name          string
	VehicleChoice int // 1 for Vehicle 1, 2 for Vehicle 2
	RouteChoice   int // 1 for Route 1, 2 for Route 2
}

// TransportPlan represents the overall transport plan
type TransportPlan struct {
	Vehicles []Vehicle
	Drivers  []Driver
	Routes   []Route
	Students []Student
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

// AddStudent adds a student to the transport plan
func (tp *TransportPlan) AddStudent(s Student) {
	tp.Students = append(tp.Students, s)
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

	fmt.Println("Students:")
	for _, s := range tp.Students {
		vehicle := "Vehicle 1"
		if s.VehicleChoice == 2 {
			vehicle = "Vehicle 2"
		}

		route := "Route 1"
		if s.RouteChoice == 2 {
			route = "Route 2"
		}
		if s.RouteChoice == 3 {
			route = "Route 3"
		}
		if s.RouteChoice == 4 {
			route = "Route 4"
		}
		if s.RouteChoice == 5 {
			route = "Route 5"
		}
		if s.RouteChoice == 6 {
			route = "Route 6"
		}
		if s.RouteChoice == 7 {
			route = "Route 7"
		}
		if s.RouteChoice == 8 {
			route = "Route 8"
		}
		fmt.Printf("  Student Name: %s, Chose: %s, %s\n", s.Name, vehicle, route)
	}
}


func main() {
	transportPlan := TransportPlan{}

	// Adding a vehicle
	vehicle1 := Vehicle{
		VehicleType:       "Bus",
		Manufacture:       "XYZ Motors",
		RouteNumber:       "R1",
		RegistrationNumber: "AB1234CD",
		OwnerName:         "Sanjay Patel",
		OwnerPhoneNumber:  "1234567890",
		OwnerAddress:      "123B, Sindhi Camp",
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

	vehicle2 := Vehicle{
		VehicleType:       "Van",
		Manufacture:       "ABC Motors",
		RouteNumber:       "R2",
		RegistrationNumber: "XY5678ZT",
		OwnerName:         "Mehul Gandhi",
		OwnerPhoneNumber:  "0987654321",
		OwnerAddress:      "456, Maharana Chowk",
		FuelType:          "Petrol",
		SeatingCapacity:   15,
		InsuranceDetails: Insurance{
			CompanyName:  "Insurance Co.",
			PolicyNumber: "POL654321",
			Covered:      "Zero Depreciation",
			AmountPaid:   15000.00,
			IssueDate:    "2023-02-01",
			ExpiryDate:   "2024-02-01",
		},
		AdditionalFeatures: AdditionalFeatures{
			AC:              true,
			Radio:           false,
			Camera:          true,
			FireExtinguisher: true,
			FirstAidKit:     true,
		},
	}

	// Adding vehicles to the transport plan
	transportPlan.AddVehicle(vehicle1)
	transportPlan.AddVehicle(vehicle2)

	// Adding routes
	route1 := Route{
		RouteNumber:   "R1",
		StartLocation: "School",
		EndLocation:   "Gandevi Road",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "07:00 AM", DropTime: "07:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "07:20 AM", DropTime: "07:30 AM", StopNumber: 2},
		},
	}

	route2 := Route{
		RouteNumber:   "R2",
		StartLocation: "School",
		EndLocation:   "Grid road",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "08:00 AM", DropTime: "08:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "08:20 AM", DropTime: "08:30 AM", StopNumber: 2},
		},
	}

	route3 := Route{
		RouteNumber:   "R3",
		StartLocation: "School",
		EndLocation:   "Mankodia",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "07:00 AM", DropTime: "07:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "07:20 AM", DropTime: "07:30 AM", StopNumber: 2},
		},
	}

	route4 := Route{
		RouteNumber:   "R4",
		StartLocation: "School",
		EndLocation:   "Station Road",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "07:00 AM", DropTime: "07:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "07:20 AM", DropTime: "07:30 AM", StopNumber: 2},
		},
	}

	route5 := Route{
		RouteNumber:   "R5",
		StartLocation: "School",
		EndLocation:   "Junathana",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "07:00 AM", DropTime: "07:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "07:20 AM", DropTime: "07:30 AM", StopNumber: 2},
		},
	}

	route6 := Route{
		RouteNumber:   "R6",
		StartLocation: "School A",
		EndLocation:   "Kabilpore",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "07:00 AM", DropTime: "07:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "07:20 AM", DropTime: "07:30 AM", StopNumber: 2},
		},
	}

	route7 := Route{
		RouteNumber:   "R7",
		StartLocation: "School",
		EndLocation:   "Lunciqui",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "07:00 AM", DropTime: "07:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "07:20 AM", DropTime: "07:30 AM", StopNumber: 2},
		},
	}

	route8 := Route{
		RouteNumber:   "R8",
		StartLocation: "School",
		EndLocation:   "Asha Bagh",
		PickupPoints: []PickupPoint{
			{Name: "Stop 1", PickupTime: "07:00 AM", DropTime: "07:10 AM", StopNumber: 1},
			{Name: "Stop 2", PickupTime: "07:20 AM", DropTime: "07:30 AM", StopNumber: 2},
		},
	}

	// Adding routes to the transport plan
	transportPlan.AddRoute(route1)
	transportPlan.AddRoute(route2)
	transportPlan.AddRoute(route3)
	transportPlan.AddRoute(route4)
	transportPlan.AddRoute(route5)
	transportPlan.AddRoute(route6)
	transportPlan.AddRoute(route7)
	transportPlan.AddRoute(route8)

	// Adding students
	// student1 := Student{
	// 	Name:          "Asmi Jain",
	// 	var i,j  int

    //     fmt.Print("Choose a vehicle: ")
    //     fmt.Scan(&i)
    //     fmt.Println("Your vehicle is:", i)
	// 	fmt.Print("Choose a Route: ")
    //     fmt.Scan(&j)
    //     fmt.Println("Your route is:", j)
        
	// 	VehicleChoice: 1, // Chooses Vehicle 1
	// 	RouteChoice:   1, // Chooses Route 1
	// }

	// student2 := Student{
	// 	Name:          "Bhoomi Bachav",
	// 	VehicleChoice: 2, // Chooses Vehicle 2
	// 	RouteChoice:   4, // Chooses Route 2
	// }

	// student3 := Student{
	// 	Name:          "Simran Bhatia",
	// 	VehicleChoice: 1, // Chooses Vehicle 1
	// 	RouteChoice:   6, // Chooses Route 2
	// }

	// student4 := Student{
	// 	Name:          "Nischay Bhatia",
	// 	VehicleChoice: 2, // Chooses Vehicle 2
	// 	RouteChoice:   1, // Chooses Route 1
	// }

	// // Adding students to the transport plan
	// transportPlan.AddStudent(student1)
	// transportPlan.AddStudent(student2)
	// transportPlan.AddStudent(student3)
	// transportPlan.AddStudent(student4)

// User input for adding students
var numStudents int
fmt.Print("Enter the number of students to add: ")
fmt.Scan(&numStudents)

for i := 0; i < numStudents; i++ {
	var name string
	var vehicleChoice, routeChoice int

	fmt.Printf("Enter name of student %d: ", i+1)
	fmt.Scan(&name)

	fmt.Print("Choose Vehicle (1 for Vehicle 1, 2 for Vehicle 2): ")
	fmt.Scan(&vehicleChoice)

	fmt.Print("Choose Route (1 for Route 1, 2 for Route 2,3 for Route 3,4 for Route 4,5 for Route 5,6 for Route 6): ")
	fmt.Scan(&routeChoice)

	student := Student{
		Name:          name,
		VehicleChoice: vehicleChoice,
		RouteChoice:   routeChoice,
	}

	// Adding the student to the transport plan
	transportPlan.AddStudent(student)
}

	// Displaying the transport plan
	transportPlan.DisplayTransportPlan()
}

