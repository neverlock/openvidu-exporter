## OpenVidu Exporter

focus on session api get only session number and all connection number
scrape data from openvidu api with basic auth 

```
 go get

 go build -o vidu-exporter
```

before run exporter you must edit config.yaml for host ,username and password

```
 cp config-template.yaml config.yaml

 vi config.yaml

 ./vidu-exporter
```

* You can add multiple domain via yaml with 2 parameter

```
 MaxSub and Vidu(number)
```

 for example

```yaml
Vidu:
   ListenHost: 0.0.0.0
   ListenPort: :8080
   MaxSub: 3
   Vidu1:
      URL: https://url1
      User: user1
      Password: pass1
   Vidu2:
      URL: https://url2
      User: user2
      Password: pass2
   Vidu3:
      URL: https://url3
      User: user3
      Password: pass3
```

