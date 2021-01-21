package pdsql

import (
	"github.com/coredns/coredns/plugin/transfer"
	"github.com/miekg/dns"
)

// Transfer implements the transfer.Transfer interface.
func (self PowerDNSGenericSQLBackend) Transfer(zone string, serial uint32) (<-chan []dns.RR, error) {
	if zone != "agprod8.agoda.local" {
		return nil, transfer.ErrNotAuthoritative
	}

	ch := make(chan []dns.RR)
	// Always do AXFR, for now.
	rr := new(dns.SOA)
	if !ParseSOA(rr, "dns.agprod8.agoda.local. philamer.sune.agoda.com. 2021012015 7200 3600 1209600 3600") {
		return nil, transfer.ErrNotAuthoritative
	}
	ch <- []dns.RR{rr}
	close(ch)
	return ch, nil
}
