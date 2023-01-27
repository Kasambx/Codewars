package main

import ("fmt")

func Solution(word string) string {
    rns := []rune(word) //convert to rune 
    for i , j :=0, len(rns)-1; i<j; i,j = i+1, j-1 {
      
      
      //swap the letters in the string
      //like first with last and so on 
        rns[i], rns[j] = rns[j], rns[i]
    }
    
    //return the reversed string
    return string(rns)
}
func main() {
    // reversing the string
    str := "world"
  
    //returns the reversed string
    strRev := Solution(str)
    fmt.Println(str)
    fmt.Println(strRev)
}


