# 🛡️ VEKO DOME – STEALTH MODE

VEKO DOME adalah alat pengaman identitas koneksi internet tingkat tinggi yang dirancang untuk melakukan **anonimisasi koneksi keluar** melalui rotasi **proxy**, **TLS fingerprint spoofing**, **user-agent acak**, dan **DNS over HTTPS**.  
Berjalan di terminal (CLI), ringan, dan tidak memerlukan browser.

> ⚠️ Untuk penggunaan etis seperti scraping sah, load testing, dan riset keamanan. Bukan untuk penyalahgunaan.

---

## ✨ FITUR UTAMA

- 🔄 Rotasi Proxy otomatis (HTTP / SOCKS5 / TOR)
- 🧅 Dukungan TOR Proxy
- 🔐 DNS-over-HTTPS (DoH) untuk menyembunyikan query DNS
- 🧬 Spoofing TLS Fingerprint (Chrome, Firefox, Safari, Edge)
- 🎭 Rotasi User-Agent acak
- 🧾 Header penyamaran (Referer, Accept, DNT, dll)
- 📡 Monitoring IP Publik
- 📟 CLI Dashboard Real-time
- 📁 Konfigurasi fleksibel melalui `config.json`

---

## 📦 STRUKTUR FILE

```
veko-dome/
├── main.go
├── build.sh
├── config.json
├── proxylist.txt
├── useragents.txt
└── builds/
```

---

## ⚙️ CARA BUILD

```bash
bash build.sh
```

📌 Ini akan:
- Mengunduh dependensi Go
- Compile binary untuk Windows, Linux, macOS
- Membuat file `config.json`, `proxylist.txt`, dan `useragents.txt`

---

## 🚀 CARA MENJALANKAN

### Windows
```bash
./veko-dome.exe
```

### Linux / macOS
```bash
./veko-dome
```

📌 Program akan menampilkan dashboard CLI real-time dan mulai melakukan rotasi identitas.

---

## ⏹️ CARA MENGHENTIKAN

Tekan `Ctrl + C` di terminal.

Program akan berhenti dengan aman.

---

## 🔧 PENGATURAN (config.json)

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

| Kunci               | Fungsi                                                                 |
|---------------------|------------------------------------------------------------------------|
| `rotation_interval_seconds` | Interval waktu rotasi identitas (detik)                   |
| `tor_enabled`       | Gunakan TOR (jika true, proxy manual akan diabaikan)                  |
| `doh_server`        | Server DNS over HTTPS (default: Cloudflare 1.1.1.1)                    |
| `tls_fingerprint`   | Spoof TLS handshake (chrome, firefox, safari, edge)                   |

---

## 🛰️ TAMPILAN DASHBOARD

```text
┌────────────────── VEKO DOME - STEALTH MODE ──────────────────┐
│ Current IP     : 185.23.104.50                               │
│ Status         : Proxy Active                                │
│ DNS Mode       : DoH (1.1.1.1)                               │
│ TLS Spoof      : CHROME                                      │
│ User Agent     : Mozilla/5.0 (Windows NT 10.0...)            │
│ Proxy Count    : 47                                          │
│ Rotation       : Every 30 seconds                            │
└──────────────────────────────────────────────────────────────┘
```

Update setiap 10 detik secara otomatis.

---

## 🧪 TEST ANONIMITAS (Opsional)

Tambahkan di `main.go`:
```go
vd.TestAnonymity()
```

Menampilkan:
- IP dari beberapa layanan
- Header aktif
- Deteksi DNS leak

---

## 📁 PENGELOLAAN FILE

| File             | Fungsi                     |
|------------------|----------------------------|
| `proxylist.txt`  | Daftar proxy aktif (acak)  |
| `useragents.txt` | Daftar User-Agent acak     |
| `config.json`    | Konfigurasi semua fitur    |

---

## 💡 CONTOH PENGGUNAAN

- Anonim scraping website publik
- Melindungi identitas saat load testing
- Riset koneksi outbound anonim
- Pengujian infrastruktur yang membutuhkan IP acak

---

## 🛑 PERINGATAN

> VEKO DOME dibuat **untuk tujuan sah & riset etis**.  
> Dilarang digunakan untuk:
> - Penipuan / DDoS
> - Peretasan / pembajakan
> - Aktivitas ilegal

Pengembang tidak bertanggung jawab atas penyalahgunaan.

---

## 🔓 LISENSI

Proyek ini open-source dan dapat dikembangkan bebas, selama mematuhi etika dan hukum.

---

## 🧠 KONTRIBUSI / PENGEMBANGAN

Untuk fitur lanjutan seperti:
- Web GUI controller
- Integrasi VPN external
- Sistem log atau output JSON

Silakan buat issue atau fork dan pull request.

---

## 🦾 Dibuat dengan 🔒 oleh Veko Project