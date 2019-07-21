@echo off
for /f %%i in ('dir /b I:\Dev\distribute-chat\proto\*.proto') do protoc -I=I:\Dev\distribute-chat\proto I:\Dev\distribute-chat\proto\%%i --go_out=I:\Dev\distribute-chat\proto-gen-go --csharp_out=I:\Dev\distribute-chat\proto-gen-cs
