0 - Validar se possui protoc instalado - protoc --version

1 - Para não receber possíveis erros do go - go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

2 - Para acessar o profile - vim ~/.bash_profile

    2.1 - 
    export GO_PATH=~/go 
    export PATH=$PATH:/$GO_PATH/bin##

3 (OPT) - Caso a linha de cima não role - 
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin

4 - Para "compilar" o profile - source ~/.bash_profile

5 - Para compilar os proto buffer - protoc --go_out=. --go-grpc_out=. proto/*.proto

6 - Para adicionar todas as libs que os arquivos compilados usam - go get google.golang.org/grpc