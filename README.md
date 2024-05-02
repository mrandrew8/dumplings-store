# Пельменная №2

<img width="900" alt="image" src="https://user-images.githubusercontent.com/9394918/167876466-2c530828-d658-4efe-9064-825626cc6db5.png">


## Инфраструктура
```bash
Предварительные настройки

1) Установка Terraform
Скачиваем дистрибутив terraform_1.6.6_linux_amd64.zip из зеркала. Ссылка на зеркало https://hashicorp-releases.yandexcloud.net/terraform/1.6.6/
После загрузки добавляем путь к папке, в которой находится исполняемый файл, в переменную PATH: export PATH=$PATH:/path/to/terraform
Указываем источник установки провайдера, добавив следующую конфигурацию в файл ~/.terraformrc (если такого файла нет на вашей машине, его нужно создать):
provider_installation {
  network_mirror {
    url = "https://terraform-mirror.yandexcloud.net/"
    include = ["registry.terraform.io/*/*"]
  }
  direct {
    exclude = ["registry.terraform.io/*/*"]
  }
} 

2) Установка и настройка интерфейса командной строки Yandex Cloud (CLI)
Инструкция по ссылке https://yandex.cloud/ru/docs/cli/quickstart#install
curl -sSL https://storage.yandexcloud.net/yandexcloud-yc/install.sh | bash
Скрипт установит CLI и добавит путь до исполняемого файла в переменную окружения PATH.
После завершения установки перезапускаем командную оболочку.

Аутентифицируемся с помощью аккаунта на Яндекс
Получаем OAuth-токен в сервисе Яндекс ID https://yandex.ru/dev/id/doc/ru/concepts/ya-oauth-intro
Перейдем по ссылке https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb. Если приложение запрашивает доступ к данным, разрешаем. Это нужно для получения токена.
Копируем в буфер обмена или сохраняем полученный токен.

Если вы аутентифицируетесь впервые, перейдите в консоль облака. Примите условия лицензионного соглашения и политики конфиденциальности.

Чтобы начать настройку профиля CLI, выполняем команду: yc init
Выберите профиль, для которого вы хотите настроить аутентификацию, или создайте новый. Если вы выполняете команду yc init впервые, этот шаг будет отсутствовать. 
Мы же выбираем создание нового профиля "Create a new profile":
Pick desired action:
[1] Re-initialize this profile 'default' with new settings
[2] Create a new profile
Please enter your numeric choice: 1

По запросу команды вводим OAuth-токен, полученный ранее:
Please go to https://oauth.yandex.ru/authorize?response_type=token&client_id=<клиент айди>
in order to obtain OAuth token.
Please enter OAuth token: y0_AgA ... wvs7N4

Выберите одно из предложенных облаков, к которым у вас есть доступ:

Please select cloud to use:
 [1] cloud1 (id = aoe2bmdcvata********)
 [2] cloud2 (id = dcvatao4faoe********)
Please enter your numeric choice: 2
Если вам доступно только одно облако, оно будет выбрано автоматически.

Выбераем каталог по умолчанию:

Please choose a folder to use:
 [1] folder1 (id = cvatao4faoe2********)
 [2] folder2 (id = tao4faoe2cva********)
 [3] Create a new folder
Please enter your numeric choice: 1
Чтобы выбрать зону доступности по умолчанию для сервиса Compute Cloud, введите Y. Чтобы пропустить настройку, введите n.

Выбираем зону по умолчанию ru-central1-a
Do you want to configure a default Yandex Compute Cloud availability zone? [Y/n] Y
Если вы ввели Y, выберите зону доступности:

Which zone do you want to use as a profile default?
 [1] ru-central1-a
 [2] ru-central1-b
 [3] ru-central1-c
 [4] ru-central1-d
 [5] Do not set default zone
Please enter your numeric choice: 2

Проверяем настройки нашего профиля CLI:
yc config list

Результат:
token: y0_AgA...wvs7N4
cloud-id: b1g159pa15cd********
folder-id: b1g8o9jbt58********
compute-default-zone: ru-central1-a

Настраиваем управление ресурсами от имени аккаунта на Яндексе.
Добавляем аутентификационные данные в переменные окружения:

export YC_TOKEN=$(yc iam create-token)
export YC_CLOUD_ID=$(yc config get cloud-id)
export YC_FOLDER_ID=$(yc config get folder-id)

Где:
YC_TOKEN — IAM-токен.
YC_CLOUD_ID — идентификатор облака.
YC_FOLDER_ID — идентификатор каталога.

Примечание
Время жизни IAM-токена — не больше 12 часов, но рекомендуется запрашивать его чаще, например каждый час.




3) Создание ресурсов в Yandex Cloud

На странице Yandex Cloud Billing убедитесь, что у вас подключен платежный аккаунт, и он находится в статусе ACTIVE или TRIAL_ACTIVE. Если платежного аккаунта нет, создайте его. Если у вас еще нет каталога, создайте его. 
Мы же будем использовать дефолтный каталог (default).

Ниже приведенеы описания модулей Terraform

module "tf-yc-network"
module "tf-yc-service-account"
module "tf-yc-kms-symmetric-key"
module "tf-yc-sgroups"
module "tf-yc-k8s-cluster" 
module "tf-yc-node-group" 
module "tf-yc-dns"
module "tf-yc-static-public-ip"
module "tf-yc-storage-bucket"


3.1) module "tf-yc-network" - создание сервисного аккаунта k8s-account
От имени этого сервисного аккаунта будут создаваться ресурсы, необходимые кластеру Managed Service for Kubernetes.
Создаем согласно документации https://yandex.cloud/ru/docs/managed-kubernetes/operations/kubernetes-cluster/kubernetes-cluster-create#tf_1 сервисный аккаунт с ролями: k8s.clusters.agent, vpc.publicAdmin, container-registry.images.puller, kms.keys.encrypterDecrypter, editor, load-balancer.admin.
Роль k8s.clusters.agent — специальная роль для сервисного аккаунта кластера Kubernetes. Дает право на создание групп узлов, дисков, внутренних балансировщиков. Позволяет использовать заранее созданные ключи Yandex Key Management Service для шифрования и расшифровки секретов, а также подключать заранее созданные группы безопасности. 
В комбинации с ролью load-balancer.admin позволяет создать сетевой балансировщик нагрузки с публичным IP-адресом.
Роль vpc.publicAdmin позволяет управлять NAT-шлюзами, публичными IP-адресами и внешней сетевой связностью, а также просматривать информацию о квотах, ресурсах и операциях с ресурсами сервиса. Роль предоставляет права администратора мультиинтерфейсных ВМ, обеспечивающих связность между несколькими сетями.
Роль container-registry.images.puller позволяет скачивать Docker-образы, а также просматривать информацию о ресурсах сервиса (реестрах, Docker-образах, репозиториях).
Роль kms.keys.encrypterDecrypter позволяет шифровать и расшифровывать данные, а также просматривать информацию о ключах. Включает все 
права ролей kms.keys.encrypter и kms.keys.decrypter.
Роль editor дает разрешения на все операции для управления ресурсом, кроме назначения ролей другим пользователям.


3.2) module "tf-yc-sgroups" - создание группы правил безопасности k8s-public-services
k8s-public-services включает в себяя набор следующих правил:
- Правило, которое разрешает проверки доступности с диапазона адресов балансировщика нагрузки. 
Нужно для работы отказоустойчивого кластера Managed Service for Kubernetes и сервисов балансировщика."
- "Правило, которое разрешает взаимодействие мастер-узел и узел-узел внутри группы безопасности."
- "Правило, которое разрешает взаимодействие под-под и сервис-сервис. Укажите подсети вашего кластера Managed Service for Kubernetes и сервисов."
- "Правило, которое разрешает отладочные ICMP-пакеты из внутренних подсетей."
- "Правило, которое разрешает входящий трафик из интернета на диапазон портов NodePort. Добавьте или измените порты на нужные вам."
- "Правило, которое разрешает весь исходящий трафик. Узлы могут связаться с Yandex Container Registry, Yandex Object Storage, Docker Hub и т. д."
- "Правило, которое разрешает подключение к API Kubernetes через порт 6443 из указанной сети."
- "Правило, которое разрешает подключение к API Kubernetes через порт 443 из указанной сети."
- "Правило, которое разрешает подключение к узлам по SSH с указанных IP-адресов

3.3) module "tf-yc-kms-symmetric-key" - создание ключа шифрования Key Management Service kms-key
Секрет — конфиденциальная информация, используемая кластером Kubernetes при управлении подами, например, OAuth-ключи, пароли, SSH-ключи и т. д. 
По умолчанию Kubernetes хранит секреты в открытом виде. Для защиты секретов сервис Managed Service for Kubernetes позволяет шифровать их с помощью ключей шифрования из сервиса Yandex Key Management Service. 
Для работы с ключами используется механизм Key Management Service-провайдеров Kubernetes.
Managed Service for Kubernetes использует для шифрования и расшифровки ключей Key Management Service-плагин. Секреты шифруются стандартными средствами Kubernetes.

3.4) module "tf-yc-network" - создание сети dumpling-networ и подсети dumpling-subnet
Подсети создаются в облачных сетях. Созданная подсеть размещается в одной из зон доступности. К подсети можно подключить ресурсы из той же зоны, где находится подсеть.

3.5) module "tf-yc-k8s-cluster" - создание зонального кластера Kubernetes k8s-zonal

3.6) module "tf-yc-node-group" - создание групп узлов k8s-node-group. 
Группу узлов Managed Service for Kubernetes создаем с автомасштабированием, делаем это добавив блок auto_scale.

3.7) module "tf-yc-dns" - создание dns зоны dumpling-zone1

3.8) module "tf-yc-static-public-ip" - Создание публичного статического ip адреса static-public-ip.
Потребуется для ingress контроллера

3.9) module "tf-yc-storage-bucket" - создание объектного хранилища s3 bucket.
Нужен для хранения terraform state, а так же для хранения картинок пельмешек.

После того как описали конфиги, выполняем:
terraform init - для инициализации провайдера и скачивания плагинов
terraform plan - для получения плана создания (изменения, удаления) ресурсов
terraform apply - для создания (изменения, удаления) ресурсов 

Ресурсы созданы

```

## Сборка образа при помощи Dockerfile

```bash
Описываем сборку образов backend и frontend в Dockerfile, который состоит из двух этапов: сборка и релиз. Собранные артефакты из этапа сборки 
копируются в релизный этап, который в свою очередь собирает основной образ для запуска контейнеров.

Настраиваем конфиг nginx следующим образом:

server {
  listen 80;

  location / {
    root /usr/share/nginx/html;
    index index.html;
    try_files $uri $uri/ /index.html;
  }

  location ~ /(products|categories|orders|auth/whoami|metrics) {
    proxy_pass http://dumpling-backend:8081;
  }
}

Директива listen 80; указывает на то, что сервер будет прослушивать запросы на порту 80, что является стандартным портом для HTTP-серверов.

Настройка корневой директории и обработки запросов: В блоке location / указывается корневая директория для обработки запросов, а также устанавливается индексный файл index.html. Директива try_files $uri $uri/ /index.html; определяет, как сервер будет обрабатывать запросы, пытаясь сначала найти запрошенный файл, затем папку, и, если ничего не найдено, возвращая index.html.

Проксирование запросов для определенных путей: В блоке location ~ /(products|categories|orders|auth/whoami|metrics) указывается регулярное выражение для путей, по которым будет осуществляться проксирование запросов. В данном случае, запросы по путям /products, /categories, /orders, /auth/whoami и /metrics будут направляться на сервер с адресом http://dumpling-backend:8081 с помощью директивы proxy_pass.

Для  файлик vue.config.js в директории фронтенда следующим образом 

module.exports = {
  devServer: {
    disableHostCheck: true
  },
  publicPath: process.env.NODE_ENV === 'production'
    ? '/'
    : '/'
};

Параметр publicPath устанавливается в зависимости от значения переменной окружения NODE_ENV. Так как у нас одна среда, для обоих случаев проставляем значение / для publicPath


```



## Установка Ingress-контроллера NGINX с менеджером для сертификатов Let's Encrypt

```bash

Покупаем домен у любого провайдера, в нашем случае рег.ру и делегируем домен. Прописываем у регистратора сервера ns1.yandexcloud.net. и ns2.yandexcloud.net.

Установка Ingress-контроллера NGINX с менеджером для сертификатов Let's Encrypt производится согласно инструкции https://yandex.cloud/ru/docs/managed-kubernetes/tutorials/ingress-cert-manager#install-controller

Предварительные действия:
Устанавливаем хелм. Ссылка на инструкцию: https://helm.sh/ru/docs/intro/install/


Для установки Helm-чарта с Ingress-контроллером NGINX выполняем команду:
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && \
helm repo update && \
helm install ingress-nginx ingress-nginx/ingress-nginx --set controller.service.loadBalancerIP=Static_IP

где Static_IP - публичный статический ip адрес созданный терраформом в модуле tf-yc-static-public-ip. 
Созданный контроллер установлен за Yandex Network Load Balancer.

Вручную устанавливаем менеджер сертификатов:

Установил версию 1.12.1 менеджера сертификатов, настроенного для выпуска сертификатов от Let's Encrypt®. 
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.1/cert-manager.yaml

Убедился, что в пространстве имен cert-manager создано три пода с готовностью 1/1 и статусом Running:
kubectl get pods -n cert-manager --watch
Результат:

NAME                                      READY  STATUS   RESTARTS  AGE
cert-manager-69********-ghw6s             1/1    Running  0         54s
cert-manager-cainjector-76********-gnrzz  1/1    Running  0         55s
cert-manager-webhook-77********-wz9bh     1/1    Running  0         54s

Создал YAML-файл /kubernetes/acme-issuer.yaml с манифестом объекта ClusterIssuer:

kubectl apply -f acme-issuer.yaml

Настройка днс записи:

Так как мы не используем ExternalDNS c плагином для Cloud DNS, то нужно настраивать DNS-запись вручную:
Узнаем IP-адрес Ingress-контроллера (значение в колонке EXTERNAL-IP) (у нас это StaticIP):
kubectl get svc

Результат:

NAME                      TYPE          CLUSTER-IP     EXTERNAL-IP     PORT(S)                     AGE
...
ingress-nginx-controller  LoadBalancer  10.96.164.252  84.201.153.122  80:31248/TCP,443:31151/TCP  2m19s
...
Размещаем в ДНС зоне A-запись, указывающую на публичный IP-адрес Ingress-контроллера:
<наш_домен> IN A <IP-адрес_Ingress-контроллера>

Проверяем работоспособность TLS
kubectl describe certificate domain-name-secret

В выводе команды содержатся подобные события (events):

Events:
  Type    Reason     Age   From                                       Message
  ----    ------     ----  ----                                       -------
  Normal  Issuing    ...   cert-manager-certificates-trigger          Issuing certificate as Secret does not exist
  Normal  Generated  ...   cert-manager-certificates-key-manager      Stored new private key in temporary Secret resource...
  
Сертификаты используются в связке с соответствующими им секретами Kubernetes, которые хранят пары ключей и служебную информацию. 
В случае отсутствия секрета сертификат перевыпускается автоматически с созданием нового секрета, о чем и сообщается в событиях. 

Поскольку сертификат выпускается впервые, то соответствующий ему секрет отсутствует. Наличие событий, сообщающих об этом, не является ошибкой.

Примечание: Проверка прав на домен сертификата Let's Encrypt® может занять несколько часов.

Выполняем команду:

curl https://<наш_домен>
Результат:

<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>

Примечание: Если ресурс недоступен по указанному URL, то убедитесь, что группы безопасности для кластера Managed Service for Kubernetes 
и его групп узлов настроены корректно. Если отсутствует какое-либо из правил — добавьте его.

```



## CI/CD
```bash


1) SonarQube 

Создаем проект в sonarqube по ссылке https://sonarqube.praktikum-services.ru/projects/favorite :
add project - указываем Project key и Display name - генерируем token - выбираем вариант, описывающий нашу сборку (и для frontend и для backend выбираем other) - выбираем OS Linux.

Получили команду для выполнения анализа средствами SonarQube, которую вставляем в секцию скрипт задания sonarqube-backend(frontend)-sast. Для безопасност сохраняем значения переменных в gitlab variables, а команду параметризуем

Пример команды: 
sonar-scaner \
  -Dsonar.projectKey=test \
  -Dsonar.sources=. \
  -Dsonar.host.url=https://sonarqube.praktikum-services.ru\
  -Dsonar.login=89138e3474346e0b22362874383a0.....


2) Создание статического файла конфигурации

Для получения доступ к кластеру Managed Service for Kubernetes без использования CLI, нам потребуется создание статического файла конфигурации.

Получите уникальный идентификатор кластера из Yandex Cloud и помещяем его в переменную:
CLUSTER_ID=catb3ppsdsh7********

Подготавливаем сертификат кластера и сохраняем сертификат кластера управляемой службы для Kubernetes в файл с именем ca.pem. Этот сертификат подтверждает подлинность управляемой службы для кластера Kubernetes.

Запускаем команду, которая:
Получает информацию об управляемой службе для кластера Kubernetes в формате JSON.
Оставляет только информацию о сертификате и удаляет лишние кавычки из содержимого сертификата.
Удаляет ненужные символы из содержимого сертификата.
Сохраняет сертификат в ca.pem файл.

yc managed-kubernetes cluster get --id $CLUSTER_ID --format json | \
  jq -r .master.master_auth.cluster_ca_certificate | \
  awk '{gsub(/\\n/,"\n")}1' > ca.pem

Примечание: Потребуется установка jq. Установить можно так: sudo apt install jq


Создаем объект ServiceAccount для взаимодействия с API Kubernetes внутри кластера управляемой службы для Kubernetes. Спецификация создания ServiceAccount объекта и его секрета сохраняем в /kubernetes/serviceaccount2.yaml

Выполняем: kubectl create -f serviceaccount2.yaml

Подготавливаем токен ServiceAccount. Токен необходим для ServiceAccount аутентификации в кластере Managed Service for Kubernetes.

Запускаем команду, которая:
Получает информацию о ранее созданной admin-user учетной записи службы в формате JSON.
Оставляет только информацию о токене и удаляет лишние кавычки из содержимого токена.
Декодирует токен из Base64.
Сохраняет содержимое токена в SA_TOKENпеременную.
SA_TOKEN=$(kubectl -n kube-system get secret $(kubectl -n kube-system get secret | \
  grep admin-user-token | \
  awk '{print $1}') -o json | \
  jq -r .data.token | \
  base64 -d)


Получаем IP-адрес кластера управляемой службы для кластера Kubernetes и добавляем его в MASTER_ENDPOINT переменную для дальнейшего использования.

Запускаем команду, которая:
Получает сведения об управляемой службе для кластера Kubernetes в формате JSON на основе его уникального идентификатора.
Оставляет только IP-адрес управляемой службы для кластера Kubernetes.
Удаляет лишние кавычки из своего содержимого.
Сохраняет IP-адрес в MASTER_ENDPOINTпеременной.
Для подключения к API кластера Managed Service for Kubernetes из Интернета (за пределами Яндекс Облака).

MASTER_ENDPOINT=$(yc managed-kubernetes cluster get --id $CLUSTER_ID \
  --format json | \
  jq -r .master.endpoints.external_v4_endpoint)

Добавляем данные о кластере Managed Service for Kubernetes в файл конфигурации

Запускаем эту команду:
kubectl config set-cluster sa-test2 \
  --certificate-authority=ca.pem \
  --server=$MASTER_ENDPOINT \
  --kubeconfig=test.kubeconfig

Добавляем информацию о токене admin-user в файл конфигурации:
kubectl config set-credentials admin-user \
  --token=$SA_TOKEN \
  --kubeconfig=test.kubeconfig

Добавьте контекстную информацию в файл конфигурации:
kubectl config set-context default \
  --cluster=sa-test2 \
  --user=admin-user \
  --kubeconfig=test.kubeconfig

Убеждаемся, что конфигурация правильная, выполнив следующую команду:

kubectl get namespace --kubeconfig=test.kubeconfig
Результат:

NAME     STATUS  AGE
default  Active  9d

Сохраняем значение из файлика ca.pem и test.kubeconfig в gitlab Variables kube_cert и kubeconfig соответственно.

4) Создание реппозитория в Nexus для хранения архивов Helm чартов

Переходим по ссылке: https://nexus.praktikum-services.tech/
Создаем новый реппозиторий: create reppository - выбираем тип helm(hosted) - задем имя реппозитория и жмем create reppository.
Реппозиторий создан. Для получения доступа к реппозиторию Nexus из CI/CD сохраняем следующие переменные в gitlab variables:
nexus_pass = <наш пароль от нексус>, nexus_repo_url = https://nexus.praktikum-services.tech/repository/dumplings_store_nikolaev/ , nexus_user = <наш логин от нексус>




Описание конфигов CI/CD

Создаем в корне проекта .gitlab-ci.yaml конфигурационный файл для определения процесса непрерывной интеграции и поставки (CI/CD).
Так же создаем .gitlab-ci.yaml в директориях backend, frontend, backend-chart, frontend-chart. И включаем их в родительский корневой .gitlab-ci.yml.

.gitlab-ci.yaml в директориях backend, frontend состоит из 4 этапов: build, test, release, notify.

build-backend(frontend)-code-job - здесь происходит сборка приложения (docker-in-docker)
sonarqube-backend(frontend)-sast - проверка кода бэкенда(фронтенда) на соответсвие quality gates заданным в sonarqube
backend-test (олько для backend)- задание для запуска тестов на GO
release - тегирование образа и его загрука в gitlab-registry
telegram-notification-backend(frontend) - уведомление в чат телеграмм о выходе новой версии приложения
Дополнительно добавляем этап проверки кода из коллекций GitLab Auto DevOps. Подключаем шаблон Security/SAST.gitlab-ci.yml, добавляя следующий код:

include:
  - template: Security/SAST.gitlab-ci.yml 


Деплой приложения будем проводить при помощи пакетного менеджера Helm. Для этого Создаем описываем два чарта backend-chart и frontend-chart.
.gitlab-ci.yaml в директориях backend-chart, frontend-chart состоит из 3 этапов: build, release, notify.

chart-realise - упаковка helm chart в архив
chart-deploy - загрузка архива в нексус реппозиторий
telegram-notification-backend(frontend) - уведомление в чат телеграмм о выходе новой версии helm chart

backend-chart содержит в себе шаблоны следующих объектов: deployment, secret, service.

Как сформировать secret.yaml. 
Аутентифицируемся в gitlab registry 
docker login gitlab.praktikum-services.ru:5050 -u <наш_логин> -p <наш пароль>
Кодируем в base64 наш докер конфиг:
cat /home/<user>/.docker/config.json | base64
Копируем полученное значение и вставляем .dockerconfigjson: >- в secret.yaml
Аналогично для фронтенда

frontend-chart содержит в себе шаблоны следующих объектов: deployment, secret, service, configmap.yaml, ingress.yaml.
В ingress.yaml в host прописываем наш домен



helm upgrade --atomic --install backend-chart ${nexus_repo_url}backend-${VERSION}.tgz --namespace default --username=${nexus_user} --password=${nexus_pass} --set backend.deployment.tag=$VERSION


```



## Глоссарий

```bash
Terraform
Yandex Cloud (CLI) - скачиваемое программное обеспечение для управления вашими облачными ресурсами через командную строку.
OAuth-токен
```