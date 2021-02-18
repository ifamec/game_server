# Game Server

- net 
    - websocket
- db
    - mysql
    - redis
- msg serialization
- ...

## websocket
- duplex communication, tcp proxy, keep-alive, web support

## protocol
- communication rules
  - s-c communication
  - block garbage msg
- clear dev routine
- publish as SDK

## message process
- network layer
  - main listen
    - socket
    - port
    - expose to network
  - websocket
    - isWebsocket?
      - no - return network layer data
      - yes - process if match rules or return error code
        - get data, process main protocol
          - no - return error code
          - yes - process child protocol
            - ...
        - http
        - tcp
    - serialization 