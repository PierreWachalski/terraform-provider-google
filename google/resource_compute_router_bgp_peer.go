// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/googleapi"
)

func resourceComputeRouterBgpPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRouterBgpPeerCreate,
		Read:   resourceComputeRouterBgpPeerRead,
		Update: resourceComputeRouterBgpPeerUpdate,
		Delete: resourceComputeRouterBgpPeerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRouterBgpPeerImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"interface": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the interface the BGP peer is associated with.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRFC1035Name(2, 63),
				Description: `Name of this BGP peer. The name must be 1-63 characters long,
and comply with RFC1035. Specifically, the name must be 1-63 characters
long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"peer_asn": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				Description: `Peer BGP Autonomous System Number (ASN).
Each BGP interface may use a different value.`,
			},
			"peer_ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `IP address of the BGP interface outside Google Cloud Platform.
Only IPv4 is supported.`,
			},
			"router": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the Cloud Router in which this BgpPeer will be configured.`,
			},
			"advertise_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DEFAULT", "CUSTOM", ""}, false),
				Description: `User-specified flag to indicate which mode to use for advertisement.
Valid values of this enum field are: 'DEFAULT', 'CUSTOM'`,
				Default: "DEFAULT",
			},
			"advertised_groups": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `User-specified list of prefix groups to advertise in custom
mode, which can take one of the following options:

* 'ALL_SUBNETS': Advertises all available subnets, including peer VPC subnets.
* 'ALL_VPC_SUBNETS': Advertises the router's own VPC subnets.
* 'ALL_PEER_VPC_SUBNETS': Advertises peer subnets of the router's VPC network.


Note that this field can only be populated if advertiseMode is 'CUSTOM'
and overrides the list defined for the router (in the "bgp" message).
These groups are advertised in addition to any specified prefixes.
Leave this field blank to advertise no custom groups.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"advertised_ip_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `User-specified list of individual IP ranges to advertise in
custom mode. This field can only be populated if advertiseMode
is 'CUSTOM' and is advertised to all peers of the router. These IP
ranges will be advertised in addition to any specified groups.
Leave this field blank to advertise no custom IP ranges.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The IP range to advertise. The value must be a
CIDR-formatted string.`,
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `User-specified description for the IP range.`,
						},
					},
				},
			},
			"advertised_route_priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The priority of routes advertised to this BGP peer.
Where there is more than one matching route of maximum
length, the routes with the lowest priority value win.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `Region where the router and BgpPeer reside.
If it is not provided, the provider region is used.`,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `IP address of the interface inside Google Cloud Platform.
Only IPv4 is supported.`,
			},
			"management_type": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource that configures and manages this BGP peer.

* 'MANAGED_BY_USER' is the default value and can be managed by
you or other users
* 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and
managed by Cloud Interconnect, specifically by an
InterconnectAttachment of type PARTNER. Google automatically
creates, updates, and deletes this type of BGP peer when the
PARTNER InterconnectAttachment is created, updated,
or deleted.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceComputeRouterBgpPeerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRouterBgpPeerName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	interfaceNameProp, err := expandComputeRouterBgpPeerInterface(d.Get("interface"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interface"); !isEmptyValue(reflect.ValueOf(interfaceNameProp)) && (ok || !reflect.DeepEqual(v, interfaceNameProp)) {
		obj["interfaceName"] = interfaceNameProp
	}
	peerIpAddressProp, err := expandComputeRouterBgpPeerPeerIpAddress(d.Get("peer_ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peer_ip_address"); !isEmptyValue(reflect.ValueOf(peerIpAddressProp)) && (ok || !reflect.DeepEqual(v, peerIpAddressProp)) {
		obj["peerIpAddress"] = peerIpAddressProp
	}
	peerAsnProp, err := expandComputeRouterBgpPeerPeerAsn(d.Get("peer_asn"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peer_asn"); !isEmptyValue(reflect.ValueOf(peerAsnProp)) && (ok || !reflect.DeepEqual(v, peerAsnProp)) {
		obj["peerAsn"] = peerAsnProp
	}
	advertisedRoutePriorityProp, err := expandComputeRouterBgpPeerAdvertisedRoutePriority(d.Get("advertised_route_priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("advertised_route_priority"); !isEmptyValue(reflect.ValueOf(advertisedRoutePriorityProp)) && (ok || !reflect.DeepEqual(v, advertisedRoutePriorityProp)) {
		obj["advertisedRoutePriority"] = advertisedRoutePriorityProp
	}
	advertiseModeProp, err := expandComputeRouterBgpPeerAdvertiseMode(d.Get("advertise_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("advertise_mode"); !isEmptyValue(reflect.ValueOf(advertiseModeProp)) && (ok || !reflect.DeepEqual(v, advertiseModeProp)) {
		obj["advertiseMode"] = advertiseModeProp
	}
	advertisedGroupsProp, err := expandComputeRouterBgpPeerAdvertisedGroups(d.Get("advertised_groups"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("advertised_groups"); ok || !reflect.DeepEqual(v, advertisedGroupsProp) {
		obj["advertisedGroups"] = advertisedGroupsProp
	}
	advertisedIpRangesProp, err := expandComputeRouterBgpPeerAdvertisedIpRanges(d.Get("advertised_ip_ranges"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("advertised_ip_ranges"); ok || !reflect.DeepEqual(v, advertisedIpRangesProp) {
		obj["advertisedIpRanges"] = advertisedIpRangesProp
	}

	lockName, err := replaceVars(d, config, "router/{{region}}/{{router}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RouterBgpPeer: %#v", obj)

	obj, err = resourceComputeRouterBgpPeerPatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpPeer: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating RouterBgpPeer",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RouterBgpPeer: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RouterBgpPeer %q: %#v", d.Id(), res)

	return resourceComputeRouterBgpPeerRead(d, meta)
}

func resourceComputeRouterBgpPeerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRouterBgpPeer %q", d.Id()))
	}

	res, err = flattenNestedComputeRouterBgpPeer(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeRouterBgpPeer because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}

	if err := d.Set("name", flattenComputeRouterBgpPeerName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("interface", flattenComputeRouterBgpPeerInterface(res["interfaceName"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("ip_address", flattenComputeRouterBgpPeerIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("peer_ip_address", flattenComputeRouterBgpPeerPeerIpAddress(res["peerIpAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("peer_asn", flattenComputeRouterBgpPeerPeerAsn(res["peerAsn"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("advertised_route_priority", flattenComputeRouterBgpPeerAdvertisedRoutePriority(res["advertisedRoutePriority"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("advertise_mode", flattenComputeRouterBgpPeerAdvertiseMode(res["advertiseMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("advertised_groups", flattenComputeRouterBgpPeerAdvertisedGroups(res["advertisedGroups"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("advertised_ip_ranges", flattenComputeRouterBgpPeerAdvertisedIpRanges(res["advertisedIpRanges"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}
	if err := d.Set("management_type", flattenComputeRouterBgpPeerManagementType(res["managementType"], d, config)); err != nil {
		return fmt.Errorf("Error reading RouterBgpPeer: %s", err)
	}

	return nil
}

func resourceComputeRouterBgpPeerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	advertisedRoutePriorityProp, err := expandComputeRouterBgpPeerAdvertisedRoutePriority(d.Get("advertised_route_priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("advertised_route_priority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, advertisedRoutePriorityProp)) {
		obj["advertisedRoutePriority"] = advertisedRoutePriorityProp
	}
	advertisedIpRangesProp, err := expandComputeRouterBgpPeerAdvertisedIpRanges(d.Get("advertised_ip_ranges"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("advertised_ip_ranges"); ok || !reflect.DeepEqual(v, advertisedIpRangesProp) {
		obj["advertisedIpRanges"] = advertisedIpRangesProp
	}

	lockName, err := replaceVars(d, config, "router/{{region}}/{{router}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RouterBgpPeer %q: %#v", d.Id(), obj)

	obj, err = resourceComputeRouterBgpPeerPatchUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RouterBgpPeer %q: %s", d.Id(), err)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating RouterBgpPeer",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeRouterBgpPeerRead(d, meta)
}

func resourceComputeRouterBgpPeerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "router/{{region}}/{{router}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceComputeRouterBgpPeerPatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return handleNotFoundError(err, d, "RouterBgpPeer")
	}
	log.Printf("[DEBUG] Deleting RouterBgpPeer %q", d.Id())

	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RouterBgpPeer")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting RouterBgpPeer",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RouterBgpPeer %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRouterBgpPeerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/routers/(?P<router>[^/]+)/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<router>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<router>[^/]+)/(?P<name>[^/]+)",
		"(?P<router>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRouterBgpPeerName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerInterface(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerPeerIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerPeerAsn(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeRouterBgpPeerAdvertisedRoutePriority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeRouterBgpPeerAdvertiseMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerAdvertisedGroups(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerAdvertisedIpRanges(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"range":       flattenComputeRouterBgpPeerAdvertisedIpRangesRange(original["range"], d, config),
			"description": flattenComputeRouterBgpPeerAdvertisedIpRangesDescription(original["description"], d, config),
		})
	}
	return transformed
}
func flattenComputeRouterBgpPeerAdvertisedIpRangesRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerAdvertisedIpRangesDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouterBgpPeerManagementType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeRouterBgpPeerName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerPeerIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerPeerAsn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerAdvertisedRoutePriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerAdvertiseMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerAdvertisedGroups(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerAdvertisedIpRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRange, err := expandComputeRouterBgpPeerAdvertisedIpRangesRange(original["range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRange); val.IsValid() && !isEmptyValue(val) {
			transformed["range"] = transformedRange
		}

		transformedDescription, err := expandComputeRouterBgpPeerAdvertisedIpRangesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["description"] = transformedDescription
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRouterBgpPeerAdvertisedIpRangesRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpPeerAdvertisedIpRangesDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func flattenNestedComputeRouterBgpPeer(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["bgpPeers"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value bgpPeers. Actual value: %v", v)
	}

	_, item, err := resourceComputeRouterBgpPeerFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeRouterBgpPeerFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName, err := expandComputeRouterBgpPeerName(d.Get("name"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemName := flattenComputeRouterBgpPeerName(item["name"], d, meta.(*Config))
		if !reflect.DeepEqual(itemName, expectedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceComputeRouterBgpPeerPatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceComputeRouterBgpPeerListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceComputeRouterBgpPeerFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create RouterBgpPeer, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"bgpPeers": append(currItems, obj),
	}

	return res, nil
}

// PatchUpdateEncoder handles creating request data to PATCH parent resource
// with list including updated object.
func resourceComputeRouterBgpPeerPatchUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	items, err := resourceComputeRouterBgpPeerListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceComputeRouterBgpPeerFindNestedObjectInList(d, meta, items)
	if err != nil {
		return nil, err
	}

	// Return error if item to update does not exist.
	if item == nil {
		return nil, fmt.Errorf("Unable to update RouterBgpPeer %q - not found in list", d.Id())
	}

	// Merge new object into old.
	for k, v := range obj {
		item[k] = v
	}
	items[idx] = item

	// Return list with new item added
	res := map[string]interface{}{
		"bgpPeers": items,
	}

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceComputeRouterBgpPeerPatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceComputeRouterBgpPeerListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceComputeRouterBgpPeerFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, &googleapi.Error{
			Code:    404,
			Message: "RouterBgpPeer not found in list",
		}
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"bgpPeers": updatedItems,
	}

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceComputeRouterBgpPeerListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*Config)
	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}")
	if err != nil {
		return nil, err
	}
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool

	v, ok = res["bgpPeers"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "bgpPeers"`)
		}
		return ls, nil
	}
	return nil, nil
}
