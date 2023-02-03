<h1 align="center">Тестирование blockchain кошельков при пакетном обновлении контраката etherium</h1>

## Реализованные функции
- Реализован контракт solidity для создания кошелька с n балансом, получение информации по созданным кошелькам, выполнение отправки баланса с одного кошелька на другой
- Реализован CLI для взаимодействия с сетью блокчейн на базе языка Go


## Download dependencies
```
go get all
```

## Создание билда и генерирование контракта для взаимодействия с go
Build abi, bin & auto-generate SimpleStorage.go file
```
solcjs --optimize --abi --bin ./SimpleStorage.sol -o build
mkdir api
abigen --bin=./build/Store_sol_Store.bin --abi=./build/Store_sol_Store.abi --pkg=api --out=./api/Store.go

```

## Ganache
Запустить Ganache Etherium

## Необходимо с Ganache взять переременнные и установить их в переменную среду app.env
```
gateway = <адрес с портом etherium>
accountPrivateKey = <Приватный ключ аккаунта etherium>
accountHexAddress = <Hex адресс используемого аккаунта>
smartContractHexAddress = <Hex адресс задеплоиного контракта>
```
## Взаимодействие с CLI
находясь в корневом каталоге проекта выполняем
```
cd cmd/pkg/app
```

### Команды CLI
#### Деплой контакта
```
go run main.go deployContract 
```
Пример положительного результата (smartContractHexAddress в последующем используется для взаимодействие с контрактом) :
``
smartContractHexAddress:  0x0784F2DBD79348E6C3Ff82C39F4065930aa49cFe
tx:  0x67a3fdacf5c7f7fd9c4c43faf9b87c5e02d2da805359b7b4756c4aca6dcf9436
``
#### Создание кошелька
```
go run main.go createWallet one 30000
```
Пример положительного результата:
``
tx sent: 0xf7caecf5cff33f55336fbe4d6a2c9ea8d37e69f8cf419a71d225df57417fcf6a
``
#### Отправка валюты с одного кошелька на другой
```
go run main.go sendBalance one two 3000
```
Пример положительного результата:
``
tx sent: 0xa9a70c91dd83f3d0717d13bff0c939ff25edf728f1a9c4a5fdf87f7486f9abde
``
#### Получение информации о баласе кошелька
```
go run main.go readWallet two
```
Пример положительного результата:
``
{two 58000}
``
## Запуск модульных и интеграционных тестов
находясь в корневом каталоге проекта выполняем
```
cd cmd/pkg/blockchain
```
выполнить команду запуска тестов
```
go test 
```