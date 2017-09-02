package main

import "fmt"

type Celsius float64
type Farenheit float64

func main()  {
  //Celsius c1 = 212.44
  //fmt.Println(c1.String())
  waterboilF:=CToF(100)
  fmt.Println(waterboilF)

  waterFreezeC Celsius = 0
  fmt.Println(waterFreezeC.String())
}

func (c Celsius) String() string {
  return fmt.Sprintf("%g C", c)
}

func CToF (c Celsius) Farenheit{
  return Farenheit(c*9/5 +32)
}
