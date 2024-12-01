package products

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

// Task one.
type DeliveryMethod interface {
	CalculateShippingCost(int) int
}

// Стратегия курьерской доставки.
type CourierDelivery struct{}

func (CourierDelivery) CalculateShippingCost(numder int) int {
	return numder + 100
}

// Стратегия почтовой доставки.
type PostDelivery struct{}

func (PostDelivery) CalculateShippingCost(number int) int {
	return number + 50
}

// Стратегия самовывоза.
type PickupDelivery struct{}

func (PickupDelivery) CalculateShippingCost(number int) int {
	return number
}

// Функция для расчета стоимости доставки
func CalculateOrderShippingCost(delivery DeliveryMethod, number int) int {
	return delivery.CalculateShippingCost(number)
}

// Task two.
type Character struct {
	attackStrategy AttackStrategy
}

type AttackStrategy interface {
	Attack() int
}

type SwordAttack struct{}

func (SwordAttack *SwordAttack) Attack() int {
	return 10
}

type BowAttack struct{}

func (BowAttack *BowAttack) Attack() int {
	return 15
}

type MagicAttack struct{}

func (MagicAttack *MagicAttack) Attack() int {
	return 20
}

func (Character *Character) SetAttackStrategy(action AttackStrategy) {
	Character.attackStrategy = action
}

func (Character *Character) AttackProcessing() int {
	if Character.attackStrategy == nil {
		fmt.Println("Нет метода атаки")
		return 0
	}
	return Character.attackStrategy.Attack()
}

//Паттерн "Синглтон".
//Task one.

type Logger struct {
	//Позволит вам работать с открытым файлом в других методах структуры Logger.
	file *os.File
	sync.RWMutex
}

// переменная для хранения единственного экземпляра
var instance *Logger
var once sync.Once

func GetInstance() *Logger {
	once.Do(func() {
		var err error
		instance = &Logger{}
		instance.file, err = os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
	})
	return instance
}

// Log записывает сообщение в лог-файл
func (l *Logger) Log(message string) {
	l.Lock()
	defer l.Unlock()
	log.SetOutput(l.file)
	log.Println(message)
}

//Task two.

type Config struct {
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}
	Port  int
	Debug bool
}

var config *Config

const configPath = "config.json"

func GetInstance1() *Config {
	once.Do(func() {
		config = &Config{}
		configFile, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("error opening config file: %v", err)
		}
		defer configFile.Close()

		jsonParser := json.NewDecoder(configFile)
		if err := jsonParser.Decode(config); err != nil {
			log.Fatalf("error decoding config: %v", err)
		}
	})
	return config
}
