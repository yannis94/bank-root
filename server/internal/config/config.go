package config

import "os"

var DB_USER = os.Getenv("DB_USER")
var DB_PASS = os.Getenv("DB_PASSWORD")
var DB_NAME = os.Getenv("DB_NAME")
