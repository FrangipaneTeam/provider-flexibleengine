package vpc

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vpc_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_v1
	p.AddResourceConfigurator("flexibleengine_vpc_v1", func(r *config.Resource) {
		r.Kind = "Vpc"
	})

	// flexibleengine_vpc_subnet_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_subnet_v1
	p.AddResourceConfigurator("flexibleengine_vpc_subnet_v1", func(r *config.Resource) {
		r.Kind = "Subnet"

		// vpc_id is the ID of the VPC
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}
	})

	// flexibleengine_vpc_route_table
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route_table
	p.AddResourceConfigurator("flexibleengine_vpc_route_table", func(r *config.Resource) {
		r.Kind = "RouteTable"

		// vpc_id is the ID of the VPC
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}

		// Subnets is a list of Subnet IDs
		r.References["subnets"] = config.Reference{
			Type: "Subnet",
		}
	})

	// flexibleengine_vpc_route
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_route
	p.AddResourceConfigurator("flexibleengine_vpc_route", func(r *config.Resource) {
		r.Kind = "Route"

		// route_table_id is the ID of the Route Table
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}

	})

	// flexibleengine_vpc_peering_connection_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_peering_v2
	p.AddResourceConfigurator("flexibleengine_vpc_peering_connection_v2", func(r *config.Resource) {
		r.Kind = "Peering"

		// vpc_id is the ID of the VPC
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}

	})

	// flexibleengine_vpc_peering_connection_accepter_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_peering_accepter_v2
	p.AddResourceConfigurator("flexibleengine_vpc_peering_connection_accepter_v2", func(r *config.Resource) {
		r.Kind = "PeeringAccepter"

		// vpc_id is the ID of the VPC
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}

		// peering_id is the ID of the VPC Peering Connection
		r.References["peering_id"] = config.Reference{
			Type: "Peering",
		}

	})

	// TODO This require LTS (Log Tank Service)
	// flexibleengine_vpc_flow_log_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_flow_log_v1
	// p.AddResourceConfigurator("flexibleengine_vpc_flow_log_v1", func(r *config.Resource) {
	// 	r.Kind = "FlowLog"

	// 	// TODO Check if this is correct
	// 	r.References["ressource_id"] = config.Reference{
	// 		Type: "NetworkPort",
	// 	}

	// 	r.References["log_group_id"] = config.Reference{
	// 		Type: "LogGroup",
	// 	}

	// })

	// flexibleengine_networking_network_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_network_v2
	p.AddResourceConfigurator("flexibleengine_networking_network_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "Network"
	})

	// flexibleengine_networking_router_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_router_v2
	p.AddResourceConfigurator("flexibleengine_networking_router_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "Router"
	})

	// flexibleengine_networking_router_interface_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_router_interface_v2
	p.AddResourceConfigurator("flexibleengine_networking_router_interface_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "RouterInterface"

		// router_id is the ID of the router to add the interface to.
		r.References["router_id"] = config.Reference{
			Type: "Router",
		}

		// subnet_id is the ID of the subnet to add to the router.
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}

		// port_id is the ID of the port to add to the router.
		r.References["port_id"] = config.Reference{
			Type: "NetworkPort",
		}
	})

	// flexibleengine_networking_port_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_port_v2
	p.AddResourceConfigurator("flexibleengine_networking_port_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "NetworkPort"

		// network_id is the ID of the network to which this port will be attached.
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}
	})

	// flexibleengine_networking_secgroup_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_v2
	p.AddResourceConfigurator("flexibleengine_networking_secgroup_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "SecurityGroup"
	})

	// flexibleengine_networking_secgroup_rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_secgroup_rule_v2
	p.AddResourceConfigurator("flexibleengine_networking_secgroup_rule_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "SecurityGroupRule"

		// security_group_id is the ID of the security group to add the rule to.
		r.References["security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}

		// remote_group_id is the ID of the security group to which this rule applies.
		r.References["remote_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	// flexibleengine_networking_subnet_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_subnet_v2
	p.AddResourceConfigurator("flexibleengine_networking_subnet_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "NetworkSubnet"

		// network_id is the ID of the network to which this subnet will be attached.
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}
	})

	// flexibleengine_networking_vip_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_vip_v2
	p.AddResourceConfigurator("flexibleengine_networking_vip_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "Vip"

		// network_id is the ID of the network to which this subnet will be attached.
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}

		// subnet_id is the ID of the subnet to which this VIP will be attached.
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
	})

	// flexibleengine_networking_vip_associate_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/networking_vip_associate_v2
	p.AddResourceConfigurator("flexibleengine_networking_vip_associate_v2", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "VipAssociate"

		// vip_id is the ID of the VIP to associate.
		r.References["vip_id"] = config.Reference{
			Type: "Vip",
		}

		// port_id is the ID of the port to associate.
		r.References["port_ids"] = config.Reference{
			Type: "NetworkPort",
		}
	})

	// flexibleengine_vpc_eip
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_eip
	// * This resource not required configuration

	// flexibleengine_vpc_eip_associate
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_eip_associate
	p.AddResourceConfigurator("flexibleengine_vpc_eip_associate", func(r *config.Resource) {
		r.Kind = "EIPAssociate"

		// eip_id is the ID of the EIP to associate.
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}

		// port_id is the ID of the port to associate.
		r.References["port_id"] = config.Reference{
			Type: "NetworkPort",
		}
	})

	// flexibleengine_vpcep_service
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_service
	p.AddResourceConfigurator("flexibleengine_vpcep_service", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "EndpointService"

		// vpc_id is the ID of the VPC to which this endpoint service will be attached.
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}

		// port_id is the ID of the port to associate.
		r.References["port_id"] = config.Reference{
			Type: "NetworkPort",
		}

	})

	// flexibleengine_vpcep_endpoint
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_endpoint
	p.AddResourceConfigurator("flexibleengine_vpcep_endpoint", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "Endpoint"

		// vpc_id is the ID of the VPC to which this endpoint will be attached.
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}

		// network_id is the ID of the network to which this endpoint will be attached.
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}

		// service_id is the ID of the service to which this endpoint will be attached.
		r.References["service_id"] = config.Reference{
			Type: "EndpointService",
		}

	})

	// flexibleengine_vpcep_approval
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_approval
	p.AddResourceConfigurator("flexibleengine_vpcep_approval", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "EndpointApproval"

		// endpoint_id is the ID of the endpoint to which this approval will be attached.
		r.References["endpoint_id"] = config.Reference{
			Type: "EndPoint",
		}

		// service_id is the ID of the service to which this approval will be attached.
		r.References["service_id"] = config.Reference{
			Type: "EndpointService",
		}

	})

}