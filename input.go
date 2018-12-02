package main

import(
    "math/rand"
    "strings"
    "strconv"
    "bufio"
    "time"
    "fmt"
    "log"
    "os"
)


func main(){

seconds := time.Now().Unix()

rand.Seed(seconds)
target := rand.Intn(100) + 1
fmt.Println("I've chosen a randon number between 1 and 100.")
fmt.Println("Can you guess it?")
fmt.Println(target)

reader := bufio.NewReader(os.Stdin)
success := false

for guesses := 0; guesses < 10; guesses++ {
    fmt.Println("You have", 10 - guesses, " guesses left.")
    fmt.Print("Make a guess: ")
    input, error := reader.ReadString('\n')
    if error != nil {
        log.Fatal(error)
    }
    
    input = strings.TrimSpace(input)
    guess, error := strconv.Atoi(input)
    if error != nil {
        log.Fatal(error)
    }
    
    if guess < target {
        fmt.Println("Oops! Your guess was LOW.")
    } else if guess > target {
        fmt.Println("Oops! Your guess was HIGH.")
    } else {
        success = true
        fmt.Println("Good job! You guessed it.")
        break
    }
    
    
} // end of for guesses

if !success {
    fmt.Println("Sorry, you didn't guess my number. It was:", target)
}


/*   GRADE APP
fmt.Print("Enter a grade: ")

reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
if err != nil {
  log.Fatal(err)   
}    

input = strings.TrimSpace(input)
grade, err := strconv.ParseFloat(input, 64)
if err != nil {
  log.Fatal(err)   
} 

var status string 
if grade >= 60 {
    status = "passing"
} else {
    status = "failing"
}

fmt.Println("A grade of ", grade, "is", status)
 */
 
 
 
} // end of main