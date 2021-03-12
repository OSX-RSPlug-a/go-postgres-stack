# go-postgres-stack

Sobre esse projeto teste:

Desenvolver uma aplicação web que responde por GET com um JSON, vindo de uma tabela em um banco relacional. 
A estrutura deve ter um processo de automação para a construção da aplicação e utilizar de container para o funcionamento da aplicação.


Tecnologias utilizadas:

    - Golang;
    - Postgresql;
    - Pgadmin;
    - Shell;
    - Docker;
    - Nginx;
    - NETDATA;



O projeto utilizou as libs da linguagem golang:

    - github.com/gorilla/mux 
    - github.com/jinzhu/gorm 
    - github.com/lib/pq 
    - github.com/rs/cors
	  

No inicio do código construímos a estrutura das tabelas e seus campos com o gorm;
Em seguida criamos os dados a serem inseridos;
Após a esrtutura do banco criada, criaremos as rotas com a lib mux para as requisições http da aplicação;
Então também criaremos as variavies com as informações de conexão; 
Em seguida faremos todas as funções necessarias para buscar as informações em banco de dados relacionado a cada rota; 


Iniciar the projeto em desenvolvimento:

      go run main.go
      


Instalar e deploy no cloud provider:

      git clone https://github.com/OSX-RSPlug-a/go-postgres-stack.git

      cd go-postgres-stack

      chmod +x initServer-install.sh
      
      sudo ./initServer-install.sh
  
      docker-compose -f gopost-stack.yaml up -d
      
