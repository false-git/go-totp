package main

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"fmt"
	"os"
	"time"
)


func main() {
        secret := os.Args[1]
        passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
                Period:    30,
                Skew:      1,
                Digits:    otp.DigitsSix,
                Algorithm: otp.AlgorithmSHA1,
        })
        if err != nil {
                panic(err)
        }
        fmt.Print(passcode)
}
