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

const (
    PI = 3.14
    Lang = "GO"
    A = iota
    B = iota
    C = iota
)

const (
    F = iota // auto increment depends of consand declared
    G
    H
)

func main() {


    fmt.Println(PI,Lang,A,B,C)
    fmt.Println(F,G,H)
}

////////////////////////////////

//    var s = Salutation{}
//    s.name = "Bob"

//    var s = Salutation{"Andrii","Hello!"}

//    var s = Salutation{greeting:"Andrii",name:"Hello!"}

////////////////////////////////

//    msg    := "myMsg"
//    var greeting *string = &msg

//    a,b,c  := 1,false,3
//    fmt.Println(msg,a,b,c)
