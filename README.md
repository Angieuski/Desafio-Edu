Após fazer o download dos arquivos executar o arquivo main.exe.

Após a execução, com o servidor ativo, ir para o endereço: http://localhost:8080/uis
O usuário e senha são respectivamente: "juan" e "angieuski".(Somente o conteúdo entre as aspas)

Para as operações de CRUD será utilizado o aplicativo Postman, onde poderão ser exercídas das seguintes maneiras:
Em todas as opções à seguir será necessário selecionar o sistema de autorização "basic auth" e preencher com os dados citados anteriormente.

Para o Post utilizamos um body-raw no formato json feito no endereço http://localhost:8080/uis
{
    "nome": "XXXXXX",
    "email": "XXXXX@XXXXX.com,
    "senha": "XXXXXXXXXXX" (min 8 dígitos, máx 24 dígitos)
}

Para o Get teremos apenas o request de Get no endereço http://localhost:8080/uis

Para o Put devemos inserir o endereço http://localhost:8080/uis/@id, onde @id é referente ao ID do usuário que será atualizado e utilizar um body-raw no formato json assim como no Post.
{
    "nome": "XXXXXX",
    "email": "XXXXX@XXXXX.com,
    "senha": "XXXXXXXXXXX" (min 8 dígitos, máx 24 dígitos)
}
PS: não é necessário alterar todos os campos, somente o email por exemplo já funciona

Para o Delete seguimos a mesma ideia do Put, onde inserimos o endereço http://localhost:8080/uis/@id, onde @id é referente ao ID do usuário que será deletado.
