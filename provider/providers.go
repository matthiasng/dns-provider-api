package provider

import (
	"fmt"
	"os"

	"github.com/matthiasng/dns-provider-api/provider/acmedns"
	"github.com/matthiasng/dns-provider-api/provider/alidns"
	"github.com/matthiasng/dns-provider-api/provider/auroradns"
	"github.com/matthiasng/dns-provider-api/provider/autodns"
	"github.com/matthiasng/dns-provider-api/provider/azure"
	"github.com/matthiasng/dns-provider-api/provider/bindman"
	"github.com/matthiasng/dns-provider-api/provider/bluecat"
	"github.com/matthiasng/dns-provider-api/provider/checkdomain"
	"github.com/matthiasng/dns-provider-api/provider/clouddns"
	"github.com/matthiasng/dns-provider-api/provider/cloudflare"
	"github.com/matthiasng/dns-provider-api/provider/cloudns"
	"github.com/matthiasng/dns-provider-api/provider/cloudxns"
	"github.com/matthiasng/dns-provider-api/provider/conoha"
	"github.com/matthiasng/dns-provider-api/provider/constellix"
	"github.com/matthiasng/dns-provider-api/provider/designate"
	"github.com/matthiasng/dns-provider-api/provider/digitalocean"
	"github.com/matthiasng/dns-provider-api/provider/dnsimple"
	"github.com/matthiasng/dns-provider-api/provider/dnsmadeeasy"
	"github.com/matthiasng/dns-provider-api/provider/dnspod"
	"github.com/matthiasng/dns-provider-api/provider/dode"
	"github.com/matthiasng/dns-provider-api/provider/dreamhost"
	"github.com/matthiasng/dns-provider-api/provider/duckdns"
	"github.com/matthiasng/dns-provider-api/provider/dyn"
	"github.com/matthiasng/dns-provider-api/provider/dynu"
	"github.com/matthiasng/dns-provider-api/provider/easydns"
	"github.com/matthiasng/dns-provider-api/provider/exoscale"
	"github.com/matthiasng/dns-provider-api/provider/fastdns"
	"github.com/matthiasng/dns-provider-api/provider/gandi"
	"github.com/matthiasng/dns-provider-api/provider/gandiv5"
	"github.com/matthiasng/dns-provider-api/provider/gcloud"
	"github.com/matthiasng/dns-provider-api/provider/glesys"
	"github.com/matthiasng/dns-provider-api/provider/godaddy"
	"github.com/matthiasng/dns-provider-api/provider/hostingde"
	"github.com/matthiasng/dns-provider-api/provider/iij"
	"github.com/matthiasng/dns-provider-api/provider/inwx"
	"github.com/matthiasng/dns-provider-api/provider/joker"
	"github.com/matthiasng/dns-provider-api/provider/lightsail"
	"github.com/matthiasng/dns-provider-api/provider/linode"
	"github.com/matthiasng/dns-provider-api/provider/linodev4"
	"github.com/matthiasng/dns-provider-api/provider/liquidweb"
	"github.com/matthiasng/dns-provider-api/provider/mydnsjp"
	"github.com/matthiasng/dns-provider-api/provider/mythicbeasts"
	"github.com/matthiasng/dns-provider-api/provider/namecheap"
	"github.com/matthiasng/dns-provider-api/provider/namedotcom"
	"github.com/matthiasng/dns-provider-api/provider/namesilo"
	"github.com/matthiasng/dns-provider-api/provider/netcup"
	"github.com/matthiasng/dns-provider-api/provider/nifcloud"
	"github.com/matthiasng/dns-provider-api/provider/ns1"
	"github.com/matthiasng/dns-provider-api/provider/oraclecloud"
	"github.com/matthiasng/dns-provider-api/provider/otc"
	"github.com/matthiasng/dns-provider-api/provider/ovh"
	"github.com/matthiasng/dns-provider-api/provider/pdns"
	"github.com/matthiasng/dns-provider-api/provider/rackspace"
	"github.com/matthiasng/dns-provider-api/provider/regru"
	"github.com/matthiasng/dns-provider-api/provider/rfc2136"
	"github.com/matthiasng/dns-provider-api/provider/rimuhosting"
	"github.com/matthiasng/dns-provider-api/provider/route53"
	"github.com/matthiasng/dns-provider-api/provider/sakuracloud"
	"github.com/matthiasng/dns-provider-api/provider/scaleway"
	"github.com/matthiasng/dns-provider-api/provider/selectel"
	"github.com/matthiasng/dns-provider-api/provider/servercow"
	"github.com/matthiasng/dns-provider-api/provider/stackpath"
	"github.com/matthiasng/dns-provider-api/provider/transip"
	"github.com/matthiasng/dns-provider-api/provider/vegadns"
	"github.com/matthiasng/dns-provider-api/provider/versio"
	"github.com/matthiasng/dns-provider-api/provider/vscale"
	"github.com/matthiasng/dns-provider-api/provider/vultr"
	"github.com/matthiasng/dns-provider-api/provider/yandex"
	"github.com/matthiasng/dns-provider-api/provider/zoneee"
	"github.com/matthiasng/dns-provider-api/provider/zonomi"
)

type Provider interface {
	Present(domain, token, fqdn, value string) error
	CleanUp(domain, token, fqdn, value string) error
}

func NewProvider(name string, vars map[string]string) (Provider, error) {
	for key, value := range vars {
		os.Setenv(key, value)
	}

	defer func(){
		for key := range vars {
			os.Unsetenv(key)
		}
	}()

	return NewDNSChallengeProviderByName(name)
}

// NewDNSChallengeProviderByName Factory for DNS providers
func NewDNSChallengeProviderByName(name string) (Provider, error) {
	switch name {
	case "acme-dns":
		return acmedns.NewDNSProvider()
	case "alidns":
		return alidns.NewDNSProvider()
	case "azure":
		return azure.NewDNSProvider()
	case "auroradns":
		return auroradns.NewDNSProvider()
	case "autodns":
		return autodns.NewDNSProvider()
	case "bindman":
		return bindman.NewDNSProvider()
	case "bluecat":
		return bluecat.NewDNSProvider()
	case "checkdomain":
		return checkdomain.NewDNSProvider()
	case "clouddns":
		return clouddns.NewDNSProvider()
	case "cloudflare":
		return cloudflare.NewDNSProvider()
	case "cloudns":
		return cloudns.NewDNSProvider()
	case "cloudxns":
		return cloudxns.NewDNSProvider()
	case "conoha":
		return conoha.NewDNSProvider()
	case "constellix":
		return constellix.NewDNSProvider()
	case "designate":
		return designate.NewDNSProvider()
	case "digitalocean":
		return digitalocean.NewDNSProvider()
	case "dnsimple":
		return dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		return dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		return dnspod.NewDNSProvider()
	case "dode":
		return dode.NewDNSProvider()
	case "dreamhost":
		return dreamhost.NewDNSProvider()
	case "duckdns":
		return duckdns.NewDNSProvider()
	case "dyn":
		return dyn.NewDNSProvider()
	case "dynu":
		return dynu.NewDNSProvider()
	case "fastdns":
		return fastdns.NewDNSProvider()
	case "easydns":
		return easydns.NewDNSProvider()
	// case "exec":
	// 	return exec.NewDNSProvider()
	case "exoscale":
		return exoscale.NewDNSProvider()
	case "gandi":
		return gandi.NewDNSProvider()
	case "gandiv5":
		return gandiv5.NewDNSProvider()
	case "glesys":
		return glesys.NewDNSProvider()
	case "gcloud":
		return gcloud.NewDNSProvider()
	case "godaddy":
		return godaddy.NewDNSProvider()
	case "hostingde":
		return hostingde.NewDNSProvider()
	// case "httpreq":
	// 	return httpreq.NewDNSProvider()
	case "iij":
		return iij.NewDNSProvider()
	case "inwx":
		return inwx.NewDNSProvider()
	case "joker":
		return joker.NewDNSProvider()
	case "lightsail":
		return lightsail.NewDNSProvider()
	case "linode":
		return linode.NewDNSProvider()
	case "linodev4":
		return linodev4.NewDNSProvider()
	case "liquidweb":
		return liquidweb.NewDNSProvider()
	// case "manual":
	// 	return dns01.NewDNSProviderManual()
	case "mydnsjp":
		return mydnsjp.NewDNSProvider()
	case "mythicbeasts":
		return mythicbeasts.NewDNSProvider()
	case "namecheap":
		return namecheap.NewDNSProvider()
	case "namedotcom":
		return namedotcom.NewDNSProvider()
	case "namesilo":
		return namesilo.NewDNSProvider()
	case "netcup":
		return netcup.NewDNSProvider()
	case "nifcloud":
		return nifcloud.NewDNSProvider()
	case "ns1":
		return ns1.NewDNSProvider()
	case "oraclecloud":
		return oraclecloud.NewDNSProvider()
	case "otc":
		return otc.NewDNSProvider()
	case "ovh":
		return ovh.NewDNSProvider()
	case "pdns":
		return pdns.NewDNSProvider()
	case "rackspace":
		return rackspace.NewDNSProvider()
	case "regru":
		return regru.NewDNSProvider()
	case "rfc2136":
		return rfc2136.NewDNSProvider()
	case "rimuhosting":
		return rimuhosting.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	case "sakuracloud":
		return sakuracloud.NewDNSProvider()
	case "scaleway":
		return scaleway.NewDNSProvider()
	case "selectel":
		return selectel.NewDNSProvider()
	case "servercow":
		return servercow.NewDNSProvider()
	case "stackpath":
		return stackpath.NewDNSProvider()
	case "transip":
		return transip.NewDNSProvider()
	case "vegadns":
		return vegadns.NewDNSProvider()
	case "versio":
		return versio.NewDNSProvider()
	case "vultr":
		return vultr.NewDNSProvider()
	case "vscale":
		return vscale.NewDNSProvider()
	case "yandex":
		return yandex.NewDNSProvider()
	case "zoneee":
		return zoneee.NewDNSProvider()
	case "zonomi":
		return zonomi.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
