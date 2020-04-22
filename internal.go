package gomasscan

// thanks 003random
// https://gist.github.com/003random/7f678dd8f50f87bc830db47a976bc4de
import (
	"net"
)

var (
	ranges []string = []string{
		"0.0.0.0/32",         // Current network (only valid as source address)
		"240.0.0.0/4",        // Reserved for future use
		"203.0.113.0/24",     // Assigned as TEST-NET-3
		"198.51.100.0/24",    // Assigned as TEST-NET-2, documentation and examples
		"198.18.0.0/15",      // Used for benchmark testing of inter-network communications between two separate subnets
		"192.0.2.0/24",       // Assigned as TEST-NET-1, documentation and examples
		"100.64.0.0/10",      // Shared address space for communications between a service provider and its subscribers when using a carrier-grade NAT.
		"255.255.255.255/32", // Reserved for the "limited broadcast" destination address
		"192.0.0.0/24",       // IETF Protocol Assignments
		"192.0.2.0/24",       // Assigned as TEST-NET-1, documentation and examples
		"192.88.99.0/24",     // Reserved. Formerly used for IPv6 to IPv4 relay (included IPv6 address block 2002::/16)
		"192.168.0.0/16",     // Used for local communications within a private network
		"172.16.0.0/12",      // Used for local communications within a private network
		"10.0.0.0/8",         // Used for local communications within a private network
		"127.0.0.0/8",        // Used for loopback addresses to the local host
		"169.254.0.0/16",     // Used for link-local addresses between two hosts on a single link when no IP address is otherwise specified
		"224.0.0.0/4",        // In use for IP multicast.[9] (Former Class D network)
	}
	blacklist []net.IPNet = createBlacklist()
)

func createBlacklist() []net.IPNet {
	b := []net.IPNet{}
	for _, sCIDR := range ranges {
		_, c, _ := net.ParseCIDR(sCIDR)
		b = append(b, *c)
	}
	return b
}

func IsInternal(sIP string) bool {
	for _, r := range blacklist {
		if r.Contains(net.ParseIP(sIP)) {
			return true
		}
	}
	return false
}
