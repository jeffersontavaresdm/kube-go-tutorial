# Conectando a um cluster Kubernetes utilizando GO: Um tutorial passo a passo

### Introdução
A aplicação GO que acompanha este repositório é um exemplo simples de como listar os pods em um cluster Kubernetes usando a biblioteca cliente-go e Kind para criar o cluster Kubernetes.<br>O objetivo deste exemplo é demonstrar como interagir com o Kubernetes usando a linguagem GO e fornecer uma base para projetos mais complexos.

### Pré-requisitos
- Docker instalado em sua máquina. Você pode baixá-lo no site oficial: [Docker](https://www.docker.com/)
- Kind instalado em sua máquina. Você pode baixá-lo no site oficial: [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- Kubectl instalado em sua máquina. Você pode baixá-lo no site oficial: [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

### Estrutura
A estrutura deste projeto é relativamente simples. A pasta raiz contém o arquivo main.go, que é o arquivo principal da aplicação GO. Ele contém o código que lista os pods em um cluster Kubernetes usando a biblioteca cliente-go. Além disso, o repositório contém pastas de exemplos, onde você pode encontrar arquivos YAML que podem ser usados para criar um cluster Kubernetes usando Kind e criar um Pod para testar a aplicação GO.

### Contribuindo
Estou sempre buscando melhorar meus conhecimentos em GO e aprender novas tecnologias, adoraría receber contribuições da comunidade. Se você quiser relatar um problema, sugerir uma melhoria ou enviar uma solicitação de recebimento, basta abrir uma nova issue ou um novo pull request neste repositório. Certifique-se de descrever o problema ou melhoria com detalhes suficientes para que eu possa entender o que você está pedindo.

### Como utilizar esta aplicação em um cluster Kubernetes usando Kind
#### Passo a passo
1. Crie um arquivo `kind.yaml` para definir o cluster Kind. Exemplo: [kind.yaml](examples/cluster/kind.yaml)
2. Crie um cluster com o Kind executando o seguinte comando: `kind create cluster --config kind.yaml`
3. Verifique se o cluster está em execução com o comando: `kubectl cluster-info`
4. Crie um arquivo de configuração kubeconfig.yaml com as configurações do cluster Kind. Você pode usar o seguinte comando para gerar o arquivo: `kind get kubeconfig > kubeconfig.yaml`
5. Crie um arquivo `hello-word.yaml` que defina um Pod Kubernetes simples. Exemplo: [hello-word.yaml](examples/pod/hello-word.yaml.yaml)
6. Execute o seguinte comando no terminal para criar o Pod no cluster Kind: `kubectl apply -f hello-world.yaml`
7. Verifique se o Pod foi criado com sucesso executando o seguinte comando: `kubectl get pods`
8. Execute o seguinte comando para compilar o código GO e criar um executável: `go build main.go`
9. Execute o seguinte comando para executar a aplicação GO no cluster Kind: `go run main.go`

### Objetivo e uso da aplicação GO
O objetivo desta aplicação é demonstrar como é possível se conectar a um cluster Kubernetes local usando Kind e listar os nomes dos Pods em um determinado namespace. É útil para quem deseja entender como trabalhar com o Kubernetes usando GO.

### Exemplo de uso
Após executar o passo a passo descrito neste tutorial, você terá uma aplicação GO que lista os nomes dos Pods em um determinado namespace no cluster Kubernetes local usando Kind.<br>Você pode usar esse exemplo como ponto de partida para desenvolver suas próprias aplicações GO que interagem com o Kubernetes.

### Conclusão
Espero que este tutorial tenha sido útil para você entender como conectar uma aplicação GO a um cluster Kubernetes local usando Kind.<br>Se você tiver alguma dúvida ou sugestão, por favor, não hesite em entrar em contato.<br>Obrigado por utilizar este exemplo de aplicação em GO e bons estudos!
