/**
 * Created by IntelliJ IDEA.
 * User: next
 * Date: 26.03.14
 * Time: 17:11
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

type Salutation struct {    // everything that is Uppercase is public and Lowercase is private
    name string
    greeting string
}

func main() {
//    var s = Salutation{}
//    s.name = "Bob"

//    var s = Salutation{"Andrii","Hello!"}

    var s = Salutation{greeting:"Andrii",name:"Hello!"}

    fmt.Println(s.name,s.greeting)
}


//    msg    := "myMsg"
//    var greeting *string = &msg

//    a,b,c  := 1,false,3
//    fmt.Println(msg,a,b,c)
