# mine-grrpc


生成私钥
openssl genrsa -out server.key 2048

证书文件
openssl req -nodes -new -x509 -sha256 -days 1825 -config openssl.cnf -extensions 'req_ext' -key server.key -out server.crt

openssl x509 -in server.crt -text

config:
  grpcServer:
  name: "stats-server"
  ip: 192.168.31.22
  port: 8022
  network: tcp
  address: ":8022"
  Weight: 100
jaeger:
  serviceName: "stats-server"
  hostPort: localhost:6831
db:
  name: "rec_visual"
  address: "localhost:3306"
  userName: root
  password: 123456
  maxIdleConnects: 10
  maxOpenConnects: 100
  connMaxLifetimeHour: 1
redis:
  address: "localhost:6379"
  userName: root
  password: 123456
  database: 0