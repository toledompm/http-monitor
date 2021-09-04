package trace

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"time"
)

type TraceResult struct {
	Response              *http.Response
	DnsTime               time.Duration
	TlsTime               time.Duration
	ConnectionTime        time.Duration
	FirstResponseByteTime time.Duration
	TotalTime             time.Duration
	RequestError          error
}

// Traces the request and returns the result
func TraceGetRequest(url string) (*TraceResult, error) {
	timer := generateHttpRequestTimer()
	totalTimer := newTimer()

	request, _ := http.NewRequest("GET", url, nil)
	context := httptrace.WithClientTrace(request.Context(), timer.ClientTrace)

	request = request.WithContext(context)

	totalTimer.Start()
	timer.FirstResponseByteTimer.Start()
	response, err := http.DefaultTransport.RoundTrip(request)
	totalTimer.Stop()

	if err != nil {
		return nil, err
	}

	return &TraceResult{
		Response:              response,
		DnsTime:               timer.DnsTimer.Measure(),
		TlsTime:               timer.TlsTimer.Measure(),
		ConnectionTime:        timer.ConnectionTimer.Measure(),
		FirstResponseByteTime: timer.FirstResponseByteTimer.Measure(),
		TotalTime:             totalTimer.Measure(),
	}, nil
}

type requestTimer struct {
	DnsTimer               *timer
	TlsTimer               *timer
	ConnectionTimer        *timer
	FirstResponseByteTimer *timer
	ClientTrace            *httptrace.ClientTrace
	RequestError           error
}

func generateHttpRequestTimer() *requestTimer {
	dnsTimer := newTimer()
	tlsTimer := newTimer()
	connectionTimer := newTimer()
	firstResponseByteTimer := newTimer()

	clientTrace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) {
			dnsTimer.Start()
		},
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			dnsTimer.Stop()
		},

		ConnectStart: func(network, addr string) {
			connectionTimer.Start()
		},
		ConnectDone: func(network, addr string, err error) {
			connectionTimer.Stop()
		},

		TLSHandshakeStart: func() {
			tlsTimer.Start()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			tlsTimer.Stop()
		},

		GotFirstResponseByte: func() {
			firstResponseByteTimer.Stop()
		},
	}

	return &requestTimer{
		DnsTimer:               dnsTimer,
		TlsTimer:               tlsTimer,
		ConnectionTimer:        connectionTimer,
		FirstResponseByteTimer: firstResponseByteTimer,
		ClientTrace:            clientTrace,
	}
}
