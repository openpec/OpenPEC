## Resumo

O DATASUS, departamento responsável pelo desenvolvimento de sistemas auxiliares ao SUS, já implementou diversas ferramentas para uso nas Unidades de Saúde nacionais, porém, esses sistemas apresentam baixa interoperabilidade entre eles, o que dificulta relacionar os dados existentes. Além disso, os sistemas são todos de código fechado, não possibilitando a edição do código por terceiros, fato que impede a adição de funcionalidades úteis aos profissionais de saúde.

Atualmente, o sistema mais importante da atenção básica nacional é o Prontuário Eletrônico do Cidadão (PEC), responsável, principalmente, por coletar as informações clínicas e administrativas do paciente durante atendimentos realizados no contexto das Unidades Básicas de Saúde (UBS), ele é disponibilizado pelo DATASUS para uso, mas é possível a utilização de sistemas desenvolvidos por terceiros que atuem em seu lugar. Portanto, objetiva-se neste projeto estruturar uma arquitetura de referência de código aberto que atenda os requisitos da estratégia e-SUS Atenção Básica e que contenha funcionalidades que agentes de saúde considerem importantes. Esse sistema, intitulado OpenPEC, está sendo desenvolvido com base nas documentações oficiais disponibilizadas pelo Ministério da Saúde, além de materiais disponibilizados por terceiros que atuam na área da
saúde.

## Contexto

A estratégia **e-SUS Atenção Básica** (e-SUS AB) foi planejada pelo Departamento de Atenção Básica do Ministério da Saúde (DAB/MS) com base nas diretrizes da Política Nacional da Atenção Básica (PNAB) que almeja, através da informatização dos estabelecimentos de saúde, reestruturar e integrar as informações da Atenção Básica em nível nacional, facilitando e agilizando o trabalho das unidades de saúde ao inserir os processos de coleta, envio e gestão da informação nas atividades rotineiras dos profissionais (1). O e-SUS AB é, também, um sistema de software público que atua captando dados de consultas nas unidades de saúde, existindo dois formatos de sistema, o PEC e o CDS.

Esse projeto foca no primeiro citado, o **Prontuário Eletrônico do Cidadão (PEC)**, que consiste em um software, de **código fechado**, que coleta todas as informações clínicas e administrativas do paciente durante atendimentos realizados no contexto das Unidade Básica de Saúde (UBS), sendo que seu principal objetivo é informatizar o fluxo de atendimento do cidadão realizado pelos profissionais de saúde (2), ou seja, atuar diretamente como um prontuário eletrônico. Além disso, o PEC possui outras funções pra atender necessidades dos usuários, como a criação do cadastro do cidadão no sistema, uma agenda pra controle de consultas e horários, funções administrativas e geração de relatórios dos dados (3). Assim, além de auxiliar os atendimentos da Atenção Básica, o PEC apresenta benefícios para gestores, profissionais de saúde e cidadãos.

Além dos sistemas PEC e CDS, é possível que a unidade de saúde utilize um sistema próprio, integrando-o ao PEC e consequentemente ao Sistema de Informação em Saúde para a Atenção Básica (SISAB), que é responsável por fazer o processamento e a disseminação de dados e informações relacionadas à Atenção Básica nacional. Assim, a estratégia permite que sejam criados sistemas próprios com mais capacidades que o PEC, considerando aspectos que esse último não atende.

# OpenPEC

## Estrutura da Aplicação

Foi definido que o OpenPEC teria uma interface semelhante ao PEC e-SUS, a fim de ser fácil a interação de novos usuários. A
Figura 1 apresenta a tela do PEC Treinamento referente à área de administração de um funcionário definido como Administrador. É importante notar que há telas diferentes para lotações diferentes.

![Imagem 1](https://github.com/openpec/OpenPEC/blob/master/Assets/pec.PNG?raw=true)

Figura 1 – Tela de Administrador do PEC Treinamento versão 3.1.11. Fonte: Screenshot da tela do PEC Treinamento.

Quanto à estrutura da aplicação, o desenvolvimento do OpenPEC contempla a criação de aplicação Web e _container_ Docker para implantação rápida do sistema. Nesta fase serão utilizadas metodologias ágeis de desenvolvimento de _software_ e técnicas e linguagens de programação web modernas, tais como Javascript, CSS3 e HTML5 no _frontend_ e a arquitetura de microsserviços e a linguagem Golang no _backend_ da aplicação, objetivando alta performance e alta escalabilidade.

## Exportação de Dados
Para fazer a integração de sistemas próprios ao SISAB é necessário a exportação de um arquivo no formato Thrift, que
é importado pela aplicação PEC Centralizador, responsável por enviar os dados para o SISAB. A Figura 1 explicita com mais clareza esse processo.

![Imagem 2](https://github.com/openpec/OpenPEC/blob/master/Assets/openPec.png?raw=true)

Figura 2 - Fluxo de transmissão de dados para integração de sistemas próprios com o SISAB. Fonte: Adaptado do Manual do PEC (4).

As especificações técnicas para integração, e os padrões usados são disponibilizados pelo Ministério da Saúde no Manual de Exportação e-SUS (5) e no Layout e-SUS AB de Dados e Interfaces (LEDI AB) (6), sendo necessário apenas a execução dos métodos disponibilizados. Nesse projeto, foi utilizada a linguagem Java para implementação dos arquivos no formato Thrift.

### Contato

Para mais informações, sugestões ou critícas, entre em contato pelo e-mail openpec@gmail.com.

### Referências

(1) O que é e-SUS AB. Disponível em: http://dab.saude.gov.br/portaldab/o_que_e_esus_ab.php.

(2) O que é Prontuário Eletrônico do Cidadão? Disponível em: http://dab.saude.gov.br/portaldab/noticias.php?conteudo=_&cod=2300.

(3) e-SUS Atenção Básica : manual de implantação / Ministério da Saúde. Disponível em: http://189.28.128.100/dab/docs/portaldab/documentos/manual_implantacao_esus.pdf.

(4) Manual de Uso do Sistema com Prontuário Eletrônico do Cidadão - PEC. Disponível em: http://189.28.128.100/dab/docs/portaldab/documentos/esus/Manual_PEc_3_1.pdf.

(5) Manual de Exportação. Departamento de Atenção Básica.

(6) Layout e-SUS AB de Dados e Interface Versão 3.0.1. Disponível em: https://integracao.esusab.ufsc.br/.


