#!/bin/bash

go build -o ./reddit-archivar ./cmds/reddit-archivar/main.go;

./reddit-archivar /r/Malware;
./reddit-archivar /r/LinuxMalware;
./reddit-archivar /r/ReverseEngineering;
./reddit-archivar /r/cybersecurity;
./reddit-archivar /r/netsec;
./reddit-archivar /r/netsecstudents;
./reddit-archivar /r/rootkit;
./reddit-archivar /r/pwned;
./reddit-archivar /r/xss;
./reddit-archivar /r/blackhat;
./reddit-archivar /r/computerforensics;

