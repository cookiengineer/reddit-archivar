#!/bin/bash

go build -o ./build/reddit-archivar ./cmds/reddit-archivar/main.go;
cp ./keywords.json ./build/keywords.json;

cd ./build;

./reddit-archivar /r/antivirus;

./reddit-archivar /r/blackhat;
./reddit-archivar /r/blueteamsec;

./reddit-archivar /r/computerforensics;
./reddit-archivar /r/ComputerSecurity;
./reddit-archivar /r/cyber;
./reddit-archivar /r/cybersecurity;
./reddit-archivar /r/Cybersecurity101;
./reddit-archivar /r/dfir;
./reddit-archivar /r/ethicalhacking;
./reddit-archivar /r/ExploitDev;
./reddit-archivar /r/fulldisclosure;

./reddit-archivar /r/HackBloc;
./reddit-archivar /r/hackers;
./reddit-archivar /r/hackersec;
./reddit-archivar /r/hacking;
./reddit-archivar /r/hackintosh;
./reddit-archivar /r/Hacking_Tutorials;
./reddit-archivar /r/hackthebox;
./reddit-archivar /r/HowToHack;

./reddit-archivar /r/Information_Security;
./reddit-archivar /r/InfoSecNews;
./reddit-archivar /r/ISO27001;

./reddit-archivar /r/linuxadmin;
./reddit-archivar /r/LinuxMalware;

./reddit-archivar /r/macsysadmin;
./reddit-archivar /r/Malware;
./reddit-archivar /r/MalwareAnalysis;
./reddit-archivar /r/MalwareDevelopment;
./reddit-archivar /r/MalwareResearch;
./reddit-archivar /r/metasploit;

./reddit-archivar /r/nessus;
./reddit-archivar /r/netsec;
./reddit-archivar /r/netsecstudents;
./reddit-archivar /r/networking;
./reddit-archivar /r/NetworkSecurity;
./reddit-archivar /r/nmap;

./reddit-archivar /r/opendirectories;
./reddit-archivar /r/OPNsenseFirewall;

./reddit-archivar /r/pwned;
./reddit-archivar /r/redteam;
./reddit-archivar /r/redteamsec;
./reddit-archivar /r/REGames;
./reddit-archivar /r/rootkit;
./reddit-archivar /r/ReverseEngineering;

./reddit-archivar /r/suricata;

./reddit-archivar /r/websec;
./reddit-archivar /r/websecurity;
./reddit-archivar /r/owasp;
./reddit-archivar /r/securityCTF;
./reddit-archivar /r/sysadmin;

./reddit-archivar /r/xss;
./reddit-archivar /r/zeroday;
