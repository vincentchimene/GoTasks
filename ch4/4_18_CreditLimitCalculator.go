package main

import "fmt"

func main(){
    var accountNumber, beginningBalance, charges, credits, creditLimit, newBalance int
    
    for {

        fmt.Print("ENter -1 to end program or Enter account number: ")
        fmt.Scan(&accountNumber)
        if accountNumber == -1{
            break
        }
        
        fmt.Println("ENter beginning balance: ")
        fmt.Scan(&beginningBalance)
        
        fmt.Println("ENter value for charges: ")
        fmt.Scan(&charges)
        
        fmt.Println("ENter value for credits: ")
        fmt.Scan(&credits)
        
        fmt.Println("ENter value for credit limit: ")
        fmt.Scan(&creditLimit)
        
        newBalance = beginningBalance + charges - credits
        
        fmt.Println("New Balance: ", newBalance)
        
        if newBalance > creditLimit{
            fmt.Println("Credit limit exceeded!")
        }else {
        
            fmt.Println("Within Credit Limit")
        }
                  
    }
    fmt.Println("You ended the program.")
    
           
}
