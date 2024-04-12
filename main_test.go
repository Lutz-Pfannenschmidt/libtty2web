package tty2web_test

import (
	"testing"

	tty2web "github.com/Lutz-Pfannenschmidt/libtty2web"
)

func TestNewTty2Web(t *testing.T) {
	tty2web := tty2web.NewTty2Web("some-command")
	if tty2web == nil {
		t.Error("Expected a new tty2web instance, got nil")
	}
}

func TestCommandBuilder(t *testing.T) {
	test := tty2web.NewTty2Web("some-command")
	test.SetBinary("custom-binary")
	test.AddOptions(
		tty2web.WithAddress("1.2.3.4"),
		tty2web.WithPort("8080"),
		tty2web.WithPermitWrite(),
		tty2web.WithCredential("user:pass"),
		tty2web.WithRandomUrl(),
		tty2web.WithWebGL(),
		tty2web.WithAllFeatures(),
		tty2web.WithApi(),
		tty2web.WithSc(),
		tty2web.WithRegeorg(),
		tty2web.WithRandomUrlLength("16"),
		tty2web.WithUrl("url"),
		tty2web.WithJsUrl("jsurl"),
		tty2web.WithDownloadDir("download"),
		tty2web.WithUploadDir("upload"),
		tty2web.WithTls(),
		tty2web.WithTlsCert("cert"),
		tty2web.WithTlsKey("key"),
		tty2web.WithTlsCa("ca"),
		tty2web.WithIndexHtml("index"),
		tty2web.WithTitleFormat("format"),
		tty2web.WithDns("dns"),
		tty2web.WithDnsListen("dnslisten"),
		tty2web.WithListen("listen"),
		tty2web.WithListenTls(),
		tty2web.WithListenCert("cert"),
		tty2web.WithForwardServer("server"),
		tty2web.WithPassword("password"),
		tty2web.WithConnect("connect"),
		tty2web.WithProxy("proxy"),
		tty2web.WithProxyAuth("proxyauth"),
		tty2web.WithReconnectTime("10"),
		tty2web.WithMaxConnection("0"),
		tty2web.WithOnce(),
		tty2web.WithTimeout("10"),
		tty2web.WithPermitArgs(),
		tty2web.WithWidth("800"),
		tty2web.WithHeight("600"),
		tty2web.WithWsOrigin("origin"),
		tty2web.WithTerminal("terminal"),
		tty2web.WithCloseSignal("signal"),
		tty2web.WithCloseTimeout("10"),
		tty2web.WithConfig("config"),
	)

	cmd := test.GetCommand()
	expected := "custom-binary --address 1.2.3.4 --port 8080 --permit-write --credential user:pass --random-url --enable-webgl --all --api --sc --regeorg --random-url-length 16 --url url --jsurl jsurl --download download --upload upload --tls --tls-cert cert --tls-key key --tls-ca-crt ca --index index --title-format format --dns dns --dnslisten dnslisten --listen listen --agenttls --listencert cert --forward server --password password --connect connect --proxy proxy --proxyauth proxyauth --reconnect-time 10 --max-connection 0 --once --timeout 10 --permit-arguments --width 800 --height 600 --ws-origin origin --terminal terminal --close-signal signal --close-timeout 10 --config config some-command"
	if cmd != expected {
		t.Errorf("Expected command: %s, got: %s", expected, cmd)
	}
}
