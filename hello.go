/**
 * Created by IntelliJ IDEA.
 * User: next
 * Date: 26.03.14
 * Time: 17:11
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

func main() {
    msg    := "myMsg"
    var greeting *string = &msg

//    a,b,c  := 1,false,3
//    fmt.Println(msg,a,b,c)

    fmt.Println(msg,*greeting)
}
