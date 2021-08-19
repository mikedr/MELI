package main


type Position struct {
  X float32 `json:"x"`
  Y float32 `json:"y"`
}

type Respuesta struct {
  Position Position `json:"position"`
  Message string `json:"message"`
}

type Satellite struct {
  Name string `json:"name"`
  Distance float32 `json:"distance"`
  Message []string `json:"message"`
}

type Satellites struct {
  Satellites []Satellite `json:"satellites"`
}

type SatelliteSplit struct {
  Distance float32 `json:"distance"`
  Message []string `json:"message"`
}

type Positions []Position
