## Índice

- [Introdução](#introdução)
    - [Resumo](#resumo)
    - [Contexto](#contexto)
- [OpenPEC](#openpec)
    - [Estrutura da Aplicação](#estrutura-da-aplicação)
- [Artigo Publicado](#artigo)
- [Contato](#contato)
- [Referências do Readme](#referências)

## Introdução

### Resumo

Nos últimos anos o Ministério da Saúde tem feito um esforço a fim de melhorar os sistemas de saúde, dentre as iniciativas, destaca-se a estratégia e-SUS Atenção Básica (e-SUS AB), que tem como principal objetivo informatizar e melhorar os atendimentos de saúde no âmbito da atenção básica nacional. Atualmente, o sistema mais importante da atenção básica é o Prontuário Eletrônico do Cidadão (PEC), responsável por registrar dados de atendimentos de saúde. O Ministério da Saúde disponibiliza uma versão do PEC, mas os municípios podem optar por versões desenvolvidas localmente. Este projeto objetiva o desenvolvimento de uma arquitetura de referência como um projeto de Software Livre que atenda aos requisitos da estratégia e-SUS AB e do Sistema de Informação em Saúde para a Atenção Básica (SISAB) e que contenha funcionalidades similares ao PEC. Esse sistema, intitulado OpenPEC, tem como principal vantagem o fato de ser de código aberto, permitindo a edição e adição de novas utilidades, além de possibilitar que o projeto esteja sempre evoluindo. O OpenPEC está sendo desenvolvido com base nas documentações oficiais disponibilizadas pelo Ministério da Saúde, utilizando linguagens e técnicas de programação web modernas.

### Contexto

A estratégia **e-SUS Atenção Básica** (e-SUS AB) foi planejada pelo Departamento de Atenção Básica do Ministério da Saúde (DAB/MS) com base nas diretrizes da Política Nacional da Atenção Básica (PNAB) que almeja, através da informatização dos estabelecimentos de saúde, reestruturar e integrar as informações da Atenção Básica em nível nacional, facilitando e agilizando o trabalho das unidades de saúde ao inserir os processos de coleta, envio e gestão da informação nas atividades rotineiras dos profissionais (1). O e-SUS AB é, também, um sistema de software público que atua captando dados de consultas nas unidades de saúde, existindo dois formatos de sistema, o PEC e o CDS.

Esse projeto foca no primeiro citado, o **Prontuário Eletrônico do Cidadão (PEC)**, que consiste em um software, de **código fechado**, que coleta todas as informações clínicas e administrativas do paciente durante atendimentos realizados no contexto das Unidade Básica de Saúde (UBS), sendo que seu principal objetivo é informatizar o fluxo de atendimento do cidadão realizado pelos profissionais de saúde (2), ou seja, atuar diretamente como um prontuário eletrônico. Além disso, o PEC possui outras funções pra atender necessidades dos usuários, como a criação do cadastro do cidadão no sistema, uma agenda pra controle de consultas e horários, funções administrativas e geração de relatórios dos dados (3). Assim, além de auxiliar os atendimentos da Atenção Básica, o PEC apresenta benefícios para gestores, profissionais de saúde e cidadãos.

Além dos sistemas PEC e CDS, é possível que a unidade de saúde utilize um sistema próprio, integrando-o ao PEC e consequentemente ao Sistema de Informação em Saúde para a Atenção Básica (SISAB), que é responsável por fazer o processamento e a disseminação de dados e informações relacionadas à Atenção Básica nacional. Assim, a estratégia permite que sejam criados sistemas próprios com mais capacidades que o PEC, considerando aspectos que esse último não atende.

## OpenPEC

### Estrutura da Aplicação

Foi definido que o OpenPEC teria uma interface e cores semelhantes ao PEC e-SUS, a fim de ser fácil a interação de novos usuários. A
Figura 1 apresenta a tela de login do OpenPEC, que solicita o CPF e senha do usuário, assim como o PEC.

![Imagem 1](https://github.com/openpec/OpenPEC/blob/master/readmeAssets/loginScreen.png?raw=true)

Figura 1 – Tela de Login do OpenPEC.

O primeiro cadastrado do sistema recebe o status de administrador, podendo configurar a lotação dos outros cadastrados, assim como definir novos administradores.

Após o login, o usuário é redirecionado para a página principal da aplicação, onde ele pode usar as funcionalidades do sistema, utilizando as funções relacionadas a saúde ou ao seu perfil. A Figura 2 apresenta a tela inicial do OpenPEC, após o login.

![Imagem 2](https://github.com/openpec/OpenPEC/blob/master/readmeAssets/homeScreen.png?raw=true)

Figura 2 – Tela de Início do OpenPEC.

Quanto à estrutura da aplicação, o desenvolvimento do OpenPEC contempla a criação de aplicação Web e _container_ Docker para implantação rápida do sistema. Nesta fase estão sendo utilizadas metodologias ágeis de desenvolvimento de _software_ e técnicas e linguagens de programação web modernas, tais como HTML5 e CSS3 no _frontend_ e a linguagem Golang no _backend_ da aplicação, objetivando alta performance e alta escalabilidade.



## Artigo

Mais informações quanto à contextualização do projeto podem ser lidas no artigo da revista Temas em Saúde [clicando aqui](http://temasemsaude.com/wp-content/uploads/2020/06/20303.pdf). É preciso notar que o projeto sofreu algumas alterações nas metodologias usadas, diferindo do que foi descrito no artigo.


### Contato

Para mais informações, sugestões ou critícas, entre em contato pelo e-mail openpec@gmail.com.


### Referências

(1) O que é e-SUS AB. Disponível em: http://dab.saude.gov.br/portaldab/o_que_e_esus_ab.php.

(2) O que é Prontuário Eletrônico do Cidadão? Disponível em: http://dab.saude.gov.br/portaldab/noticias.php?conteudo=_&cod=2300.

(3) e-SUS Atenção Básica : manual de implantação / Ministério da Saúde. Disponível em: http://189.28.128.100/dab/docs/portaldab/documentos/manual_implantacao_esus.pdf.


