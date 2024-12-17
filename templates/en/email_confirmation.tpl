<html>
  <head>
    <meta name="viewport" content="width=device-width"/>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>Confirm Your Account</title>
    <style type="text/css">
      body{
        margin: 0 auto;
        padding: 0;
        min-width: 100%;
        font-family: sans-serif;
      }
      table{
        margin: 50px 0 50px 0;
      }
      .speza{
        width: 356px;
        height: 96px;
        text-align: center;
      }
      .image{
        width: 600px;
        height: 250;
        text-align: center;
      }
      .header{
        height: 40px;
        text-align: center;
        text-transform: uppercase;
        font-size: 24px;
        font-weight: normal;
      }
      .content{
        height: 100px;
        font-size: 18px;
        line-height: 30px;
        text-align: center;
      }
      .subscribe{
        height: 70px;
        text-align: center;
      }
      .button{
        text-align: center;
        font-size: 18px;
        font-family: sans-serif;
        font-weight: bold;
        padding: 0 30px 0 30px;
      }
      .button a{
        color: #FFFFFF;
        text-decoration: none;
      }
      .buttonwrapper{
        margin: 0 auto;
      }
      .footer{
        text-align: center;
        height: 100px;
        font-size: 14px;
      }
      .footer a{
        color: #000000;
        text-decoration: none;
        font-style: normal;
      }
    </style>
  </head>
  <body>
    <table bgcolor="#FFFFFF" width="100%" border="0" cellspacing="0" cellpadding="0">
      <tr class="header">
        <td style="padding: 20px;">
        <img src="https://i.imgur.com/tH3ti99.jpeg" class="abc">
        </td>
      </tr>
      <tr class="header">
        <td style="padding: 10px;">
          Welcome to abc!
        </td>
      </tr>
      <tr class="header">
        <td>
          Welcome
        </td>
      </tr>
      <tr class="header">
        <td style="padding: 20px;">
        <img src="https://i.imgur.com/mbejVX3.png" class="image">
        </td>
      </tr>
      <tr class="content">
        <td style="padding:10px;">
          <p>
            Hi <b>{{ .record.user.email }}</b>, <br/>
            To activate your account, click "Confirm"
          </p>
        </td>
      </tr>
      <tr class="subscribe">
        <td style="padding: 20px 0 0 0;">
          <table bgcolor="#823179" border="0" cellspacing="0" cellpadding="0" class="buttonwrapper">
            <tr>
              <td class="button" height="45">
                <a href="{{ .EmailConfirmationURI }}" target="_blank">CONFIRM</a>
              </td>
            </tr>
          </table>
        </td>
      </tr>
      <tr class="footer">
        <td>
          Need help?
          Check out the <a href="https://abc.zendesk.com/hc/en-us" style="color:#823179" target="_blank">help panel</a> for more information <br/>
        </td>
      </tr>
    </table>
  </body>
</html>