package email

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"testing"
)

// Client simple email client support ssl
type Client struct {
	user     string
	addr     string
	nickName string
	isSSL    bool
	auth     smtp.Auth
}

// NewEmailClient  New new email client
func NewEmailClient(user, password, nickName, host, port string, isSsl bool) *Client {
	ec := &Client{
		user:  user,
		addr:  fmt.Sprintf("%s:%s", host, port),
		isSSL: isSsl,
		auth:  smtp.PlainAuth("", user, password, host),
	}
	if nickName == "" {
		ec.nickName = user
	} else {
		ec.nickName = nickName
	}
	return ec
}

func (ec *Client) generateEmailMsg(toUser []string, subject, content string) []byte {
	return ec.generateEmailMsgByte(toUser, subject, []byte(content))
}

func (ec *Client) generateEmailMsgByte(toUser []string, subject string, body []byte) []byte {
	msgStr := fmt.Sprintf("To: %s\r\nFrom: %s<%s>\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n",
		strings.Join(toUser, ","), ec.nickName, ec.user, subject)
	return append([]byte(msgStr), body...)
}

func (ec *Client) sendMailTLS(toUser []string, msg []byte) error {
	host, _, _ := net.SplitHostPort(ec.addr)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", ec.addr, tlsconfig)
	if err != nil {
		return fmt.Errorf("DialConn:%v", err)
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("Client:generateClient:%v", err)
	}
	defer client.Close()
	if ec.auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(ec.auth); err != nil {
				return fmt.Errorf("Client:clientAuth:%v", err)
			}
		}
	}
	if err = client.Mail(ec.user); err != nil {
		return fmt.Errorf("Client:clientMail:%v", err)
	}

	for _, addr := range toUser {
		if err = client.Rcpt(addr); err != nil {
			return fmt.Errorf("Client:Rcpt:%v", err)
		}
	}
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("Client:%v", err)
	}
	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("Client:WriterBody:%v", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("Client:CloseBody:%v", err)
	}
	return client.Quit()
}

func (ec *Client) sendMail(toUser []string, msg []byte) error {
	return smtp.SendMail(ec.addr, ec.auth, ec.user, toUser, msg)
}

// SendEmail send email by string content
func (ec *Client) SendEmail(toUser []string, subject string, content string) error {
	msg := ec.generateEmailMsg(toUser, subject, content)
	if ec.isSSL {
		return ec.sendMailTLS(toUser, msg)
	}
	return ec.sendMail(toUser, msg)
}

// SendEmailByte send email by byte body
func (ec *Client) SendEmailByte(toUser []string, subject string, body []byte) error {
	msg := ec.generateEmailMsgByte(toUser, subject, body)
	if ec.isSSL {
		return ec.sendMailTLS(toUser, msg)
	}
	return ec.sendMail(toUser, msg)
}

func Test_SSL(t *testing.T) {
	client := NewEmailClient("3330125541@qq.com", "hhplrasikbegchfd", "韦达的小助手", "smtp.qq.com", "465", true)
	if err := client.SendEmail([]string{"weida@ones.cn"}, "email send test", "this is a test email for sending.\r\nJust Testing"); err != nil {
		fmt.Println(err)
	}
}