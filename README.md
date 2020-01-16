# drone-wechat-work

A drone plugin to push messages to WeChat Work with robot api.

## How to use

1. Pull image:

```sh
docker pull sqrtthree/drone-wechat-work
```

2. Use plugin:

```yml
kind: pipeline
type: docker
name: default

steps:
  # Your tasks.

  - name: notify
    image: sqrtthree/drone-wechat-work
    settings:
      hook_url:
        from_secret: hook_url
      # Or with key
      key:
        from_secret: key
    when:
      status:
        - failure
        - success
```

## Parameter Reference

### hook_url

Hook url of wechat for work. See [work.weixin.qq.com](https://work.weixin.qq.com/api/doc/90000/90136/91770) to get more details.

### key

Or you can also set the key if you don't want to set hook url.
