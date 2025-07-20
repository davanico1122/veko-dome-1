# VEKO DOME - STEALTH MODE

VEKO DOME adalah sistem perlindungan dan anonimitas tingkat lanjut yang dirancang untuk:

* Menyembunyikan identitas asli (IP publik)
* Melindungi aktivitas scraping, crawling, dan load testing
* Menghindari pendeteksian bot
* Mengelola proxy, Tor, DoH, dan user-agent secara otomatis

> Cocok digunakan bersama VEKO GO untuk pengetesan trafik anonim dan sah secara maksimal.

---

## FITUR UNGGULAN

* Dukungan HTTP Proxy dan SOCKS5
* Integrasi mode TOR (127.0.0.1:9050)
* DNS-over-HTTPS (DoH via 1.1.1.1)
* Rotasi otomatis IP dan User-Agent
* Simulasi fingerprint TLS seperti Chrome, Firefox, dll
* Monitoring IP dan status secara real-time
* File konfigurasi otomatis

---

## STRUKTUR FILE

```
veko-dome/
├── main.go               # Kode utama Veko Dome
├── config.json           # Konfigurasi awal (otomatis dibuat)
├── proxylist.txt         # Daftar proxy (otomatis dibuat)
├── useragents.txt        # Daftar user agent (otomatis dibuat)
├── README.md             # Panduan penggunaan
└── veko-dome             # Binary hasil build (setelah dibuild)
```

---

## CARA BUILD MANUAL

### SYARAT

* Sudah menginstall Go: [https://go.dev/dl/](https://go.dev/dl/)

### BUILD (Linux/Mac/Windows):

```bash
# Linux/macOS:
go build -o veko-dome main.go

# Windows:
go build -o veko-dome.exe main.go
```

### CROSS-BUILD (jika ingin membuat versi .exe dari Linux/macOS):

```bash
GOOS=windows GOARCH=amd64 go build -o veko-dome.exe main.go
```

---

## CARA MENJALANKAN VEKO DOME

```bash
./veko-dome
```

Jika berhasil, terminal akan menampilkan informasi seperti ini:

```
┌────────────────── VEKO DOME - STEALTH MODE ──────────────────┐
│ Current IP     : 185.xxx.xxx.xxx                            │
│ Status         : Proxy Active                               │
│ DNS Mode       : DoH (1.1.1.1)                               │
│ TLS Spoof      : CHROME                                     │
│ User Agent     : Mozilla/5.0 ...                            │
│ Proxy Count    : 12                                         │
│ Rotation       : Every 30 seconds                           │
└──────────────────────────────────────────────────────────────┘
```

> VEKO DOME akan terus berjalan & memutar identitas (IP, UA, TLS) otomatis.

Untuk **berhenti**, cukup tekan **Ctrl + C** pada terminal.

---

## FILE KONFIGURASI OTOMATIS

Saat pertama dijalankan, file berikut otomatis dibuat:

### `config.json`

```json
{
  "proxy_file": "proxylist.txt",
  "rotation_interval_seconds": 30,
  "tor_enabled": false,
  "tor_proxy": "socks5://127.0.0.1:9050",
  "doh_server": "https://1.1.1.1/dns-query",
  "useragent_file": "useragents.txt",
  "tls_fingerprint": "chrome"
}
```

Ubah nilai sesuai kebutuhan:

* `tor_enabled`: true untuk aktifkan Tor
* `rotation_interval_seconds`: waktu rotasi IP/UA
* `tls_fingerprint`: `chrome`, `firefox`, `safari`, `edge`

---

## FORMAT PROXYLIST

Isi `proxylist.txt` dengan format:

```
# HTTP
123.456.78.90:8080
123.456.78.91:3128:user:pass

# SOCKS5
socks5://127.0.0.1:9050
socks5://123.123.123.123:1080:user:pass
```

---

## FORMAT USERAGENT FILE

Contoh isi `useragents.txt`:

```
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 ...
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 ...
```

> Jika file tidak ada, akan dibuat otomatis dengan beberapa default UA.

---

## INTEGRASI DENGAN VEKO GO

* Jalankan VEKO DOME terlebih dahulu
* Pastikan koneksi keluar sudah terproteksi (lihat IP, status)
* Jalankan VEKO GO (load test) secara bersamaan

```bash
./veko-go --url https://target.site --vus 100 --duration 30s
```

---

## PERINGATAN ETIS

Tool ini **tidak untuk digunakan secara sembarangan**.

> VEKO DOME + VEKO GO = senjata kuat. Gunakan hanya untuk:

* Penelitian keamanan
* Pengujian performa situs milik sendiri
* Audit jaringan yang legal

**Jangan gunakan untuk kegiatan ilegal, penyerangan, atau abuse.**

---

## KONTRIBUSI & LISENSI

* Proyek terbuka dan bebas digunakan
* Beri kredit dan jangan ubah nama "VEKO"
* Boleh di-clone dan dibagikan

---

## PENUTUP

VEKO DOME adalah lapisan perlindungan trafik yang tangguh, cocok dipakai sendiri maupun terintegrasi ke VEKO GO.

Selalu gunakan dengan tanggung jawab dan etika.

> Untuk pertanyaan atau dukungan, silakan hubungi melalui GitHub!

---
