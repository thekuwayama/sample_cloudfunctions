# sample_cloudfunctions

[![Actions Status](https://github.com/thekuwayama/sample_cloudfunctions/workflows/CI/badge.svg)](https://github.com/thekuwayama/sample_cloudfunctions/actions?workflow=CI)
[![MIT licensed](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://raw.githubusercontent.com/thekuwayama/sample_cloudfunctions/master/LICENSE.txt)

## Usage

```bash
$ gcloud config configurations list | grep True | awk '{print $4}'
MYPROJECT

$ gcloud functions deploy listofname \
    --trigger-http \
    --allow-unauthenticated \
    --runtime go113 \
    --entry-point ListOfName \
    --region asia-northeast1
```

```bash
$ curl "https://asia-northeast1-${MYPROJECT}.cloudfunctions.net/listofname" -d '[{"name": "Bob"},{"name":"Alice"}]' | jq '.'
[
  {
    "age": 20,
    "sex": "male"
  },
  {
    "age": 10,
    "sex": "female"
  }
]
```

## License

The CLI is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
