// 代码生成时间: 2025-09-15 03:54:03
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// EncryptionKey is the key used for AES encryption/decryption.
// It must be 16, 24, or 32 bytes long.
var EncryptionKey = []byte("your-encryption-key")

// Encrypt encrypts the given password using AES.
func Encrypt(text string) (string, error) {
    textBytes := []byte(text)
    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        return "", err
    }

    // PKCS#7 padding
    padding := aes.BlockSize - len(textBytes)%aes.BlockSize
    textBytes = append(textBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)

    // CBC mode
    ciphertext := make([]byte, aes.BlockSize+len(textBytes))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], textBytes)

    // Base64 encode
    return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given password using AES.
func Decrypt(encryptedText string) (string, error) {
    ciphertext, err := base64.URLEncoding.DecodeString(encryptedText)
    if err != nil {
        return "", err
    }

    if len(ciphertext) < aes.BlockSize {
        return "", fmt.Errorf("ciphertext too short")
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        return "", err
    }

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    // Unpadding
    padding := int(ciphertext[len(ciphertext)-1])
    if padding < 1 || padding > aes.BlockSize {
        return "", fmt.Errorf("invalid padding")
    }
    return string(ciphertext[:(len(ciphertext) - padding)]), nil
}

// ResponseData is used to return API responses.
type ResponseData struct {
    Message string `json:"message"`
}

// EncodeResponse is a helper function to encode a response.
func EncodeResponse(c *gin.Context, message string) {
    response := ResponseData{Message: message}
    c.JSON(http.StatusOK, response)
}

func main() {
    router := gin.Default()

    // Define routes
    router.POST("/encrypt", func(c *gin.Context) {
        password := c.PostForm("password")
        if password == "" {
            EncodeResponse(c, "Password is required")
            return
        }

        encryptedPassword, err := Encrypt(password)
        if err != nil {
            EncodeResponse(c, "Encryption failed: "+err.Error())
            return
        }

        EncodeResponse(c, encryptedPassword)
    })

    router.POST("/decrypt", func(c *gin.Context) {
        encryptedPassword := c.PostForm("password")
        if encryptedPassword == "" {
            EncodeResponse(c, "Encrypted password is required")
            return
        }

        decryptedPassword, err := Decrypt(encryptedPassword)
        if err != nil {
            EncodeResponse(c, "Decryption failed: "+err.Error())
            return
        }

        EncodeResponse(c, decryptedPassword)
    })

    // Start the server
    log.Fatal(router.Run(":8080"))
}
