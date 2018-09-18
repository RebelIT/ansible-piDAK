# piDAK

---
## DAK Digital Wall Mount Calendar on RaspberryPi3
### Notes:
* Tested on Stretch - NOOBS
* You will need a DAK account (free)
* You will need a gmail account & calendar (free)

### Usage:
* Update hosts with IP or Hostname under the [piCalendar]

  ```
  --ask-sudo-pass may be required if running reboot role due to your local setup
  ansible-playbook piCalendar_setup.yml --ask-vault-pass -i hosts
  ```

### Secrets:
* Upload your ssh key to your pi else ansible will fail

  ```
  wifi_ssid: 'xxx' #wifi ssid
  wifi_psk: 'xxx' #wifi password
  dak_token: 'xxx' #your personal DAK url
  ```
