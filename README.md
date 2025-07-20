VEKO DOME - STEALTH MODE 
Advanced Network Identity Stealth System
VEKO DOME adalah tool CLI berbasis Golang yang dirancang untuk memberikan perlindungan anonimitas jaringan tingkat tinggi dengan target 99% anonymity. Tool ini melindungi identitas pengguna dari tracking, fingerprinting, dan WAF detection.
 PERINGATAN PENTING
VEKO DOME hanya untuk penggunaan yang ETIS:

 Perlindungan privasi dan anonimitas
 Penelitian keamanan dan penetration testing
 Load testing dan performance testing
 Web scraping yang legitimate
 Penelitian akademis

 JANGAN gunakan untuk:

Aktivitas ilegal atau melanggar hukum
Menyerang sistem tanpa izin
Melanggar terms of service website
Spamming atau abuse

 Fitur Utama
1. Multi-Layer Proxy System

Rotasi proxy otomatis dari file list
Support HTTP & SOCKS5 proxy
Authentikasi proxy (username/password)
Tor integration (socks5://127.0.0.1:9050)

2. Advanced TLS Fingerprinting

Menggunakan uTLS library untuk spoof JA3 fingerprint
Simulasi TLS ClientHello dari browser nyata
Support: Chrome, Firefox, Safari, Edge
Menghindari deteksi fingerprint default Golang

3. DNS over HTTPS (DoH)

DNS queries melalui HTTPS (Cloudflare 1.1.1.1)
Mencegah DNS leak dan monitoring
Bypass DNS censorship

4. User-Agent Randomization

Rotasi User-Agent dari database lengkap
User-Agent realistis dari browser populer
Sinkronisasi dengan TLS fingerprint

5. HTTP Header Spoofing

Headers realistis: Accept, Referer, Accept-Language
Anti-fingerprinting headers
Randomized header values

6. Real-time Monitoring

Dashboard CLI dengan status real-time
IP leak detection
Anonymity status indicator
Connection statistics

 Instalasi
Prasyarat:
bash# Install Go (v1.19+)
go version

# Install uTLS dependency
go mod init veko-dome
go get github.com/refraction-networking/utls
Kompilasi:
bash# Clone atau copy file main.go
go build -o veko-dome main.go

# Atau untuk Windows
go build -o veko-dome.exe main.go
 Konfigurasi
1. File Konfigurasi (config.json)
json{
  "proxy_file": "proxylist.txt",
  "rotation_interval_seconds": 30,
  "tor_enabled": false,
  "tor_proxy": "socks5://127.0.0.1:9050",
  "doh_server": "https://1.1.1.1/dns-query",
  "useragent_file": "useragents.txt",
  "tls_fingerprint": "chrome"
}
2. Proxy List (proxylist.txt)
# Format: IP:PORT atau IP:PORT:USER:PASS
# HTTP Proxies
123.456.789.10:8080
123.456.789.11:3128:username:password

# SOCKS5 Proxies  
socks5://127.0.0.1:9050
socks5://123.456.789.12:1080:user:pass
3. User Agent List (useragents.txt)
File ini akan dibuat otomatis dengan user-agent populer, atau Anda bisa menambahkan custom user-agent.
 Penggunaan
Menjalankan VEKO DOME:
bash# Basic usage
./veko-dome

# Untuk Windows
veko-dome.exe
Mode Tor:

Install dan jalankan Tor Browser atau Tor service
Set tor_enabled: true di config.json
Jalankan VEKO DOME

Testing Anonymity:
bash# Test anonymity secara manual
curl -x socks5://127.0.0.1:9050 https://api.ipify.org
 Dashboard Monitoring
VEKO DOME menampilkan dashboard real-time:
┌────────────────── VEKO DOME - STEALTH MODE ──────────────────┐
│ Current IP     : 185.22.xxx.xxx                              │
│ Status         :  TOR Active                                │
│ DNS Mode       : DoH (1.1.1.1)                              │
│ TLS Spoof      : CHROME                                      │
│ User Agent     : Mozilla/5.0 (Windows NT 10.0...)          │
│ Proxy Count    : 15                                         │
│ Rotation       : Every 30 seconds                           │
└──────────────────────────────────────────────────────────────┘
 Tingkat Perlindungan
LayerProtectionStatusIP AddressProxy/Tor RotationDNS QueriesDNS-over-HTTPSTLS FingerprintuTLS SpoofingHTTP HeadersRandomizationUser AgentBrowser SimulationConnection TimingRandomized
 Tips Penggunaan
1. Proxy Quality:

Gunakan proxy berkualitas tinggi dan terpercaya
Rotasi proxy setiap 30-60 detik
Test proxy secara berkala

2. Tor Configuration:
bash# Install Tor (Ubuntu/Debian)
sudo apt install tor

# Edit /etc/tor/torrc
ControlPort 9051
SocksPort 9050

# Restart Tor
sudo systemctl restart tor
3. Optimal Settings:

Rotation interval: 30-60 seconds
TLS fingerprint: Chrome (paling umum)
DoH server: Cloudflare (1.1.1.1)

 Testing & Verification
IP Leak Test:
bash# Test multiple IP services
curl https://api.ipify.org
curl https://ifconfig.me/ip
curl https://icanhazip.com
DNS Leak Test:
bash# Check DNS resolution
nslookup google.com 1.1.1.1
TLS Fingerprint Test:
Gunakan tools seperti JA3er.com atau Wireshark untuk memverifikasi TLS fingerprint.
 Struktur Direktori
/veko-dome/
├── main.go                 # Source code utama
├── config.json            # Konfigurasi
├── proxylist.txt          # Daftar proxy
├── useragents.txt         # Daftar user agent
├── veko-dome              # Binary (Linux/Mac)
├── veko-dome.exe          # Binary (Windows)
└── README.md              # Dokumentasi
 Troubleshooting
Masalah Umum:

Proxy tidak berfungsi:

Periksa format proxy di proxylist.txt
Test proxy secara manual
Periksa authentikasi proxy


Tor tidak connect:

Pastikan Tor service berjalan
Periksa port 9050 terbuka
Restart Tor service


DNS leak:

Periksa konfigurasi DoH
Test dengan DNS leak checker online



 Legal & Ethical Use
VEKO DOME dibuat untuk:

Melindungi privasi pengguna
Penelitian keamanan yang legitimate
Testing dan development

Pengguna bertanggung jawab atas penggunaan tool ini sesuai hukum dan regulasi yang berlaku.
 Contributing
Jika ingin berkontribusi:

Fork repository
Buat feature branch
Test thoroughly
Submit pull request

 License
Open source - gunakan secara bertanggung jawab.

Made with  for privacy and security research
VEKO DOME - Your Digital Stealth Companion
