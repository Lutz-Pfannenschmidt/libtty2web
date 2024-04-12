package tty2web

type option struct {
	name     string
	value    string
	hasValue bool
}

// IP address to listen (default: "0.0.0.0")
func WithAddress(address string) option {
	return option{name: "address", value: address, hasValue: true}
}

// Port number to listen (default: "8080")
func WithPort(port string) option {
	return option{name: "port", value: port, hasValue: true}
}

// Permit clients to write to the TTY (BE CAREFUL) (default: false)
func WithPermitWrite() option {
	return option{name: "permit-write", hasValue: false}
}

// Credential for Basic Authentication (ex: user:pass, default disabled)
func WithCredential(credential string) option {
	return option{name: "credential", value: credential, hasValue: true}
}

// Add a random string to the URL (default: false)
func WithRandomUrl() option {
	return option{name: "random-url", hasValue: false}
}

// Enable WebGL renderer (default: false)
func WithWebGL() option {
	return option{name: "enable-webgl", hasValue: false}
}

// Turn on all features: download /, upload /, api, regeorg, ... (default: false)
func WithAllFeatures() option {
	return option{name: "all", hasValue: false}
}

// Enable API for executing commands on the system (BE CAREFUL!) (default: false)
func WithApi() option {
	return option{name: "api", hasValue: false}
}

// Enable API for executing sc on the system (BE CAREFUL!) (default: false)
func WithSc() option {
	return option{name: "sc", hasValue: false}
}

// Enable socks4/socks5 proxy using regeorg (default: false)
func WithRegeorg() option {
	return option{name: "regeorg", hasValue: false}
}

// Random URL length (default: 8)
func WithRandomUrlLength(length string) option {
	return option{name: "random-url-length", value: length, hasValue: true}
}

// Specify string for the URL
func WithUrl(url string) option {
	return option{name: "url", value: url, hasValue: true}
}

// Specify string for custom URL serving Javascript files (useful for DNS)
func WithJsUrl(url string) option {
	return option{name: "jsurl", value: url, hasValue: true}
}

// Serve files to download from specified dir
func WithDownloadDir(dir string) option {
	return option{name: "download", value: dir, hasValue: true}
}

// Enable uploading of files to the specified dir (BE CAREFUL!)
func WithUploadDir(dir string) option {
	return option{name: "upload", value: dir, hasValue: true}
}

// Enable TLS/SSL (default: false)
func WithTls() option {
	return option{name: "tls", hasValue: false}
}

// TLS/SSL certificate file path (default: "~/.tty2web.crt")
func WithTlsCert(cert string) option {
	return option{name: "tls-cert", value: cert, hasValue: true}
}

// TLS/SSL key file path (default: "~/.tty2web.key")
func WithTlsKey(key string) option {
	return option{name: "tls-key", value: key, hasValue: true}
}

// TLS/SSL CA certificate file for client certifications (default: "~/.tty2web.ca.crt")
func WithTlsCa(ca string) option {
	return option{name: "tls-ca-crt", value: ca, hasValue: true}
}

// Custom index.html file
func WithIndexHtml(file string) option {
	return option{name: "index", value: file, hasValue: true}
}

// Title format of browser window (default: "{{ .command }}@{{ .hostname }}")
func WithTitleFormat(format string) option {
	return option{name: "title-format", value: format, hasValue: true}
}

// Use domain for DNS tunneling (ex. example.com)
func WithDns(domain string) option {
	return option{name: "dns", value: domain, hasValue: true}
}

// Listen for reverse connection agents (ex. 0.0.0.0:53)
func WithDnsListen(address string) option {
	return option{name: "dnslisten", value: address, hasValue: true}
}

// Password/Key to use for DNS tunnel
func WithDnsKey(key string) option {
	return option{name: "dnskey", value: key, hasValue: true}
}

// Delay time between polling for DNS requests (default: "200ms")
func WithDnsDelay(delay string) option {
	return option{name: "dnsdelay", value: delay, hasValue: true}
}

// Listen for reverse connection agents (ex. 0.0.0.0:4444)
func WithListen(address string) option {
	return option{name: "listen", value: address, hasValue: true}
}

// Enable TLS for listening for agents and clients itself (default: false)
func WithListenTls() option {
	return option{name: "agenttls", hasValue: false}
}

// Certificate and key for listen server (ex. mycert)
func WithListenCert(cert string) option {
	return option{name: "listencert", value: cert, hasValue: true}
}

// Server for forwarding reverse connections (ex. 127.0.0.1:6000) (default: "127.0.0.1:6000")
func WithForwardServer(address string) option {
	return option{name: "forward", value: address, hasValue: true}
}

// Password for reverse server connection
func WithPassword(key string) option {
	return option{name: "password", value: key, hasValue: true}
}

// Connect to host for reverse connection (ex. 192.168.1.1:4444)
func WithConnect(address string) option {
	return option{name: "connect", value: address, hasValue: true}
}

// Use proxy for reverse server connection (ex. 192.168.1.1:8080)
func WithProxy(address string) option {
	return option{name: "proxy", value: address, hasValue: true}
}

// Use proxy authentication for reverse server connection (ex. DOMAIN/user:password)
func WithProxyAuth(auth string) option {
	return option{name: "proxyauth", value: auth, hasValue: true}
}

// Use user agent for reverse server connection (ex. Mozilla)
func WithUserAgent(agent string) option {
	return option{name: "useragent", value: agent, hasValue: true}
}

// Enable reconnection (default: false)
func WithReconnect() option {
	return option{name: "reconnect", hasValue: false}
}

// Enable verbose messages (default: false)
func WithVerbose() option {
	return option{name: "verbose", hasValue: false}
}

// Time to reconnect (default: 10)
func WithReconnectTime(time string) option {
	return option{name: "reconnect-time", value: time, hasValue: true}
}

// Maximum connection to tty2web (default: 0)
func WithMaxConnection(max string) option {
	return option{name: "max-connection", value: max, hasValue: true}
}

// Accept only one client and exit on disconnection (default: false)
func WithOnce() option {
	return option{name: "once", hasValue: false}
}

// Timeout seconds for waiting a client(0 to disable) (default: 0)
func WithTimeout(timeout string) option {
	return option{name: "timeout", value: timeout, hasValue: true}
}

// Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB) (default: false)
func WithPermitArgs() option {
	return option{name: "permit-arguments", hasValue: false}
}

// Static width of the screen, 0(default) means dynamically resize (default: 0)
func WithWidth(width string) option {
	return option{name: "width", value: width, hasValue: true}
}

// Static height of the screen, 0(default) means dynamically resize (default: 0)
func WithHeight(height string) option {
	return option{name: "height", value: height, hasValue: true}
}

// A regular expression that matches origin URLs to be accepted by WebSocket. No cross origin requests are acceptable by default
func WithWsOrigin(origin string) option {
	return option{name: "ws-origin", value: origin, hasValue: true}
}

// Terminal name to use on the browser, one of xterm or hterm. (default: "xterm")
func WithTerminal(terminal string) option {
	return option{name: "terminal", value: terminal, hasValue: true}
}

// Signal sent to the command process when tty2webclose it (default: SIGHUP) (default: 1)
func WithCloseSignal(signal string) option {
	return option{name: "close-signal", value: signal, hasValue: true}
}

// Time in seconds to force kill process after client is disconnected (default: -1)
func WithCloseTimeout(time string) option {
	return option{name: "close-timeout", value: time, hasValue: true}
}

// Config file path (default: "~/.tty2web")
func WithConfig(file string) option {
	return option{name: "config", value: file, hasValue: true}
}
