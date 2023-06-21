package config

import "os"

var DB_USER = os.Getenv("DB_USER")
var DB_PASS = os.Getenv("DB_PASS")
var DB_NAME = os.Getenv("DB_NAME")
var DB_PORT = os.Getenv("DB_PORT")
var JWT_SECRET = os.Getenv("JWT_SECRET")
var MOIRAI_SECRET = os.Getenv("MOIRAI_SECRET")
