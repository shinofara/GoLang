package main
 
import (
        "bytes"
        "log"
        "net/smtp"
)
 
func main() {
        // Connect to the remote SMTP server.
        c, err := smtp.Dial("localhost:25")
        if err != nil {
                log.Fatal(err)
        }
        // Set the sender and recipient.
        c.Mail("sender@localhost")      // メールの送り主を指定
        c.Rcpt("shinofara@gmail.com")      // 受信者を指定
        //c.Rcpt("yyyyyy@gmail.com")      // Ccにしたい場合も同様に指定
 
        // Send the email body.
        wc, err := c.Data()
        if err != nil {
                log.Fatal(err)
        }
        defer wc.Close()
        //ToにするかCcにするかBccにするかはDATAメッセージ次第
        buf := bytes.NewBufferString("To:shinofara@gmail.com")     
        buf.WriteString("\r\n") // DATA メッセージはCRLFのみ
        //buf.WriteString("Cc:yyyyyy@gmail.com")
        buf.WriteString("\r\n")
        buf.WriteString("Subject:this is Subject") //件名
        buf.WriteString("\r\n")
        buf.WriteString("This is the email body.")
        if _, err = buf.WriteTo(wc); err != nil {
                log.Fatal(err)
        }
 
        c.Quit()  //メールセッションの終了
}
