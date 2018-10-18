# piDAK

---
## DAK Digital Wall Mount Calendar on RaspberryPi3
### Notes:
* Ansible for installing and building a DAK wall mounted calendar
* Telegraf for metrics
* GO web webService for controlling your calendar via rest API
* Tested on Stretch - NOOBS
* Assumptions:
   You have a base pi with NOOBS installed and on the network
   You have your SSH key uploaded for user `pi`
   You have gmail account & calendar (free)
   You will need a DAK account (free)
   You have created ./ansible/roles/dak_calendar/vars/secrets.yml vault populated with secrets (noted below)

### Usage:
* Update `hosts` in hosts with your server IP or Hostname
* Update `gopath` in hosts if you want it changed
* Update `repo` in hosts if you use this for another project
* Update `hosts` with any other relevant info specific to your setup

  ```
  --ask-sudo-pass may be required due to your local setup (reboot role requires sudo)
  ansible-playbook piCalendar_setup.yml --ask-vault-pass -i hosts
     or
  ansible-playbook piCalendar_setup.yml -i hosts --ask-vault-pass --ask-sudo-pass
  ```

### Secrets.yml:

  ```
  dak_token: 'xxx' #your personal DAK API token (not the complete http URL... vars adds the http://xxx prefix)
  ```
