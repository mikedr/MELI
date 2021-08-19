package main

import(
  "strings"
  )

// input: los satellites enviados por POST
// output
func sendSatellitesToGetMessage(satellites Satellites)(msg string) {
  return GetMessage(satellites.Satellites[0].Message,satellites.Satellites[1].Message,satellites.Satellites[2].Message)
}

// input: array de satellites cargados individualmente
// output
func sendSatellitesToGetMessageArray(satellitesCargados []Satellite)(msg string) {
  return GetMessage(satellitesCargados[0].Message,satellitesCargados[1].Message,satellitesCargados[2].Message)
}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string){
  var eosKenobi bool = false
  var eosSkywalker bool = false
  var eosSato bool = false

  var mensaje string
  var palabrasKenobi = messages[0]
  var palabrasSkywalker = messages[1]
  var palabrasSato = messages[2]
  var i = 0
  for {
    if len(palabrasKenobi)>=(i+1) {
      if(strings.Compare(palabrasKenobi[i], "\"\"") != 0) {
        mensaje += (palabrasKenobi[i]+" ")
      }
    } else {
      eosKenobi = true
    }

    if len(palabrasSkywalker)>=(i+1) {
      if(strings.Compare(palabrasSkywalker[i], "\"\"") != 0) {
        mensaje += (palabrasSkywalker[i]+" ")
      }
    } else {
      eosSkywalker = true
    }

    if len(palabrasSato)>=(i+1) {
      if(strings.Compare(palabrasSato[i], "\"\"") != 0) {
        mensaje += (palabrasSato[i]+" ")
      }
    } else {
      eosSato = true
    }

    if (eosKenobi && eosSkywalker && eosSato) {
      break
    }
    i++
  }
  return mensaje
}
