# Run-code

http://www.51cto.com

run-code is fast and pretty used for online judge in Go.

## Environment variables
run-code takes its configuration from environment variables.

| Variable name            | Allowed values                     | Example                  | Description              |
|:-------------------------|:-----------------------------------|:-------------------------|:-------------------------|
| model                    | debug                              | debug                    | run model                |
| address                  | :8989                              | :8989                    | Listen port              |
| run_auth_token           | 123456                             | 123456                   | run auth token           |


## Quick Start

#### Before we get started
    1.Install the docker https://github.com/docker/docker-ce
    2.Pull docker images https://github.com/prasmussen/glot-containers

#### Create a run directory
    mkdir run-code
    cd run-code
    mkdir runtime
    mkdir conf
    cd conf
    vim app.json
    ```
    {
      "model": "debug",
      "address": ":8989",
      "run_auth_token": "123456"
    }
    ```

#### Download release
    run-code/tree/master/release/1.0/run-code

#### Run the backend
    ./run-code

#### Start your journey
    e.g: via http
    POST /run/run-code?run_auth_token=123456 HTTP/1.1
    Host: 127.0.0.1:8989
    Content-Type: application/json
    cache-control: no-cache
    Postman-Token: 0e08fa0c-b1f7-44b9-8e67-d74ec29ee971
    {
        "language": "php",
        "stdin": "",
        "script": "<?php echo 'Hello Kitty\n';"
    }------WebKitFormBoundary7MA4YWxkTrZu0gW--

    ![image](https://gitee.com/shijiaoliang/src/raw/master/51cto/run-code.png)

## License

run-code source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
