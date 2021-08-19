package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "log"
  "github.com/gorilla/mux"
)

var satellitesCargados []Satellite

func HealthCheck(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"I feel good")
}

func PostSatellites(w http.ResponseWriter, r *http.Request){
  var x, y float32
  var posicionValida bool
  var mensaje string
  var satellites Satellites
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&satellites)
  if(err != nil) {
    panic(err)
  }
  defer r.Body.Close()
  x, y = sendSatellitesToGetLocation(satellites)
  posicionValida = validarPosicionEncontrada(satellites,x,y)
  if(posicionValida) {
    mensaje = sendSatellitesToGetMessage(satellites)
    ResponseOK(w,r,x,y,mensaje)
  } else {
    ResponseNotFound(w,r)
  }
  log.Println(x,y)
  log.Println(posicionValida)
}

func ResponseOK(w http.ResponseWriter, r *http.Request,x float32, y float32, mensaje string){
  position := Position { x, y}

  respuesta := Respuesta{
    Position:   position,
    Message:    mensaje,
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(respuesta)
}

func ResponseNotFound(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNotFound)
}

func PostSatellitesSplit(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  satelliteName := params["name_satellite"]
  var satelliteSplit SatelliteSplit
  var satellite Satellite
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&satelliteSplit)
  if(err != nil) {
    panic(err)
  }
  defer r.Body.Close()
  satellite.Name = satelliteName
  satellite.Distance = satelliteSplit.Distance
  satellite.Message = satelliteSplit.Message
  satellitesCargados = append(satellitesCargados,satellite)
}

func GetSatellitesSplit(w http.ResponseWriter, r *http.Request){
  if(len(satellitesCargados)<3) {
    ResponseNotEnoughInformation(w,r)
  } else {
    calculateLocation(w,r)
  }
}

func ResponseNotEnoughInformation(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"No hay suficiente informacion aun para calcular la posicion de la nave enemiga")
}

func calculateLocation(w http.ResponseWriter, r *http.Request){
  var x, y float32
  var posicionValida bool
  var mensaje string

  x, y = sendSatellitesToGetLocationArray(satellitesCargados)
  posicionValida = validarPosicionEncontradaArray(satellitesCargados,x,y)
  if(posicionValida) {
    mensaje = sendSatellitesToGetMessageArray(satellitesCargados)
    ResponseOK(w,r,x,y,mensaje)
  } else {
    ResponseNotFound(w,r)
  }
  log.Println(x,y)
  log.Println(posicionValida)
}
