# piDAK

---
## DAK Digital Wall Mount Calendar on RaspberryPi3
### Notes:
* Ansible for installing and building
* go web webService for controlling your calendar via rest API
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

  ```
  --ask-sudo-pass may be required due to your local setup
  ansible-playbook piCalendar_setup.yml -i hosts
     or
  ansible-playbook piCalendar_setup.yml -i hosts --ask-sudo-pass
  ```

### Secrets.yml:

  ```
  dak_token: 'xxx' #your personal DAK url
  ```
