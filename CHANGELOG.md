# Changelog

## 0.1.0-alpha.1 (2025-07-08)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/KarthikBoddeda/rzp-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([39ccd65](https://github.com/KarthikBoddeda/rzp-go/commit/39ccd65a78e687f0e33a6e2acab88e02e3b74433))
* **client:** add debug log helper ([46175ad](https://github.com/KarthikBoddeda/rzp-go/commit/46175ada657892d91f935869fac45d37c15ebcbd))
* **client:** add escape hatch for null slice & maps ([895d330](https://github.com/KarthikBoddeda/rzp-go/commit/895d3300898765939373454b566605cb9d59aa3b))
* **client:** add support for endpoint-specific base URLs in python ([b4727f0](https://github.com/KarthikBoddeda/rzp-go/commit/b4727f092a2fb04eca80b3c7a9c244cca75c8fd4))
* **client:** allow overriding unions ([2da22e9](https://github.com/KarthikBoddeda/rzp-go/commit/2da22e9a38371de02ac674e129947179c843074f))


### Bug Fixes

* **client:** correctly set stream key for multipart ([e542d8e](https://github.com/KarthikBoddeda/rzp-go/commit/e542d8e4dacc8f55bb7bc1680f25321a14b22fe8))
* **client:** don't panic on marshal with extra null field ([3608896](https://github.com/KarthikBoddeda/rzp-go/commit/3608896b340898fbbca32c61abe7a81146206e40))
* don't try to deserialize as json when ResponseBodyInto is []byte ([10b4d31](https://github.com/KarthikBoddeda/rzp-go/commit/10b4d31a2bf9d5ee56759fea9a100c5371a42a81))
* fix error ([924a224](https://github.com/KarthikBoddeda/rzp-go/commit/924a224ad75ab5dc35f6152b18ed65535f94c33c))


### Chores

* **ci:** enable for pull requests ([99fd52e](https://github.com/KarthikBoddeda/rzp-go/commit/99fd52e9728985e4dde96428eaf90d45259964cd))
* **ci:** only run for pushes and fork pull requests ([c86de9f](https://github.com/KarthikBoddeda/rzp-go/commit/c86de9fc347dd5d708bd7fc501dc2c4c6dfc916f))
* **docs:** grammar improvements ([fedbc36](https://github.com/KarthikBoddeda/rzp-go/commit/fedbc36a236571ae044c7de6091eeadc6bd10ff4))
* fix documentation of null map ([7cbf94e](https://github.com/KarthikBoddeda/rzp-go/commit/7cbf94e8ca10b81fd34ad1d2ed68bd534751d29e))
* improve devcontainer setup ([d930aca](https://github.com/KarthikBoddeda/rzp-go/commit/d930acadb36d18d696a6bbf5e81b5fafed4147fd))
* lint tests ([ec7e3c7](https://github.com/KarthikBoddeda/rzp-go/commit/ec7e3c7f1fe6dd341f9e0a2ebd5e5c8b5b31f9b8))
* make go mod tidy continue on error ([28d98d7](https://github.com/KarthikBoddeda/rzp-go/commit/28d98d7caf59a99ab33aeaaf234a424e06d61088))
* update SDK settings ([c9cca4a](https://github.com/KarthikBoddeda/rzp-go/commit/c9cca4ac233fb3ced77a8b062e830eae8f7fc7fb))
