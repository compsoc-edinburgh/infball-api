package base

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

type EmailStruct struct {
	Name      string
	AuthToken string
	OrderID   string
}

func SendTicketEmail(c *gin.Context, mailgun mailgun.Mailgun, name, to_address, orderID, authToken string) (_ bool) {
	var tpl bytes.Buffer
	if err := Email.Execute(&tpl, EmailStruct{
		Name:      name,
		OrderID:   orderID,
		AuthToken: authToken,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Ticket was bought but email generation was unsuccessful. Please email us for assistance.",
		})
		return
	}

	message := mailgun.NewMessage(
		"Informatics Ball <infball@comp-soc.com>",
		"Your Informatics Ball ticket! [#"+orderID+"]",
		"",
		to_address,
	)
	message.SetHtml(tpl.String())

	_, _, err := mailgun.Send(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Ticket was bought but an email was not sent. Please email us for assistance.",
		})
		return
	}

	return true
}

var Email = template.Must(template.New("ball-email").Parse(`
	<!doctype html>
	<html>
	  <head>
		<meta name="viewport" content="width=device-width" />
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
		<title>Simple Transactional Email</title>
		<style>
		  
		  img {
			border: none;
			-ms-interpolation-mode: bicubic;
			max-width: 100%; }
	
		  body {
			background-color: #f6f6f6;
			font-family: sans-serif;
			-webkit-font-smoothing: antialiased;
			font-size: 14px;
			line-height: 1.4;
			margin: 0;
			padding: 0;
			-ms-text-size-adjust: 100%;
			-webkit-text-size-adjust: 100%; }
	
		  table {
			border-collapse: separate;
			mso-table-lspace: 0pt;
			mso-table-rspace: 0pt;
			width: 100%; }
			table td {
			  font-family: sans-serif;
			  font-size: 14px;
			  vertical-align: top; }
	
		  
	
		  .body {
			background-color: #f6f6f6;
			width: 100%; }
	
		  
		  .container {
			display: block;
			Margin: 0 auto !important;
			
			max-width: 580px;
			padding: 10px;
			width: 580px; }
	
		  
		  .content {
			box-sizing: border-box;
			display: block;
			Margin: 0 auto;
			max-width: 580px;
			padding: 10px; }
	
		  
		  .main {
			background: #ffffff;
			border-radius: 3px;
			width: 100%; }
	
		  .wrapper {
			box-sizing: border-box;
			padding: 20px; }
	
		  .content-block {
			padding-bottom: 10px;
			padding-top: 10px;
		  }
	
		  .footer {
			clear: both;
			Margin-top: 10px;
			text-align: center;
			width: 100%; }
			.footer td,
			.footer p,
			.footer span,
			.footer a {
			  color: #999999;
			  font-size: 12px;
			  text-align: center; }
	
		  
		  h1,
		  h2,
		  h3,
		  h4 {
			color: #000000;
			font-family: sans-serif;
			font-weight: 400;
			line-height: 1.4;
			margin: 0;
			Margin-bottom: 30px; }
	
		  h1 {
			font-size: 35px;
			font-weight: 300;
			text-align: center;
			text-transform: capitalize; }
	
		  p,
		  ul,
		  ol {
			font-family: sans-serif;
			font-size: 14px;
			font-weight: normal;
			margin: 0;
			Margin-bottom: 15px; }
			p li,
			ul li,
			ol li {
			  list-style-position: inside;
			  margin-left: 5px; }
	
		  a {
			color: #3498db;
			text-decoration: underline; }
	
		  
		  .btn {
			box-sizing: border-box;
			width: 100%; }
			.btn > tbody > tr > td {
			  padding-bottom: 15px; }
			.btn table {
			  width: auto; }
			.btn table td {
			  background-color: #ffffff;
			  border-radius: 5px;
			  text-align: center; }
			.btn a {
			  background-color: #ffffff;
			  border: solid 1px #3498db;
			  border-radius: 5px;
			  box-sizing: border-box;
			  color: #3498db;
			  cursor: pointer;
			  display: inline-block;
			  font-size: 14px;
			  font-weight: bold;
			  margin: 0;
			  padding: 12px 25px;
			  text-decoration: none;
			  text-transform: capitalize; }
	
		  .btn-primary table td {
			background-color: #3498db; }
	
		  .btn-primary a {
			background-color: #3498db;
			border-color: #3498db;
			color: #ffffff; }
	
		  
		  .last {
			margin-bottom: 0; }
	
		  .first {
			margin-top: 0; }
	
		  .align-center {
			text-align: center; }
	
		  .align-right {
			text-align: right; }
	
		  .align-left {
			text-align: left; }
	
		  .clear {
			clear: both; }
	
		  .mt0 {
			margin-top: 0; }
	
		  .mb0 {
			margin-bottom: 0; }
	
		  .preheader {
			color: transparent;
			display: none;
			height: 0;
			max-height: 0;
			max-width: 0;
			opacity: 0;
			overflow: hidden;
			mso-hide: all;
			visibility: hidden;
			width: 0; }
	
		  .powered-by a {
			text-decoration: none; }
	
		  hr {
			border: 0;
			border-bottom: 1px solid #f6f6f6;
			Margin: 20px 0; }
	
		  
		  @media only screen and (max-width: 620px) {
			table[class=body] h1 {
			  font-size: 28px !important;
			  margin-bottom: 10px !important; }
			table[class=body] p,
			table[class=body] ul,
			table[class=body] ol,
			table[class=body] td,
			table[class=body] span,
			table[class=body] a {
			  font-size: 16px !important; }
			table[class=body] .wrapper,
			table[class=body] .article {
			  padding: 10px !important; }
			table[class=body] .content {
			  padding: 0 !important; }
			table[class=body] .container {
			  padding: 0 !important;
			  width: 100% !important; }
			table[class=body] .main {
			  border-left-width: 0 !important;
			  border-radius: 0 !important;
			  border-right-width: 0 !important; }
			table[class=body] .btn table {
			  width: 100% !important; }
			table[class=body] .btn a {
			  width: 100% !important; }
			table[class=body] .img-responsive {
			  height: auto !important;
			  max-width: 100% !important;
			  width: auto !important; }}
	
		  
		  @media all {
			.ExternalClass {
			  width: 100%; }
			.ExternalClass,
			.ExternalClass p,
			.ExternalClass span,
			.ExternalClass font,
			.ExternalClass td,
			.ExternalClass div {
			  line-height: 100%; }
			.apple-link a {
			  color: inherit !important;
			  font-family: inherit !important;
			  font-size: inherit !important;
			  font-weight: inherit !important;
			  line-height: inherit !important;
			  text-decoration: none !important; }
			.btn-primary table td:hover {
			  background-color: #34495e !important; }
			.btn-primary a:hover {
			  background-color: #34495e !important;
			  border-color: #34495e !important; } }
	
		</style>
	  </head>
	  <body class="" style="background-color:#f6f6f6;font-family:sans-serif;-webkit-font-smoothing:antialiased;font-size:14px;line-height:1.4;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;padding-top:0;padding-bottom:0;padding-right:0;padding-left:0;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;" >
		<table border="0" cellpadding="0" cellspacing="0" class="body" style="border-collapse:separate;mso-table-lspace:0pt;mso-table-rspace:0pt;background-color:#f6f6f6;width:100%;" >
		  <tr>
			<td style="font-family:sans-serif;font-size:14px;vertical-align:top;" >&nbsp;</td>
			<td class="container" style="font-family:sans-serif;font-size:14px;vertical-align:top;display:block;Margin:0 auto !important;max-width:580px;padding-top:10px;padding-bottom:10px;padding-right:10px;padding-left:10px;width:580px;" >
			  <div class="content" style="box-sizing:border-box;display:block;Margin:0 auto;max-width:580px;padding-top:10px;padding-bottom:10px;padding-right:10px;padding-left:10px;" >
	
				<!-- START CENTERED WHITE CONTAINER -->
				<span class="preheader" style="color:transparent;display:none;height:0;max-height:0;max-width:0;opacity:0;overflow:hidden;mso-hide:all;visibility:hidden;width:0;" >Join us on the 8th of March 2018.</span>
				<table class="main" style="border-collapse:separate;mso-table-lspace:0pt;mso-table-rspace:0pt;background-color:#ffffff;background-image:none;background-repeat:repeat;background-position:top left;background-attachment:scroll;border-radius:3px;width:100%;" >
	
				  <!-- START MAIN CONTENT AREA -->
				  <tr>
					<td class="wrapper" style="font-family:sans-serif;font-size:14px;vertical-align:top;box-sizing:border-box;padding-top:20px;padding-bottom:20px;padding-right:20px;padding-left:20px;" >
					  <table border="0" cellpadding="0" cellspacing="0" style="border-collapse:separate;mso-table-lspace:0pt;mso-table-rspace:0pt;width:100%;" >
						<tr>
						  <td style="font-family:sans-serif;font-size:14px;vertical-align:top;" >
							<p style="font-family:sans-serif;font-size:14px;font-weight:normal;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;Margin-bottom:15px;" >Dear {{.Name}},</p>
							<p style="font-family:sans-serif;font-size:14px;font-weight:normal;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;Margin-bottom:15px;" >Thank you for getting a ticket for the Informatics Ball.</p>
							<table border="0" cellpadding="0" cellspacing="0" class="btn btn-primary" style="border-collapse:separate;mso-table-lspace:0pt;mso-table-rspace:0pt;box-sizing:border-box;width:100%;" >
							  <tbody>
								<tr>
								  <td align="left" style="font-family:sans-serif;font-size:14px;vertical-align:top;padding-bottom:15px;" >
									<table border="0" cellpadding="0" cellspacing="0" style="border-collapse:separate;mso-table-lspace:0pt;mso-table-rspace:0pt;width:auto;" >
									  <tbody>
										<tr>
										  <td style="font-family:sans-serif;font-size:14px;vertical-align:top;border-radius:5px;text-align:center;background-color:#3498db;" > <a href="https://comp-soc.com/infball-ticket?id={{.OrderID}}" target="_blank" style="border-width:1px;border-style:solid;border-radius:5px;box-sizing:border-box;cursor:pointer;display:inline-block;font-size:14px;font-weight:bold;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;padding-top:12px;padding-bottom:12px;padding-right:25px;padding-left:25px;text-decoration:none;text-transform:capitalize;background-color:#3498db;border-color:#3498db;color:#ffffff;" >View Ticket</a> </td>
										</tr>
									  </tbody>
									</table>
								  </td>
								</tr>
							  </tbody>
							</table>
							<p style="font-family:sans-serif;font-size:14px;font-weight:normal;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;Margin-bottom:15px;" >If you would like to modify your ticket you can use the following auth token to make changes:
							<code>{{.AuthToken}}</code>
							</p>
							<p style="font-family:sans-serif;font-size:14px;font-weight:normal;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;Margin-bottom:15px;" >See you at 7PM on 8th April 2018!</p>
							<p style="font-family:sans-serif;font-size:14px;font-weight:normal;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;Margin-bottom:15px;" >- CompSoc</p>
						  </td>
						</tr>
					  </table>
					</td>
				  </tr>
	
				<!-- END MAIN CONTENT AREA -->
				</table>
	
				<!-- START FOOTER -->
				<div class="footer" style="clear:both;Margin-top:10px;text-align:center;width:100%;" >
				  <table border="0" cellpadding="0" cellspacing="0" style="border-collapse:separate;mso-table-lspace:0pt;mso-table-rspace:0pt;width:100%;" >
					<tr>
					  <td class="content-block" style="font-family:sans-serif;vertical-align:top;padding-bottom:10px;padding-top:10px;color:#999999;font-size:12px;text-align:center;" >
						<p style="font-family:sans-serif;font-weight:normal;margin-top:0;margin-bottom:0;margin-right:0;margin-left:0;Margin-bottom:15px;color:#999999;font-size:12px;text-align:center;" >You are receiving this because you have a ticket to the Informatics Ball.<br>If you think this
	is a mistake please email us at hello@comp-soc.com.</p>
						<span class="" style="color:#999999;font-size:12px;text-align:center;" >CompSoc Edinburgh, Informatics Forum, Edinburgh, EH8 9AB</span>
					  </td>
					</tr>
					<tr>
					  <td class="content-block powered-by" style="font-family:sans-serif;vertical-align:top;padding-bottom:10px;padding-top:10px;color:#999999;font-size:12px;text-align:center;" >
						CompSoc &hearts; You!
					  </td>
					</tr>
				  </table>
				</div>
				<!-- END FOOTER -->
	
			  <!-- END CENTERED WHITE CONTAINER -->
			  </div>
			</td>
			<td style="font-family:sans-serif;font-size:14px;vertical-align:top;" >&nbsp;</td>
		  </tr>
		</table>
	  </body>
	</html>
	
`))
