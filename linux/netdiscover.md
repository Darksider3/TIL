# Netdiscover - in case you don't know which subnet a device is in!

It's pretty simple. Netdiscover just sniffs for ARP Broadcasts and shows you the corresponding IP. 
```
> sudo netdiscover -S -f -i enp3s0

 Currently scanning: 172.27.147.0/16   |   Screen View: Unique Hosts                                                                                                                    
                                                                                                                                                                                        
 3 Captured ARP Req/Rep packets, from 3 hosts.   Total size: 180                                                                                                                        
 _____________________________________________________________________________
   IP            At MAC Address     Count     Len  MAC Vendor / Hostname      
 -----------------------------------------------------------------------------
 192.xxx.x.2     x:x:x:x:x:1e      1      60  TP-LINK TECHNOLOGIES CO.,LTD.                                                                                                        
 192.xxx.x.1     x:x:x:x:x:ec      1      60  Sercomm Corporation.                                                                                                                 
 192.xxx.x.1     x:x:x:x:x:38      1      60  Hewlett Packard
```
