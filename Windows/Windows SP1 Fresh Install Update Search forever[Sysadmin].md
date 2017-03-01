If Service Pack 1 is not installed, install it before following this guide.

Download KB-3138612 and save it where you can find it later https://support.microsoft.com/en-us/kb/3138612

Download SUR Tool save it to same place  http://windows.microsoft.com/en-us/windows7/what-is-the-system-update-readiness-tool

Restart the PC and disconnect from internet before Windows loads, 

this is important because at every boot windows will check for updates 

in the background and this will start the checking for updates hang all over again and will prevent 

the install of the downloaded packages until it finishes checking, so disconnecting from the internet 

before Windows loads prevents this.
Once booted install KB-3138612, if reboot is required do so and stay disconnected from internet.

Now install the SUR Tool package, this is a big package and will install many updates along with cleaning up and repairing

the Windows update store. It will also cut down on how many more Windows updates will need to be installed later.

After install of SUR package reboot, connect to internet and do a manual Windows Update, it should work much faster now. 
Manually Update.
If you have other Windows updates issues and the 2 updates above are installed, download this Microsoft Windows Update fixit tool (right click "save link as") run it and select aggressive mode to completely reset Windows updates. Reboot and try Windows Updates from the Control Panel again.

It could help disabling all automatic updates.

Shamelessy taken from http://superuser.com/questions/951960/windows-7-sp1-windows-update-stuck-checking-for-updates
