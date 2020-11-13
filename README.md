# slack-emoji-backup

1. Open chrome to `https://{your_slack_space}.slack.com/customize/emoji`
and press F12, then type
`window.prompt("your api token is: ", TS.boot_data.api_token)` in console. Copy the xoxs token from the popup window.

2. Paste xoxs token into main.go, then run it. The backuped emoji will be in `backup-emojis`.