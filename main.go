package main

import (
	"fmt"
	"log"
	"os"

	gecko "github.com/HackLike-co/Gecko/Gecko"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("gecko", "encrypt/encode payloads")

	payloadFile := parser.String("f", "file", &argparse.Options{Required: true, Help: "path to .bin file"})
	outputFile := parser.String("o", "output", &argparse.Options{Required: false, Help: "file to write output to. if none specified, will print to stdout"})

	aesEncodeCmd := parser.NewCommand("aes", "AES Encrypt the Shellcode with randomly generated key and iv using AES256-CBC")

	rc4EncodedCmd := parser.NewCommand("rc4", "RC4 Encrypt the Shellcode with randomly generated key of x number of bytes")
	keyLen := rc4EncodedCmd.Int("k", "key-len", &argparse.Options{Required: false, Default: 32, Help: "Number of bytes in the rc4 key"})

	// parse arguments
	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// should we write output to file or print to stdout
	var doStdout bool = true
	if *outputFile != "" {
		doStdout = false
	}

	// read payload file
	payload, err := os.ReadFile(*payloadFile)
	if err != nil {
		log.Fatal(err)
	}

	// aes encryption
	if aesEncodeCmd.Happened() {
		// generate key
		key, err := gecko.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		// generate iv
		iv, err := gecko.GenerateIV()
		if err != nil {
			log.Fatal(err)
		}

		// encrypt payload
		ePayload, err := gecko.AES_CBCEncrypt(payload, key, iv)
		if err != nil {
			log.Fatal(err)
		}

		// generate final output string
		output := fmt.Sprintf("%s\n\n%s\n\n%s", gecko.C_FormatArray(ePayload, "Payload"), gecko.C_FormatArray(key, "Key"), gecko.C_FormatArray(iv, "IV"))

		if doStdout {
			fmt.Print(output)
		} else {
			outFile, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatal(err)
			}

			outFile.Write([]byte(output))
		}
	} else if rc4EncodedCmd.Happened() {
		// generate key
		key, err := gecko.GenerateSecureBytes(*keyLen)
		if err != nil {
			log.Fatal(err)
		}

		// encrypt payload
		ePayload, err := gecko.RC4_Encrypt(payload, key)
		if err != nil {
			log.Fatal(err)
		}

		// generate final output string
		output := fmt.Sprintf("%s\n\n%s", gecko.C_FormatArray(ePayload, "Payload"), gecko.C_FormatArray(key, "Key"))

		if doStdout {
			fmt.Print(output)
		} else {
			outFile, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatal(err)
			}

			outFile.Write([]byte(output))
		}
	}

}
