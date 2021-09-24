@echo off
set exepath=D:\WS\fgame\src\ztests\proto\tools\
set dispath=D:\WS\fgame\src\ztests\
%exepath%protoc --go_out=plugins=grpc:%dispath% hello.proto
