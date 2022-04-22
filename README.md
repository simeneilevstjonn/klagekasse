# Klagekasse
The solution for businesses not bothering to deal with complaints.

A system that automatically replies to inquiries with a convincing inquiry ID, and then later rejects the complaint after a configurable random delay.

## Example

A user sends a complaint to the complaint email address.

He then immediately receives a confirmation email. This can be customised using PHP, with the following being the default.
```
To: complainer@example.org
From: noreply@example.com
Subject: [3307445] Your inquiry has been received

Thank you for your inquiry. 
Your case has been assigned case ID 3307445.
Please refer to this ID for all further communication.
We will get back to you within five business days. 
Best regards,

Company Name 

```

After a random delay within a configurable bound, the user receives a notice that his complain has been rejected.
```
To: complainer@example.org
From: noreply@example.com
Subject: [3307445] Your inquiry has been closed

Dear customer 
We are referring to your inquiry with ID 3307445. 
We have evaluated your inquiry, and found that it is not applicable. 
Your inquiry has been closed.
Please open a new inquiry if you still need support. 
Best regards,

Company Name 
```

## Installation

### Prerequisites
You need a working Postfix installation, as well as golang, php, at and sendmail.


### Instructions
- Start by cloning the code. The reccomended location is `/usr/lib/klagekasse`, but you could put it anywhere, just make sure to change the paths in `klagekasse.go` and in the Postfix config.
```bash
git clone https://github.com/simeneilevstjonn/klagekasse.git
cd klagekasse
```
- Then do the nescessary changes to the code. You should change the noreply address in both `klagekasse.go` and `rejection.php`. You could also change the PHP files. For `acknowledgement.php`, the ID is passed as the only command line argument, for `rejection.php`, the ID, customer email and subject is sent as arguments in that order.
- Compile the go code and make it executable
```bash
go build klagekasse
chmod 755 klagekasse
```
- Next you should make a user account for the service. You can name it whatever you want, and set whatever password you want, as we will never use that password.
```bash
sudo adduser klagekasse
```
- You should make sure that the user has access to `at`, by adding it's name to the end of `/etc/at.allow`.
- Add the following lines to `/etc/postfix/master.cf` to add the etof transfer to the program, filling in your username and path if changed.
```
etof  unix  -       n       n       -       10      pipe
  flags=Rq user=klagekasse null_sender=
  argv=/usr/lib/klagekasse/klagekasse ${sender}
```
- Next you need to create a transport map. You could map a single address or a wildcard to the transport. Create `/etc/posstfix/transport` with the following line for a basic setup, replacing the address with your desired address.
```
inquiries@example.com etof:
```
- Then you must create a Postfix map of this file by running the following:
```bash
sudo postmap /etc/postfix/transport
```
- In order for Postfix to use your transport map, it must be reffered to in `main.cf`. Add the following line to the end:
```
transport_maps = hash:/etc/postfix/transport
```
- Then, reload Postfix and the service should work.
```bash
sudo service postfix reload
```


