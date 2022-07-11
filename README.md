<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<div align="center">

[![Build][actions-build-shield]][actions-build-url]
[![Tests][actions-tests-shield]][actions-tests-url]
[![Go Report Card](https://goreportcard.com/badge/github.com/rs401/letsgorip)](https://goreportcard.com/report/github.com/rs401/letsgorip)
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![BSD 3-Clause License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

</div>


<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">Lets Go Rip</h3>

  <p align="center">
    Lets Go Rip is an application hosting a community to help like minded motorsports enthusiasts, find places to ride/drive/drift/jump/race and people to go with. 
    <br />
    <a href="https://github.com/rs401/letsgorip"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://www.letsgo.rip">View Demo</a>
    ·
    <a href="https://github.com/rs401/letsgorip/issues">Report Bug</a>
    ·
    <a href="https://github.com/rs401/letsgorip/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><a href="#usage">Example Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project is a work in progress ~~that does not yet have a live demo~~ that can be seen at [https://www.letsgo.rip](https://www.letsgo.rip).

Lets Go Rip is a portfolio project. The project is written in Go(lang) and has microservices for Auth, Forums, Places and an HTTP API. I used this project to learn more about gRPC and Kubernetes. I use PostgreSQL for my data store and Angular for the front end. The app is currently running in a cluster on Google Kubernetes Engine (GKE). 

## The Journey

After finishing school and learning 4 languages in the process (C++, C#, Java, Python), I felt the need to pick a language and start a great portfolio project. Looking around a little in different communities and job postings, I decided to learn Go and build microservice's. I wanted to learn more about Kubernetes so I needed the services to be deployed in easily replicable containers for scaling and be ephemeral for easy CD. I also wanted to learn about gRPC because some article or click baity title somewhere, painted gRPC as a "Hot Topic", I also love learning so adding another concept to the list seemed like a good idea.

### Challenges

First challenge was finding out what gRPC is and how it works. I found that gRPC is an RPC framework that uses protocol buffers to define how data is to be structured. You define your data structures as "Messages", because your data will be transmitted as a message. You define your service with RPC functions that use messages as parameters and return types. Once messages and services are defined, the protocol buffer compiler (protoc) can be used to generate code in the language of your choosing. The generate code defines service and client interfaces to be used to transmit the messages. The messages will be serialized into a compact binary format and transmitted.

OK, so I planned out my data structures and basic CRUD methods for services in .proto files and compiled them. Next challenge was implementing the services servers. This consisted of creating the data structures as models (so I can add gorm tags as I will be using gorm as my ORM), creating repository interfaces to utilize gorm for database storage and create a service structure that implements the protocol buffers NameServiceServer interface. The structure that implements the ServiceServer interface receives a repository, the service methods call the repository methods to interact with the data store. The service methods also convert the data into the protocol buffers data structures. Once the service methods are sorted out, all that's left is to create a grpc.NewServer(), register the service with the server and listen on my chosen port.

After my services were all set, I created an HTTP API that has separate handlers for each service, each handler implementing the given services NameServiceClient interface.

To be continued...




<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Go](https://go.dev/)
* [gRPC](https://grpc.io/)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://www.docker.com/)
* [Kubernetes](https://kubernetes.io/)
* [mini-Kube](https://github.com/kubernetes/minikube)
* [Angular](https://angular.io/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To run this application you will need to install the prerequisites and use the different Make targets.

### Prerequisites

* Go
* Docker
* MiniKube
* Make


<!-- USAGE EXAMPLES -->
## Usage

1. Install Prerequisites
2. Clone the repo
   ```sh
   git clone https://github.com/rs401/letsgorip.git
   ```
3. Start MiniKube and set local docker repo
   ```sh
   minikube start
   eval $(minikube -p minikube docker-env)
   ```
4. Build the container images
   ```sh
   make build-docker
   ```
5. Start the kubernetes deployments and services
   ```sh
   make kube
   ```
6. To stop the kubernetes deployments and services
   ```sh
   make kube-down
   ```


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Complete Protocol Buffers Messages
- [x] Generate Protocol Buffers Go files
- [x] Complete API
    - [x] Auth
        - [x] Handlers
        - [x] Routes
        - [x] Implement JWT 
        - [x] Middleware
    - [x] Forums
        - [x] Forum handlers
        - [x] Thread handlers
        - [x] Post handlers
        - [x] Routes
    - [x] Places
        - [x] Handlers
        - [x] Routes
- [x] Build docker image from API
- [x] Build docker image from auth
- [x] Build docker image from forums
- [x] Build docker image from places
- [x] Build k8s configs with label selectors and environment variables for communication
- [x] Build k8s configs for services
- [x] Rebuild images for k8s
- [x] UI
    - [x] Models
    - [x] Auth components
    - [x] Forum components
    - [x] Thread components
    - [x] Post components
    - [x] Places components
        - [x] Google maps
- [x] GKE
    - [x] Setup cluster
    - [x] Setup artifact registry
        - [x] Push images to registry
    - [x] Execute kubernetes configs
    - [x] Setup letsgo.rip domain with GKE DNS
    - [x] Setup ingress
    - [x] Setup SSL certificates with GKE
    - [x] Setup Persistent Volume for Postgres db
- [x] Add Google OAuth for identity

See the [open issues](https://github.com/rs401/letsgorip/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

> **WANTED** Code reviews: I am eager to be a better programmer.

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the BSD 3-Clause License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

[![LinkedIn][linkedin-shield]][linkedin-url]

Project Link: [https://github.com/rs401/letsgorip](https://github.com/rs401/letsgorip)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* []()
* []()
* []()

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[stars-shield]: https://img.shields.io/github/stars/rs401/letsgorip.svg?style=plastic
[stars-url]: https://github.com/rs401/letsgorip/stargazers
[issues-shield]: https://img.shields.io/github/issues/rs401/letsgorip.svg?style=plastic
[issues-url]: https://github.com/rs401/letsgorip/issues
[license-shield]: https://img.shields.io/github/license/rs401/letsgorip.svg?style=plastic
[license-url]: https://github.com/rs401/letsgorip/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=plastic&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/richard-stadnick-3b4ab53b
[product-screenshot]: images/screenshot.png
[actions-tests-shield]: https://github.com/rs401/letsgorip/actions/workflows/tests.yaml/badge.svg
[actions-tests-url]: https://github.com/rs401/letsgorip/actions/workflows/tests.yaml
[actions-build-shield]: https://github.com/rs401/letsgorip/actions/workflows/build.yaml/badge.svg
[actions-build-url]: https://github.com/rs401/letsgorip/actions/workflows/build.yaml
[stars-shield]: https://img.shields.io/github/stars/rs401/letsgorip.svg?style=plastic
[stars-url]: https://github.com/rs401/letsgorip./stargazers