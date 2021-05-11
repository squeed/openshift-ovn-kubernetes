// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

import (
	"encoding/json"

	"github.com/ovn-org/libovsdb/model"
	"github.com/ovn-org/libovsdb/ovsdb"
)

// FullDatabaseModel returns the DatabaseModel object to be used in libovsdb
func FullDatabaseModel() (*model.DBModel, error) {
	return model.NewDBModel("OVN_Northbound", map[string]model.Model{
		"ACL":                         &ACL{},
		"Address_Set":                 &AddressSet{},
		"Connection":                  &Connection{},
		"DHCP_Options":                &DHCPOptions{},
		"DNS":                         &DNS{},
		"Forwarding_Group":            &ForwardingGroup{},
		"Gateway_Chassis":             &GatewayChassis{},
		"HA_Chassis":                  &HAChassis{},
		"HA_Chassis_Group":            &HAChassisGroup{},
		"Load_Balancer":               &LoadBalancer{},
		"Load_Balancer_Health_Check":  &LoadBalancerHealthCheck{},
		"Logical_Router":              &LogicalRouter{},
		"Logical_Router_Policy":       &LogicalRouterPolicy{},
		"Logical_Router_Port":         &LogicalRouterPort{},
		"Logical_Router_Static_Route": &LogicalRouterStaticRoute{},
		"Logical_Switch":              &LogicalSwitch{},
		"Logical_Switch_Port":         &LogicalSwitchPort{},
		"Meter":                       &Meter{},
		"Meter_Band":                  &MeterBand{},
		"NAT":                         &NAT{},
		"NB_Global":                   &NBGlobal{},
		"Port_Group":                  &PortGroup{},
		"QoS":                         &QoS{},
		"SSL":                         &SSL{},
	})
}

var schema = `{
  "name": "OVN_Northbound",
  "version": "5.28.0",
  "tables": {
    "ACL": {
      "columns": {
        "action": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "allow",
                  "allow-related",
                  "drop",
                  "reject"
                ]
              ]
            }
          }
        },
        "direction": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "from-lport",
                  "to-lport"
                ]
              ]
            }
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "log": {
          "type": "boolean"
        },
        "match": {
          "type": "string"
        },
        "meter": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        },
        "name": {
          "type": {
            "key": {
              "type": "string",
              "minLength": 63,
              "maxLength": 63
            },
            "min": 0,
            "max": 1
          }
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32767
            }
          }
        },
        "severity": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "alert",
                  "warning",
                  "notice",
                  "info",
                  "debug"
                ]
              ]
            },
            "min": 0,
            "max": 1
          }
        }
      }
    },
    "Address_Set": {
      "columns": {
        "addresses": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Connection": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "inactivity_probe": {
          "type": {
            "key": {
              "type": "integer"
            },
            "min": 0,
            "max": 1
          }
        },
        "is_connected": {
          "type": "boolean",
          "ephemeral": true
        },
        "max_backoff": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1000
            },
            "min": 0,
            "max": 1
          }
        },
        "other_config": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "status": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          },
          "ephemeral": true
        },
        "target": {
          "type": "string"
        }
      },
      "indexes": [
        [
          "target"
        ]
      ]
    },
    "DHCP_Options": {
      "columns": {
        "cidr": {
          "type": "string"
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "DNS": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "records": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "Forwarding_Group": {
      "columns": {
        "child_port": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 1,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "liveness": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "vip": {
          "type": "string"
        },
        "vmac": {
          "type": "string"
        }
      }
    },
    "Gateway_Chassis": {
      "columns": {
        "chassis_name": {
          "type": "string"
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32767
            }
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "HA_Chassis": {
      "columns": {
        "chassis_name": {
          "type": "string"
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32767
            }
          }
        }
      }
    },
    "HA_Chassis_Group": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ha_chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "HA_Chassis",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Load_Balancer": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "health_check": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Load_Balancer_Health_Check",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ip_port_mappings": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "protocol": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "tcp",
                  "udp",
                  "sctp"
                ]
              ]
            },
            "min": 0,
            "max": 1
          }
        },
        "selection_fields": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "eth_src",
                  "eth_dst",
                  "ip_src",
                  "ip_dst",
                  "tp_src",
                  "tp_dst"
                ]
              ]
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "vips": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "Load_Balancer_Health_Check": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "vip": {
          "type": "string"
        }
      }
    },
    "Logical_Router": {
      "columns": {
        "enabled": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0,
            "max": 1
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "load_balancer": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Load_Balancer",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "nat": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "NAT",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "policies": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Logical_Router_Policy",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ports": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Logical_Router_Port",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "static_routes": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Logical_Router_Static_Route",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "Logical_Router_Policy": {
      "columns": {
        "action": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "allow",
                  "drop",
                  "reroute"
                ]
              ]
            }
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "match": {
          "type": "string"
        },
        "nexthop": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32767
            }
          }
        }
      }
    },
    "Logical_Router_Port": {
      "columns": {
        "enabled": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0,
            "max": 1
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "gateway_chassis": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Gateway_Chassis",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ha_chassis_group": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "HA_Chassis_Group",
              "refType": "strong"
            },
            "min": 0,
            "max": 1
          }
        },
        "ipv6_prefix": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ipv6_ra_configs": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "mac": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "networks": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 1,
            "max": "unlimited"
          }
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "peer": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Logical_Router_Static_Route": {
      "columns": {
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ip_prefix": {
          "type": "string"
        },
        "nexthop": {
          "type": "string"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "output_port": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        },
        "policy": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "src-ip",
                  "dst-ip"
                ]
              ]
            },
            "min": 0,
            "max": 1
          }
        }
      }
    },
    "Logical_Switch": {
      "columns": {
        "acls": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "ACL",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "dns_records": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "DNS",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "forwarding_groups": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Forwarding_Group",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "load_balancer": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Load_Balancer",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "other_config": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ports": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Logical_Switch_Port",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "qos_rules": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "QoS",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      }
    },
    "Logical_Switch_Port": {
      "columns": {
        "addresses": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "dhcpv4_options": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "DHCP_Options",
              "refType": "weak"
            },
            "min": 0,
            "max": 1
          }
        },
        "dhcpv6_options": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "DHCP_Options",
              "refType": "weak"
            },
            "min": 0,
            "max": 1
          }
        },
        "dynamic_addresses": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        },
        "enabled": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0,
            "max": 1
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "ha_chassis_group": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "HA_Chassis_Group",
              "refType": "strong"
            },
            "min": 0,
            "max": 1
          }
        },
        "name": {
          "type": "string"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "parent_name": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        },
        "port_security": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "tag": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 4095
            },
            "min": 0,
            "max": 1
          }
        },
        "tag_request": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 4095
            },
            "min": 0,
            "max": 1
          }
        },
        "type": {
          "type": "string"
        },
        "up": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0,
            "max": 1
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Meter": {
      "columns": {
        "bands": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Meter_Band",
              "refType": "strong"
            },
            "min": 1,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "fair": {
          "type": {
            "key": {
              "type": "boolean"
            },
            "min": 0,
            "max": 1
          }
        },
        "name": {
          "type": "string"
        },
        "unit": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "kbps",
                  "pktps"
                ]
              ]
            }
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "Meter_Band": {
      "columns": {
        "action": {
          "type": {
            "key": {
              "type": "string",
              "enum": "drop"
            }
          }
        },
        "burst_size": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 4294967295
            }
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "rate": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 4294967295
            }
          }
        }
      }
    },
    "NAT": {
      "columns": {
        "allowed_ext_ips": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Address_Set",
              "refType": "strong"
            },
            "min": 0,
            "max": 1
          }
        },
        "exempted_ext_ips": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Address_Set",
              "refType": "strong"
            },
            "min": 0,
            "max": 1
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "external_ip": {
          "type": "string"
        },
        "external_mac": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        },
        "external_port_range": {
          "type": "string"
        },
        "logical_ip": {
          "type": "string"
        },
        "logical_port": {
          "type": {
            "key": {
              "type": "string"
            },
            "min": 0,
            "max": 1
          }
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "type": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "dnat",
                  "snat",
                  "dnat_and_snat"
                ]
              ]
            }
          }
        }
      }
    },
    "NB_Global": {
      "columns": {
        "connections": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Connection"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "hv_cfg": {
          "type": "integer"
        },
        "hv_cfg_timestamp": {
          "type": "integer"
        },
        "ipsec": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "nb_cfg": {
          "type": "integer"
        },
        "nb_cfg_timestamp": {
          "type": "integer"
        },
        "options": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "sb_cfg": {
          "type": "integer"
        },
        "sb_cfg_timestamp": {
          "type": "integer"
        },
        "ssl": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "SSL"
            },
            "min": 0,
            "max": 1
          }
        }
      }
    },
    "Port_Group": {
      "columns": {
        "acls": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "ACL",
              "refType": "strong"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "name": {
          "type": "string"
        },
        "ports": {
          "type": {
            "key": {
              "type": "uuid",
              "refTable": "Logical_Switch_Port",
              "refType": "weak"
            },
            "min": 0,
            "max": "unlimited"
          }
        }
      },
      "indexes": [
        [
          "name"
        ]
      ]
    },
    "QoS": {
      "columns": {
        "action": {
          "type": {
            "key": {
              "type": "string",
              "enum": "dscp"
            },
            "value": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 63
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "bandwidth": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "rate",
                  "burst"
                ]
              ]
            },
            "value": {
              "type": "integer",
              "minInteger": 1,
              "maxInteger": 4294967295
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "direction": {
          "type": {
            "key": {
              "type": "string",
              "enum": [
                "set",
                [
                  "from-lport",
                  "to-lport"
                ]
              ]
            }
          }
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "match": {
          "type": "string"
        },
        "priority": {
          "type": {
            "key": {
              "type": "integer",
              "minInteger": 0,
              "maxInteger": 32767
            }
          }
        }
      }
    },
    "SSL": {
      "columns": {
        "bootstrap_ca_cert": {
          "type": "boolean"
        },
        "ca_cert": {
          "type": "string"
        },
        "certificate": {
          "type": "string"
        },
        "external_ids": {
          "type": {
            "key": {
              "type": "string"
            },
            "value": {
              "type": "string"
            },
            "min": 0,
            "max": "unlimited"
          }
        },
        "private_key": {
          "type": "string"
        },
        "ssl_ciphers": {
          "type": "string"
        },
        "ssl_protocols": {
          "type": "string"
        }
      }
    }
  }
}`

func Schema() ovsdb.DatabaseSchema {
	var s ovsdb.DatabaseSchema
	err := json.Unmarshal([]byte(schema), &s)
	if err != nil {
		panic(err)
	}
	return s
}
