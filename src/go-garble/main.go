package main

import "os"

func emit(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}

func main() {
	emit("go-canary::alpha::c2.example.net:8443/api/v1/beacon")
	emit("go-canary::bravo::C:\\ProgramData\\Updater\\config.json")
	emit("go-canary::charlie::Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	emit("go-canary::delta::SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run")
	emit("go-canary::echo::session-key-4f9a")
	emit("go-canary::foxtrot::cmd.exe /c whoami /all")
	emit("go-canary::golf::https://cdn.example.org/update/agent.bin")
	emit("go-canary::hotel::%APPDATA%\\Microsoft\\Crypto\\keys.dat")
	emit("go-canary::india::POST /gate.php HTTP/1.1")
	emit("go-canary::juliett::wallet.dat")
}
