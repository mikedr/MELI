package main

import (
  "github.com/sandertv/go-formula/v2"
  "log"
  "math"
  )

  //posicion Kenobi:: [A1,B1]=[-500,-200]
  var A1 = float32(-500)
  var B1 = float32(-200)
  //posicion Skywalker: [A2,B2]=[100,-100]
  var A2 = float32(100)
  var B2 = float32(-100)
  //posicion Sato: [A3,C3]=[500,100]
  var A3 = float32(500)
  var B3 = float32(100)

// input: los satellites enviados por POST
// output
func sendSatellitesToGetLocation(satellites Satellites) (x, y float32){
  var distanciaKenobi = satellites.Satellites[0].Distance
  var distanciaSkywalker float32 = satellites.Satellites[1].Distance
  var distanciaSato float32 = satellites.Satellites[2].Distance
  return GetLocation(distanciaKenobi,distanciaSkywalker,distanciaSato)
}

// input: array de satellites cargados individualmente
// output
func sendSatellitesToGetLocationArray(satellitesCargados []Satellite) (x, y float32){
  var distanciaKenobi = satellitesCargados[0].Distance
  var distanciaSkywalker float32 = satellitesCargados[1].Distance
  var distanciaSato float32 = satellitesCargados[2].Distance
  return GetLocation(distanciaKenobi,distanciaSkywalker,distanciaSato)
}

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
  var D1 = distances[0] //distancia desde el emisor a Kenobi
  var D2 = distances[1] //distancia desde el emisor a Skywalker
  var D3 = distances[2] //distancia desde el emisor a Sato
  return calcular(D1,A1,B1,D2,A2,B2,D3,A3,B3)
}

//input: D1. Distancia desde el emisor al punto Kenobi
//input: A1,B1. Coordenadas de punto Kenobi
//input: D2. Distancia desde el emisor al punto Skywalker
//input: A2,B2. Coordenadas de punto Skywalker
//input: D3. Distancia desde el emisor al punto Sato
//input: A3,B3. Coordenadas de punto Sato
func calcular(D1 float32, A1 float32, B1 float32, D2 float32, A2 float32, B2 float32, D3 float32, A3 float32, B3 float32) (x, y float32) {
  d1 := formula.Var("d1", D1)
  a1 := formula.Var("a1", A1)
  b1 := formula.Var("b1", B1)
  d2 := formula.Var("d2", D2)
  a2 := formula.Var("a2", A2)
  b2 := formula.Var("b2", B2)
  d3 := formula.Var("d3", D3)
  a3 := formula.Var("a3", A3)
  b3 := formula.Var("b3", B3)

  termino01, err := formula.New("pow(d2,2)*(2*b3-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado01 := termino01.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino02, err := formula.New("pow(d1,2)*(2*b3-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado02 := termino02.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino03, err := formula.New("pow(a1,2)*(2*b3-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado03 := termino03.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino04, err := formula.New("pow(b1,2)*(2*b3-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado04 := termino04.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino05, err := formula.New("pow(a2,2)*(2*b3-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado05 := termino05.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino06, err := formula.New("pow(b2,2)*(2*b3-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado06 := termino06.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino07, err := formula.New("pow(d1,2)*(2*b2-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado07 := termino07.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino08, err := formula.New("pow(d3,2)*(2*b2-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado08 := termino08.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino09, err := formula.New("pow(a1,2)*(2*b2-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado09 := termino09.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino10, err := formula.New("pow(b1,2)*(2*b2-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado10 := termino10.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino11, err := formula.New("pow(a3,2)*(2*b2-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado11 := termino11.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino12, err := formula.New("pow(b3,2)*(2*b2-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado12 := termino12.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino13, err := formula.New("4*b3*a1-4*b1*a1")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado13 := termino13.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino14, err := formula.New("4*b3*a2-4*b1*a2")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado14 := termino14.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino15, err := formula.New("4*b2*a1-4*b1*a1")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado15 := termino15.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  termino16, err := formula.New("4*b2*a3-4*b1*a3")
  if err != nil {
      log.Print(err)
      //return
  }
  resultado16 := termino16.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3)

  dividendo := resultado01 - resultado02 + resultado03 + resultado04 - resultado05 - resultado06 + resultado07 - resultado08 - resultado09 - resultado10 + resultado11 + resultado12

  divisor := resultado13 - resultado14 - resultado15 + resultado16

  x = float32(math.Round(dividendo/divisor))

  X := formula.Var("X", x)

  terminoY, err := formula.New("(pow(d1,2)-pow(d3,2)+2*X*a1-pow(a1,2)-pow(b1,2)-2*X*a3+pow(a3,2)+pow(b3,2))/(2*b3-2*b1)")
  if err != nil {
      log.Print(err)
      //return
  }
  y = float32(math.Round(terminoY.MustEval(d1,a1,b1,d2,a2,b2,d3,a3,b3,X)))

	return x,y
}

//input: D1. Distancia desde el emisor al punto Kenobi
//input: A1,B1. Coordenadas de punto Kenobi
//input: D2. Distancia desde el emisor al punto Skywalker
//input: A2,B2. Coordenadas de punto Skywalker
//input: D3. Distancia desde el emisor al punto Sato
//input: A3,B3. Coordenadas de punto Sato
//input: x,y coordenadas calculadas
//output: boolean. true si es valida, false si no es valida
func validarPosicionEncontrada(satellites Satellites, x float32, y float32) (bool) {
  var D1 = satellites.Satellites[0].Distance
  var D2 float32 = satellites.Satellites[1].Distance
  var D3 float32 = satellites.Satellites[2].Distance
  var dist1 = math.Sqrt( math.Pow(float64(x) - float64(A1),2) + math.Pow(float64(y) - float64(B1),2) )
  var firstVerif = (dist1-float64(D1) <= 1)
  var dist2 = math.Sqrt( math.Pow(float64(x) - float64(A2),2) + math.Pow(float64(y) - float64(B2),2) )
  var secondVerif = (dist2-float64(D2) <= 1)
  var dist3 = math.Sqrt( math.Pow(float64(x) - float64(A3),2) + math.Pow(float64(y) - float64(B3),2) )
  var thirdVerif = (dist3-float64(D3) <= 1)
  return (firstVerif&&secondVerif&&thirdVerif)
}

//input: D1. Distancia desde el emisor al punto Kenobi
//input: A1,B1. Coordenadas de punto Kenobi
//input: D2. Distancia desde el emisor al punto Skywalker
//input: A2,B2. Coordenadas de punto Skywalker
//input: D3. Distancia desde el emisor al punto Sato
//input: A3,B3. Coordenadas de punto Sato
//input: x,y coordenadas calculadas
//output: boolean. true si es valida, false si no es valida
func validarPosicionEncontradaArray(satellitesCargados []Satellite, x float32, y float32) (bool) {
  var D1 = satellitesCargados[0].Distance
  var D2 float32 = satellitesCargados[1].Distance
  var D3 float32 = satellitesCargados[2].Distance
  var dist1 = math.Sqrt( math.Pow(float64(x) - float64(A1),2) + math.Pow(float64(y) - float64(B1),2) )
  var firstVerif = (dist1-float64(D1) <= 1)
  var dist2 = math.Sqrt( math.Pow(float64(x) - float64(A2),2) + math.Pow(float64(y) - float64(B2),2) )
  var secondVerif = (dist2-float64(D2) <= 1)
  var dist3 = math.Sqrt( math.Pow(float64(x) - float64(A3),2) + math.Pow(float64(y) - float64(B3),2) )
  var thirdVerif = (dist3-float64(D3) <= 1)
  return (firstVerif&&secondVerif&&thirdVerif)
}
