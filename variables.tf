# variables.tf
variable "ACCESS_KEY" {
 default = "Enter_Access_KEY"
}
variable "SECRET_KEY" {
 default = "Enter_Secret_KEY"
}
variable "region" {
 default = "us-east-1"
}
variable "availabilityZone" {
        default = "us-east-1a"
}
variable "instanceTenancy" {
 default = "default"
}
variable "dnsSupport" {
 default = true
}
variable "dnsHostNames" {
        default = true
}
variable "vpcCIDRblock" {
 default = "10.0.0.0/16"
}
variable "subnetCIDRblock" {
        default = "10.0.1.0/24"
}
variable "destinationCIDRblock" {
        default = "0.0.0.0/0"
}
variable "ingressCIDRblock" {
        type = "list"
        default = [ "0.0.0.0/0" ]
}
variable "mapPublicIP" {
        default = true
}
variable "ami" {
        default = "ami-cd0f5cb6"
}
variable "instance_type" {
        default = "t2.small"
}
# end of variables.tf