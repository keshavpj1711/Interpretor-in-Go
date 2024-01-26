package main 

import "fmt"

// Our task is to Write a program to compare two dates entered by user.
// Make a structure named Date to store the elements day, 
// month and year to store the dates. If the dates are equal, 
// display "Dates are equal" otherwise display "Dates are not equal".

type Date struct{
	day int
	month int
	year int
}

func comparingDates(date1 Date, date2 Date) string{
	//Comparing the Year
	if date1.year > date2.year{
		return "Date 1 is later"
	}else if date1.year < date2.year{
		return "Date 2 is later"
	}else{
		if date1.month > date2.month{
			return "Date 1 is later"
		}else if date1.month < date2.month{
			return "Date 2 is later"
		}else{
			if date1.day > date2.day{
				return "Date 1 is later"
			}else if date1.day < date2.day{
				return "Date 2 is later"
			}else{
				return "Dates are equal"
			}
		}
	}
}

// Checking if the year is leap or not
func isLeap(year int) bool{
	if year % 4 == 0{
		if year % 100 == 0{
			if year % 400 == 0{
				return true
			}else{
				return false
			}
		}else{
			return true
		}
	}else{
		return false
	}
}

// Checking for validity of day and month
func validDayMonth(date Date) bool{

	// Checking for valid month
	if date.month > 12{
		fmt.Println("Please enter a valid month for date 1")
		return false
	}

	// Checking for valid day 
	if date.day > 31{
		fmt.Println("Please enter a valid day")
		return false
	}else if date.month == 4 || date.month == 6|| date.month == 9 || date.month == 11{
		if date.day > 30{
			fmt.Println("Please enter a valid day")
			return false
		}
	}else if isLeap(date.year){
		if date.day > 29{
			fmt.Println("Please enter a valid day")
			return false
		}
	}else{
		if date.day > 28{
			fmt.Println("Please enter a valid day")
			return false
		}
	}

	// If all conditions satisfied 
	return true
}

func main(){
	var date1, date2 Date

	for {
        // Taking input of date1
        fmt.Printf("Enter date 1 in the format DD/MM/YYYY: ")
        fmt.Scanf("%d/%d/%d\n", &date1.day, &date1.month, &date1.year)

        if validDayMonth(date1) {
            break // Exit the loop if date1 is valid
        } else {
            continue // Keep prompting for date1 if invalid
        }
    }

    // Get valid date2, preserving date1
    for {
        fmt.Printf("Enter date 2 in the format DD/MM/YYYY: ")
        fmt.Scanf("%d/%d/%d\n", &date2.day, &date2.month, &date2.year)

        if validDayMonth(date2) {
            break // Exit the loop if date2 is valid
        } else {
            continue // Keep prompting for date2 only
        }
    }
	
	// Calling the compare Date Function
	result := comparingDates(date1, date2)
	fmt.Println(result)
}