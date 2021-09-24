@echo off
set exepath=D:\WS\proto\bin\
set dispath=D:\WS\fgame\src
%exepath%protoc --go_out=%dispath% cmd.proto
%exepath%protoc --go_out=%dispath% service.proto
%exepath%protoc --go_out=%dispath% server_msg.proto
%exepath%protoc --go_out=%dispath% player_msg.proto
