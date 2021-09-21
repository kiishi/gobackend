package main

import "os"

var MONGO_URI = os.Getenv("MONGO_URI")
var PORT = os.Getenv("PORT")
