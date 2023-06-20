@ECHO off
FOR /F "tokens=1,2 delims=\=" %%G IN (.env) DO (set %%G=%%H)
