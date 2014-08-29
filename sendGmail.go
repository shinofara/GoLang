package main
 
import (
        "log"
        "net/smtp"
)
 
func main() {
        // Set up authentication information.
        auth := smtp.PlainAuth(
                "",
                "info@finatext.com", //ユーザ名
                "Hayato1014", //パスワード,
                "smtp.gmail.com",
        )
        // Connect to the server, authenticate, set the sender and recipient,
        // and send the email all in one step.
        err := smtp.SendMail(
                "smtp.gmail.com:587",
                auth,
                "info@finatext.com", //送信元
                []string{"shinofara@gmail.com"}, //宛先
                []byte("This is the email body."),
        )
        if err != nil {
                log.Fatal(err)
        }
}
