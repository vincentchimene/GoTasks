package main

import "fmt"

func main(){
    var gallons, miles int
    var totalMiles, totalGallons int
    
    for {
        fmt.Print("Enter miles driven or -1 to quit: ")
        fmt.Scan(&miles)
        
        if miles == -1{
            break
        }
        
        fmt.Print("ENter value for gallons: ")
        fmt.Scan(&gallons)
        
        totalMiles += miles
        totalGallons += gallons
        milesPerGallon := float64(miles) / float64(gallons)
        totalMilesPerGallon := float64(totalMiles) / float64(totalGallons)  
        
        fmt.Printf("MilesPerGallon: %.2f\n", milesPerGallon)
        fmt.Printf("TotalMilesPerGallon: %.2f\n", totalMilesPerGallon)
              
    }
    
    fmt.Println("You ended the program.")

    
}
