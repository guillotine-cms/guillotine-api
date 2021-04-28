# Guillotine CMS is a headless CMS

Guillotine CMS places these pillars at the heart of its design

* High performance, Scalable design
* Security by default
* Multilingual support and search
 
## Cloud Agnostic

You can host Guillotine CMS on AWS, Azure and Google Cloud along with other platforms

[Hosting Guillotine on AWS](/deployments/terraform-aws)

[Hosting Guillotine on Azure](/deployments/terraform-azure)

[Hosting Guillotine on Google](/deployments/terraform-google)

[Hosting Guillotine on Kubernetes](/deployments/terraform-google)

## User Interface Agnostic

The platform will work with developers to allow content to be easily integrated into

To date we have integrated with 

* Angular
* Ionic
* ReactJS
* Flutter
  
In-game marketing, new bulletins and rich content

* Unity
* Unreal Engine

Integrated systems

* Raspirian
* Andorid OS

## API server design

### Language

The API server is written using Golang

https://golang.org/

> Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

### Persistent storage

The content is stored using ElasticSearch

https://www.elastic.co/elasticsearch/

> Elasticsearch is a distributed, RESTful search and analytics engine capable of addressing a growing number of use cases. As the heart of the Elastic Stack, it centrally stores your data for lightning fast search, fine‑tuned relevancy, and powerful analytics that scale with ease

### Caching

The application is written with care and full attention to the RFC2616 to correctly use headers such as Modified-Since, 
Etags etc... therefore is easily cachable in web cache proxies, we highly advocate using a content distribution network such as:


* AWS Cloudfront
* Google Cloud CDN
* Azure Content Delivery Network

They are globally distributed, have a formidable network backing and greatly excel at distributing content beyond what 
would be possible with a Redis, Memcache or other  

## Content Management Client design

The front end application is design in Ionic framework which under the hood uses Angular.

## FAQ's

Q: Do you plan to support other databases

A: No, whilst traditional SQL databases can be used to store content they do not offer the same search and language analysis
tools that power some of Guillotines most powerful features, we could, but in doing so we would nerf our offering.

Q: Why Golang

A: Authors perogative, I needed a highly efficient language both in coding style and execution Golang won in my books.