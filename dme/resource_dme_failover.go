package dme

import (
	"fmt"
	"log"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDMEFailover() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMEFailoverCreate,
		Read:   resourceDMEFailoverRead,
		Update: resourceDMEFailoverUpdate,
		Delete: resourceDMEFailoverDelete,

		Schema: map[string]*schema.Schema{
			"record_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"system_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"max_emails": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"protocol_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"auto_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"contact_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"send_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"dns_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"dns_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_query_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDMEFailoverCreate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	failoverAttr := &models.FailoverAttribute{}

	if monitor, ok := d.GetOk("monitor"); ok {
		failoverAttr.Monitor = monitor.(string)
	}

	if system_description, ok := d.GetOk("system_description"); ok {
		failoverAttr.SystemDescription = system_description.(string)
	}

	if max_emails, ok := d.GetOk("max_emails"); ok {
		failoverAttr.MaxEmails = max_emails.(string)
	}

	if sensitivity, ok := d.GetOk("sensitivity"); ok {
		failoverAttr.Sensitivity = sensitivity.(string)
	}

	if protocol_id, ok := d.GetOk("protocol_id"); ok {
		failoverAttr.ProtocolId = protocol_id.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		failoverAttr.Port = port.(string)
	}

	if failover, ok := d.GetOk("failover"); ok {
		failoverAttr.Failover = failover.(string)
	}

	if auto_failover, ok := d.GetOk("auto_failover"); ok {
		failoverAttr.AutoFailover = auto_failover.(string)
	}

	if ip1, ok := d.GetOk("ip1"); ok {
		failoverAttr.Ip1 = ip1.(string)
	}

	if ip2, ok := d.GetOk("ip2"); ok {
		failoverAttr.Ip2 = ip2.(string)
	}

	if ip3, ok := d.GetOk("ip3"); ok {
		failoverAttr.Ip3 = ip3.(string)
	}

	if ip4, ok := d.GetOk("ip4"); ok {
		failoverAttr.Ip4 = ip4.(string)
	}

	if ip5, ok := d.GetOk("ip5"); ok {
		failoverAttr.Ip5 = ip5.(string)
	}

	if contact_list, ok := d.GetOk("contact_list"); ok {
		failoverAttr.ContactList = contact_list.(string)
	}

	if http_fqdn, ok := d.GetOk("http_fqdn"); ok {
		failoverAttr.HttpFqdn = http_fqdn.(string)
	}

	if http_file, ok := d.GetOk("http_file"); ok {
		failoverAttr.HttpFile = http_file.(string)
	}

	if http_query_string, ok := d.GetOk("http_query_string"); ok {
		failoverAttr.HttpQueryString = http_query_string.(string)
	}

	if send_string, ok := d.GetOk("send_string"); ok {
		failoverAttr.SendString = send_string.(string)
	}

	if timeout, ok := d.GetOk("timeout"); ok {
		failoverAttr.Timeout = timeout.(string)
	}

	if dns_fqdn, ok := d.GetOk("dns_fqdn"); ok {
		failoverAttr.DNSFqdn = dns_fqdn.(string)
	}

	if dns_timeout, ok := d.GetOk("dns_timeout"); ok {
		failoverAttr.DNSTimeout = dns_timeout.(string)
	}

	log.Println("Failover insside create structure is :", failoverAttr)
	_, err := dmeClient.Update(failoverAttr, "monitor/"+d.Get("record_id").(string))
	if err != nil {
		return err
	}
	// log.Println("value of container in create:", con)
	// log.Println("Output containier create domain :", con.S("recordId"))
	// d.SetId(fmt.Sprintf("%v", con.S("recordId")))
	return resourceDMEFailoverRead(d, m)
}

func resourceDMEFailoverRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)

	con, err := dmeClient.GetbyId("monitor/" + d.Get("record_id").(string))
	if err != nil {
		return err
	}
	log.Println("Inside read, container value: ", con)

	d.SetId(fmt.Sprintf("%v", con.S("recordId")))
	d.Set("monitor", StripQuotes(con.S("monitor").String()))
	d.Set("system_description", StripQuotes(con.S("systemDescription").String()))
	d.Set("max_emails", StripQuotes(con.S("maxEmails").String()))
	d.Set("sensitivity", StripQuotes(con.S("sensitivity").String()))
	d.Set("protocol_id", StripQuotes(con.S("protocolId").String()))
	d.Set("port", StripQuotes(con.S("port").String()))
	d.Set("failover", StripQuotes(con.S("failover").String()))
	d.Set("auto_failover", StripQuotes(con.S("autoFailover").String()))
	d.Set("ip1", StripQuotes(con.S("ip1").String()))
	d.Set("ip2", StripQuotes(con.S("ip2").String()))
	d.Set("ip3", StripQuotes(con.S("ip3").String()))
	d.Set("ip4", StripQuotes(con.S("ip4").String()))
	d.Set("ip5", StripQuotes(con.S("ip5").String()))
	d.Set("contact_list", StripQuotes(con.S("contactListId").String()))
	d.Set("http_fqdn", StripQuotes(con.S("httpFqdn").String()))
	d.Set("http_file", StripQuotes(con.S("httpFile").String()))
	d.Set("http_query_string", StripQuotes(con.S("httpQueryString").String()))
	d.Set("send_string", StripQuotes(con.S("sendString").String()))
	d.Set("timeout", StripQuotes(con.S("timeout").String()))
	// d.Set("dns_fqdn", StripQuotes(con.S("dnsFqdn").String()))
	// d.Set("dns_timeout", StripQuotes(con.S("dnsTimeout").String()))
	return nil
}

func resourceDMEFailoverUpdate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	failoverAttr := &models.FailoverAttribute{}

	if monitor, ok := d.GetOk("monitor"); ok {
		failoverAttr.Monitor = monitor.(string)
	}

	if system_description, ok := d.GetOk("system_description"); ok {
		failoverAttr.SystemDescription = system_description.(string)
	}

	if max_emails, ok := d.GetOk("max_emails"); ok {
		failoverAttr.MaxEmails = max_emails.(string)
	}

	if sensitivity, ok := d.GetOk("sensitivity"); ok {
		failoverAttr.Sensitivity = sensitivity.(string)
	}

	if protocol_id, ok := d.GetOk("protocol_id"); ok {
		failoverAttr.ProtocolId = protocol_id.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		failoverAttr.Port = port.(string)
	}

	if failover, ok := d.GetOk("failover"); ok {
		failoverAttr.Failover = failover.(string)
	}

	if auto_failover, ok := d.GetOk("auto_failover"); ok {
		failoverAttr.AutoFailover = auto_failover.(string)
	}

	if ip1, ok := d.GetOk("ip1"); ok {
		failoverAttr.Ip1 = ip1.(string)
	}

	if ip2, ok := d.GetOk("ip2"); ok {
		failoverAttr.Ip2 = ip2.(string)
	}

	if ip3, ok := d.GetOk("ip3"); ok {
		failoverAttr.Ip3 = ip3.(string)
	}

	if ip4, ok := d.GetOk("ip4"); ok {
		failoverAttr.Ip4 = ip4.(string)
	}

	if ip5, ok := d.GetOk("ip5"); ok {
		failoverAttr.Ip5 = ip5.(string)
	}

	if contact_list, ok := d.GetOk("contact_list"); ok {
		failoverAttr.ContactList = contact_list.(string)
	}

	if http_fqdn, ok := d.GetOk("http_fqdn"); ok {
		failoverAttr.HttpFqdn = http_fqdn.(string)
	}

	if http_file, ok := d.GetOk("http_file"); ok {
		failoverAttr.HttpFile = http_file.(string)
	}

	if http_query_string, ok := d.GetOk("http_query_string"); ok {
		failoverAttr.HttpQueryString = http_query_string.(string)
	}

	if send_string, ok := d.GetOk("send_string"); ok {
		failoverAttr.SendString = send_string.(string)
	}

	if timeout, ok := d.GetOk("timeout"); ok {
		failoverAttr.Timeout = timeout.(string)
	}

	if dns_fqdn, ok := d.GetOk("dns_fqdn"); ok {
		failoverAttr.DNSFqdn = dns_fqdn.(string)
	}

	if dns_timeout, ok := d.GetOk("dns_timeout"); ok {
		failoverAttr.DNSTimeout = dns_timeout.(string)
	}

	log.Println("Failover insside create structure is :", failoverAttr)
	_, err := dmeClient.Update(failoverAttr, "monitor/"+d.Get("record_id").(string))
	if err != nil {
		return err
	}
	return resourceDMEFailoverRead(d, m)
}
func resourceDMEFailoverDelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	failoverAttr := &models.FailoverAttribute{}
	failoverAttr.Failover = "false"

	if port, ok := d.GetOk("port"); ok {
		failoverAttr.Port = port.(string)
	}

	if sensitivity, ok := d.GetOk("sensitivity"); ok {
		failoverAttr.Sensitivity = sensitivity.(string)
	}

	log.Println("inside delete", failoverAttr)

	_, err := dmeClient.Update(failoverAttr, "monitor/"+d.Get("record_id").(string))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
