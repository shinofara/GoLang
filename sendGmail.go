package main
 
import (
        "log"
        "net/smtp"
)
 
func main() {
        // Set up authentication information.
        auth := smtp.PlainAuth(
                "",
                "info@finatext.com", //"<username>", // foo@gmail.com
                "Hayato1014", //"<password>",
                "smtp.gmail.com",
        )
        // Connect to the server, authenticate, set the sender and recipient,
        // and send the email all in one step.
        err := smtp.SendMail(
                "smtp.gmail.com:587",
                auth,
                "info@finatext.com", //"送信アドレス", //foo@gmail.com
                //[]string{"受信アドレス"},
                []string{"shinofara@gmail.com"},
                []byte("This is the email body."),
        )
        if err != nil {
                log.Fatal(err)
        }
}
