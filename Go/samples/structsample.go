package main

import "fmt"

type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}

type Wheel struct {
    Circle Circle
    Spokes int
}

func main()  {

  var w Wheel
  w.Circle.Center.X = 10
  w.Circle.Center.Y = 20

  fmt.Printf("Centre is at %d\t%d" , w.Circle.Center.X, w.Circle.Center.Y)

}
