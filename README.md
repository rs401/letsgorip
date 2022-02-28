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
    Lets Go Rip is a project to show potential employers my capabilities. The project is written in Go(lang) and has microservices for Auth, Forums, Places and an HTTP API. I used this project to learn more about gRPC and Kubernetes. I used Angular for the front end.
    <br />
    <a href="https://github.com/rs401/letsgorip"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/rs401/letsgorip">View Demo</a>
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

This project is a work in progress that does not yet have a live demo.

Lets Go Rip is a project to show potential employers my capabilities. The project is written in Go(lang) and has microservices for Auth, Forums, Places and an HTTP API. I used this project to learn more about gRPC and Kubernetes. I used Angular for the front end.

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
- [ ] UI
    - [x] Models
    - [x] Auth components
    - [ ] Forum components
    - [ ] Thread components
    - [ ] Post components
    - [ ] Places components

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