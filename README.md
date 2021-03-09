# api-gateway


sudo rabbitmq-plugins enable rabbitmq_management

# REQUIREMENTS

https://golang.github.io/dep/docs/installation.html

GOBIN 
export PATH=$PATH:$(go env GOPATH)/bin

#------------
# STARTUP
#-----------

service redis-server start
service rabbitmq-server start

tail -f api_log_****

#------------
# DOCUMENTATION
#-----------

http://127.0.0.1:1323/

https://echo.labstack.com/guide

https://github.com/eurie-inc/echo-sample

https://golang.org/doc/effective_go.html

#-----------
# RABBITMQ
#-----------

https://gist.github.com/fernandoaleman/72f0ad39e11915c0077d544b50096b50 (install for ubuntu 16)
https://www.rabbitmq.com/management.html (install management interface)

http://localhost:15672/
user:admin
password:password

https://wiki.betgames.tv/index.php?title=Web_integration_API_1.4#Token_system