package smtp

import (
	"capstone/utils"
	"net/smtp"
	"strings"
)

func SendMail(to []string, subject, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		utils.ReadENV("SMTP_USER"),
		utils.ReadENV("SMTP_PASSWORD"),
		utils.ReadENV("SMTP_HOST"),
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n")
	err := smtp.SendMail(utils.ReadENV("SMTP_HOST")+":"+utils.ReadENV("SMTP_PORT"), auth, utils.ReadENV("SMTP_USER"), to, msg)
	if err != nil {
		return err
	}

	return nil
}

func SendOTP(to []string, otp string) error {
	subject := "OTP Verification"
	body := "Your " + utils.ReadENV("APP_NAME") + " OTP is " + otp + ". Please enter this code to verify your account."

	err := SendMail(to, subject, body)
	if err != nil {
		return err
	}

	return nil
}