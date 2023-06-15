package config

import "os"

var DB_USER = os.Getenv("DB_USER")
var DB_PASS = os.Getenv("DB_PASS")
var DB_NAME = os.Getenv("DB_NAME")
var DB_PORT = os.Getenv("DB_PORT")
