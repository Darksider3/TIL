# OpenWRT - change lost root password over telnet in recovery mode

For an Archer C7 v2 you've to reboot your router, press the RESET-Button till the systems LED is flashing rapidly.
Then you can connect to it, in my case it was 192.168.1.1, by setting your local interface into that subnet. Note: Dont forget the netmask!
```
sudo ifconfig enp3s0 192.168.1.101 netmask 255.255.255.0
```

When you've got no idea which IP nor subnet it has, you can watch for it's ARP Packet with netdiscover or wireshark. Dont forget to plug your PC into the router first and disconnect it from any working DHCP:
```
netdiscover -S -f -i $INTERFACE
```

Then you'll be able to connect, when the light is still flashing rapidly, with telnet:
```
telnet 192.168.1.1
```
And you're greeted with the normal interface! 
```
BusyBox v1.11.2 (2009-12-02 06:19:32 UTC) built-in shell (ash)
Enter 'help' for a list of built-in commands.

  _______                     ________        __
 |       |.-----.-----.-----.|  |  |  |.----.|  |_
 |   -   ||  _  |  -__|     ||  |  |  ||   _||   _|
 |_______||   __|_____|__|__||________||__|  |____|
          |__| W I R E L E S S   F R E E D O M
 KAMIKAZE (8.09.2, r18961) -------------------------
  * 10 oz Vodka       Shake well with ice and strain
  * 10 oz Triple sec  mixture into 10 shot glasses.
  * 10 oz lime juice  Salute!
 ---------------------------------------------------
root@OpenWrt:~#
```
Mount root:
```
mount_root
```
and set your new password!

```
passwd
```
, finally reboot! Wait a minute or two until you try to login
```
reboot
```
Voila! Your new password is set for root! :)
