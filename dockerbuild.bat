@echo off
cd /d "%~dp0"
cls

docker build -t bwlee2018/virtual_exporter1:1.0.0 . --push
REM docker run -it --name virtual_exporter1 -p:bwlee2018/virtual_exporter1:1.0.0