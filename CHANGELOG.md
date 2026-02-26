# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.0](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/compare/v0.3.0...v0.4.0) (2026-02-26)


### Features

* add optimistic locking tests for computer and mobile device prestages ([dc99830](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/dc998301f10779e7fc91a2fab5d3ffd66f6f6dcb))
* add optimistic locking tests for computer and mobile device prestages ([2c42cec](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/2c42cec3ba5cf0b77930fe9cbae941d99f93c38b))
* add unit tests for various API services and enhance mock responders ([39aafe6](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/39aafe6de27d551095c3ba402a7de818c6bf08fb))
* add unit tests for various API services and enhance mock responders ([f640371](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/f640371898c988ae6c7aec2c3fe1de69a977fa20))
* enhance unit tests for API services with error handling and validation ([a41c71a](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/a41c71a22dc7bf2488fc83f457f5d4c0201e7f76))
* implement UUID preservation in macOS configuration profile updates ([9b2de85](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/9b2de85cb1047d1806e834d3911a602c14c6e353))
* implement UUID preservation in macOS configuration profile updates ([ec879c5](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/ec879c5d8e4fc36c83c6fd5e67ee772e0bb29e42))

## [0.3.0](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/compare/v0.2.0...v0.3.0) (2026-02-25)


### Features

* feat:  ([65eba71](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/65eba7111ba2aba053472af17a25de24fe61b698))
* enhance API client with new services and refactor structure ([b1c1e61](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/b1c1e616c9164ce55c3964bc11616411e36655d2))
* enhance ClassicCommandFlush service with additional status handling ([a66b525](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/a66b5256e2e253db5e34e653feef1553bdcf9d21))
* expand computer inventory API documentation and models ([f004ae8](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/f004ae89f1362af854a2c5a2b46f7f57e8512897))
* update account preferences API to v3 and enhance examples ([74231ad](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/74231ad98c7c7b47c4d97d218b75dc1582440a9b))


### Bug Fixes

* standardize endpoint usage across services ([51a4c27](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/51a4c27f1163d553c609bb6733d53d8dc5787b0e))

## [0.2.0](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/compare/v0.1.0...v0.2.0) (2026-02-23)


### Features

* feat:  ([1e28505](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/1e2850586360115921ac8ba0868921226e179de2))
* add ClassicCommandFlush service to API client ([0c0380d](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/0c0380d616bd6e179f0635c8d7da27d7e4660f54))
* enhance onboarding API with history and notes functionality ([3fdd241](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/3fdd2412938ed171b94baeea83ee46579a339af9))
* update .gitignore and enhance jamfpro API client ([7d9c09f](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/7d9c09fa01f13d41b2c440b5e7bf67531d1a02b7))

## 0.1.0 (2026-02-20)


### Features

* add ComputerPrestages service to Jamf Pro client ([11c3d99](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/11c3d999270853981a8ec3af215fe8c4d4252a0f))
* add GreaterThanJamfProVersion function for version-specific test skipping ([bd79754](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/bd79754a38520901b1d5003d0a0ffa9ee9036fc6))
* add new Jamf Pro API services for buildings, computer groups, dock items, and packages; update go.mod and go.sum for new dependencies ([f57338f](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/f57338feef6b289d900d31526efa770b3e77465f))
* add new services and methods for Jamf Pro API, including smart and static computer groups, enhancing the client structure and API interactions ([f89f1d4](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/f89f1d4b89add58289c745f266040ac9e4f78c44))
* added all scaffolding and first service jamf_pro_api categories ([b424d51](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/b424d516a946e9eb0f981ecd587dba219a8aa213))
* added departments ([1b563a7](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/1b563a7fe9d5ce3b284178f594f575c4753e278f))
* enhance SDK documentation with timeouts, retries, and TLS configuration ([a68136e](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/a68136e614daf3b7da4c7e4b9e4d5536f3705e6f))
* enhance weekly test workflow with unit tests and coverage reporting ([851b7ac](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/851b7ac19d74efc5f14c6bde41b9438476c72fe3))
* integrate Classic API services and refactor headers in categories and scripts services ([3420fbf](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/3420fbf903482047dca0884408293a8004d7fd6a))
* update README and error handling in Jamf Pro client ([a531e34](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/a531e34069aaf5e15cbaf3c24b27d1e1e10f011d))


### Bug Fixes

* enhance login customization test with default disclaimer and action text ([8c3468a](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/8c3468a5ce11fac6e7a8ff237500f0912aac2ccb))
* update FillUserTemplate to false for non-.dmg packages and adjust dock item type constants to uppercase ([560a4bb](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/commit/560a4bb9f4e642869ddd44a667c32196c055a9d9))

## [Unreleased]

### Added

- Added xyz [@your_username](https://github.com/your_username)

### Fixed

- Fixed zyx [@your_username](https://github.com/your_username)

## [1.1.0] - 2021-06-23

### Added

- Added x [@your_username](https://github.com/your_username)

### Changed

- Changed y [@your_username](https://github.com/your_username)

## [1.0.0] - 2021-06-20

### Added

- Inititated y [@your_username](https://github.com/your_username)
- Inititated z [@your_username](https://github.com/your_username)
