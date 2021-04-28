resource "aws_elasticsearch_domain" "example" {
  domain_name           = "guillotine-cms"
  elasticsearch_version = "7.10"

  cluster_config {
    instance_type = "t3.small.elasticsearch"
  }
}