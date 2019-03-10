/*
Towers of Hanoi: In the classic problem of the Towers of Hanoi, you have 3 towers and N disks of
different sizes which can slide onto any tower. The puzzle starts with disks sorted in ascending order
of size from top to bottom (Le., each disk sits on top of an even larger one). You have the following
constraints:
(1) Only one disk can be moved at a time.
(2) A disk is slid off the top of one tower onto another tower.
(3) A disk cannot be placed on top of a smaller disk.
Write a program to move the disks from the first tower to the last using stacks. 
*/
package main

import "fmt"

type solver interface {
    play(int)
}
   
// towers is example of type satisfying solver interface
type towers struct {
    // an empty struct
}
  
// play is sole method required to implement solver type
func (t *towers) play(n int) {    
    t.moveN(n, 1, 2, 3)
}

// recursive algorithm
func (t *towers) moveN(n, from, to, via int) {
    if n > 0 {
        t.moveN(n-1, from, via, to)
        t.moveM(from, to)
        t.moveN(n-1, via, to, from)
    }
}
 
func (t *towers) moveM(from, to int) {
    fmt.Println("Move disk from rod", from, "to rod", to)
}
 
func main() {
    var t solver    
    t = new(towers) // type towers must satisfy solver interface
    t.play(2)
}