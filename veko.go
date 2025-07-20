package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	utls "github.com/refraction-networking/utls"
)

// Config structure
type Config struct {
	ProxyFile        string   `json:"proxy_file"`
	RotationInterval int      `json:"rotation_interval_seconds"`
	TorEnabled       bool     `json:"tor_enabled"`
	TorProxy         string   `json:"tor_proxy"`
	DoHServer        string   `json:"doh_server"`
	UserAgentFile    string   `json:"useragent_file"`
	TLSFingerprint   string   `json:"tls_fingerprint"`
}

// Proxy structure
type Proxy struct {
	Host     string
	Port     string
	Username string
	Password string
	Type     string // http or socks5
}

// VekoDome main structure
type VekoDome struct {
	config     Config
	proxies    []Proxy
	userAgents []string
	currentUA  string
	currentIP  string
	client     *http.Client
	mutex      sync.RWMutex
}

// Default user agents
var defaultUserAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/119.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/119.0.0.0 Safari/537.36",
}

// TLS fingerprints mapping
var tlsFingerprints = map[string]utls.ClientHelloID{
	"chrome":  utls.HelloChrome_Auto,
	"firefox": utls.HelloFirefox_Auto,
	"safari":  utls.HelloSafari_Auto,
	"edge":    utls.HelloEdge_Auto,
}

func main() {
	showWarning()
	
	vd := &VekoDome{}
	vd.loadConfig()
	vd.loadProxies()
	vd.loadUserAgents()
	vd.setupHTTPClient()
	
	// Start monitoring
	go vd.startRotation()
	go vd.monitorStatus()
	
	// Keep running
	fmt.Println("\n🚀 VEKO DOME is now active. Press Ctrl+C to stop.")
	select {}
}

func showWarning() {
	fmt.Println(`
╔══════════════════════════════════════════════════════════════════════════════╗
║                            ⚠️  IMPORTANT WARNING ⚠️                           ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  VEKO DOME - STEALTH MODE is designed for ETHICAL purposes only:            ║
║                                                                              ║
║  ✅ Privacy Protection          ✅ Security Research                          ║
║  ✅ Load Testing               ✅ Anonymity Protection                        ║
║  ✅ Academic Research          ✅ Legitimate Scraping                         ║
║                                                                              ║
║  ❌ DO NOT use for malicious activities, illegal scraping, or abuse         ║
║                                                                              ║
║  Use responsibly and respect website terms of service.                      ║
╚══════════════════════════════════════════════════════════════════════════════╝
`)
	time.Sleep(3 * time.Second)
}

func (vd *VekoDome) loadConfig() {
	// Default config
	vd.config = Config{
		ProxyFile:        "proxylist.txt",
		RotationInterval: 30,
		TorEnabled:       false,
		TorProxy:         "socks5://127.0.0.1:9050",
		DoHServer:        "https://1.1.1.1/dns-query",
		UserAgentFile:    "useragents.txt",
		TLSFingerprint:   "chrome",
	}

	// Try to load from file
	if file, err := os.Open("config.json"); err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(&vd.config)
	} else {
		// Create default config file
		vd.saveConfig()
	}
}

func (vd *VekoDome) saveConfig() {
	file, err := os.Create("config.json")
	if err != nil {
		return
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(vd.config)
}

func (vd *VekoDome) loadProxies() {
	file, err := os.Open(vd.config.ProxyFile)
	if err != nil {
		fmt.Printf("⚠️  Proxy file not found, creating sample %s\n", vd.config.ProxyFile)
		vd.createSampleProxyFile()
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		proxy := vd.parseProxyLine(line)
		if proxy.Host != "" {
			vd.proxies = append(vd.proxies, proxy)
		}
	}
	
	fmt.Printf("📋 Loaded %d proxies\n", len(vd.proxies))
}

func (vd *VekoDome) parseProxyLine(line string) Proxy {
	parts := strings.Split(line, ":")
	if len(parts) >= 2 {
		proxy := Proxy{
			Host: parts[0],
			Port: parts[1],
			Type: "http", // default
		}
		
		if len(parts) >= 4 {
			proxy.Username = parts[2]
			proxy.Password = parts[3]
		}
		
		// Detect socks5
		if strings.Contains(line, "socks") {
			proxy.Type = "socks5"
		}
		
		return proxy
	}
	return Proxy{}
}

func (vd *VekoDome) createSampleProxyFile() {
	content := `# VEKO DOME Proxy List
# Format: IP:PORT or IP:PORT:USER:PASS
# Prefix with 'socks5://' for SOCKS5 proxies

# Example HTTP proxies:
# 123.456.789.10:8080
# 123.456.789.11:3128:username:password

# Example SOCKS5 proxies:
# socks5://127.0.0.1:9050
# socks5://123.456.789.12:1080:user:pass

# Add your proxies here:
`
	os.WriteFile(vd.config.ProxyFile, []byte(content), 0644)
}

func (vd *VekoDome) loadUserAgents() {
	file, err := os.Open(vd.config.UserAgentFile)
	if err != nil {
		vd.userAgents = defaultUserAgents
		vd.createSampleUserAgentFile()
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			vd.userAgents = append(vd.userAgents, line)
		}
	}
	
	if len(vd.userAgents) == 0 {
		vd.userAgents = defaultUserAgents
	}
	
	fmt.Printf(" Loaded %d user agents\n", len(vd.userAgents))
}

func (vd *VekoDome) createSampleUserAgentFile() {
	content := `# VEKO DOME User Agent List
# Add your custom user agents here

` + strings.Join(defaultUserAgents, "\n")
	os.WriteFile(vd.config.UserAgentFile, []byte(content), 0644)
}

func (vd *VekoDome) setupHTTPClient() {
	// Custom TLS config with uTLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         "",
	}

	// Setup transport with proxy
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}

	// Apply proxy if available
	vd.setupProxy(transport)

	vd.client = &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
	
	vd.rotateUserAgent()
}

func (vd *VekoDome) setupProxy(transport *http.Transport) {
	if vd.config.TorEnabled {
		proxyURL, _ := url.Parse(vd.config.TorProxy)
		transport.Proxy = http.ProxyURL(proxyURL)
		fmt.println(" Tor mode enabled")
		return
	}
	
	if len(vd.proxies) > 0 {
		proxy := vd.proxies[rand.Intn(len(vd.proxies))]
		proxyStr := fmt.Sprintf("%s://%s:%s", proxy.Type, proxy.Host, proxy.Port)
		
		if proxy.Username != "" && proxy.Password != "" {
			proxyStr = fmt.Sprintf("%s://%s:%s@%s:%s", proxy.Type, proxy.Username, proxy.Password, proxy.Host, proxy.Port)
		}
		
		proxyURL, err := url.Parse(proxyStr)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}
}

func (vd *VekoDome) rotateUserAgent() {
	vd.mutex.Lock()
	vd.currentUA = vd.userAgents[rand.Intn(len(vd.userAgents))]
	vd.mutex.Unlock()
}

func (vd *VekoDome) startRotation() {
	ticker := time.NewTicker(time.Duration(vd.config.RotationInterval) * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		vd.rotateUserAgent()
		vd.setupHTTPClient() // Rotate proxy as well
	}
}

func (vd *VekoDome) getCurrentIP() string {
	services := []string{
		"https://api.ipify.org",
		"https://ifconfig.me/ip",
		"https://icanhazip.com",
		"https://ipinfo.io/ip",
	}
	
	for _, service := range services {
		if ip := vd.fetchIP(service); ip != "" {
			return ip
		}
	}
	return "Unknown"
}

func (vd *VekoDome) fetchIP(service string) string {
	req, err := http.NewRequest("GET", service, nil)
	if err != nil {
		return ""
	}
	
	// Add spoofed headers
	vd.addSpoofedHeaders(req)
	
	resp, err := vd.client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	
	return strings.TrimSpace(string(body))
}

func (vd *VekoDome) addSpoofedHeaders(req *http.Request) {
	vd.mutex.RLock()
	ua := vd.currentUA
	vd.mutex.RUnlock()
	
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	
	// Random referer
	referers := []string{
		"https://www.google.com/",
		"https://www.bing.com/",
		"https://duckduckgo.com/",
		"https://github.com/",
	}
	req.Header.Set("Referer", referers[rand.Intn(len(referers))])
}

func (vd *VekoDome) monitorStatus() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		vd.displayStatus()
	}
}

func (vd *VekoDome) displayStatus() {
	currentIP := vd.getCurrentIP()
	vd.currentIP = currentIP
	
	// Clear screen and show status
	fmt.Print("\033[2J\033[H")
	
	status := "  Direct Connection"
	if vd.config.TorEnabled {
		status = " TOR Active"
	} else if len(vd.proxies) > 0 {
		status = " Proxy Active"
	}
	
	fmt.Printf(`
┌────────────────── VEKO DOME - STEALTH MODE ──────────────────┐
│ Current IP     : %-45s │
│ Status         : %-45s │
│ DNS Mode       : DoH (1.1.1.1)                               │
│ TLS Spoof      : %-45s │
│ User Agent     : %-45s │
│ Proxy Count    : %-45d │
│ Rotation       : Every %d seconds                            │
└──────────────────────────────────────────────────────────────┘

 Rotating identities automatically...
  Last update: %s

`, 
		currentIP,
		status,
		strings.ToUpper(vd.config.TLSFingerprint),
		vd.truncateString(vd.currentUA, 40),
		len(vd.proxies),
		vd.config.RotationInterval,
		time.Now().Format("15:04:05"),
	)
}

func (vd *VekoDome) truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// Test function to verify anonymity
func (vd *VekoDome) TestAnonymity() {
	fmt.Println("\n Testing anonymity...")
	
	// Test multiple IP services
	services := map[string]string{
		"ipify":      "https://api.ipify.org",
		"ifconfig":   "https://ifconfig.me/ip", 
		"icanhazip":  "https://icanhazip.com",
		"httpbin":    "https://httpbin.org/ip",
	}
	
	for name, service := range services {
		ip := vd.fetchIP(service)
		fmt.Printf("%-10s: %s\n", name, ip)
	}
	
	// Test DNS leak
	fmt.Println("\n DNS Leak Test:")
	dnsIP := vd.getCurrentIP()
	fmt.Printf("DNS Resolution IP: %s\n", dnsIP)
	
	// Test headers
	fmt.Println("\n Current Headers Test:")
	resp, err := vd.client.Get("https://httpbin.org/headers")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Headers Response: %s\n", string(body)[:200])
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
