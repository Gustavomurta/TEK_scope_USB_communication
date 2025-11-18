```
This project was based on USB communications of the OpenChoice Desktop program from Tektronix.
https://www.tek.com/en/support/software/utility/tektronix-openchoice-desktop-application-tdspcs1--v28

But does not work with the original scope driver - TekVisa.
It is necessary to install the WinUSB driver for the scope.

1) Install the Zadig program:
   https://zadig.akeo.ie/
2) Click on the options tab and click on all devices. Select TEK scope. (ex: TBS1062). 
3) Select the WinUSB driver and click reinstall.
4) After a few seconds, check in the Windows device manager if your scope appears with the Win USB driver.

VISA (Virtual Instrument Software Architecture)

In the TBS1000 programmer's Version 3 manual, you can find the SCPI commands (Standard Commands for Programmable Instruments).

```

<img width="575" height="254" alt="image" src="https://github.com/user-attachments/assets/0887574a-643e-4eba-ae87-40f44a0fe33b" />

