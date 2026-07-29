package main

import (
	"fmt"
	"time"
)

type Personagem struct {
	Nome   string
	Raca   string
	Classe string
	Nivel  int
	Forca  int
	Defesa int
	Vida   int
}

type Jogador struct {
	Personagem
	TempoGameplay int64
	UltimoLogin   time.Time
	Online        bool
}

func (j *Jogador) CalcForcaAtaque() int {
	return j.Forca + j.Nivel + 10
}

func (j *Jogador) CalcForcaDefesa() int {
	return j.Defesa + j.Nivel
}

type Criatura struct {
	Personagem
	RespawTempo time.Duration
}

func (c *Criatura) CalcForcaAtaque() int {
	return c.Forca
}

func (c *Criatura) CalcForcaDefesa() int {
	return c.Defesa
}

type NPC struct {
	Personagem
	Funcao string
}

// Atacante é quem implementa o método CalcForcaAtaque()
type Atacante interface {
	CalcForcaAtaque() int
}

// Defensor é quem implementa o método CalcForcaDefesa()
type Defensor interface {
	CalcForcaDefesa() int
}

type Lutador interface {
	CalcForcaAtaque() int
	CalcForcaDefesa() int
}

func CalculoDano(a Atacante, d Defensor) int {
	return a.CalcForcaAtaque() - d.CalcForcaDefesa()
}

func RinhaDeGo(p1, p2 Lutador){
	fmt.Println(CalculoDano(p1, p2))
	fmt.Println(CalculoDano(p2, p1))
}

func main() {
	mi := Jogador{
		Personagem: Personagem{
			Nome:   "Morfina",
			Raca:   "Alado",
			Classe: "Sacerdote",
			Nivel:  1,
			Forca:  2,
			Defesa: 0,
			Vida:   10,
		},
		TempoGameplay: 60,
		UltimoLogin:   time.Date(2026, 01, 25, 12, 0, 0, 0, time.UTC),
		Online:        true,
	}

	orc := Criatura{
		Personagem: Personagem{
			Nome:   "Orc 1",
			Raca:   "Orc",
			Classe: "Guerreiro",
			Nivel:  2,
			Forca:  3,
			Defesa: 1,
			Vida:   5,
		},
		RespawTempo: 10,
	}

	jormungandr := NPC{
		Personagem: Personagem{
			Nome:   "Jormungandr",
			Raca:   "Jötunn",
			Classe: "Monstro Mitológico",
			Nivel:  1,
			Forca:  2,
			Defesa: 0,
			Vida:   10,
		},
		Funcao: "Historia Antiga",
	}

	fmt.Println("Personagem 1", mi)
	fmt.Println("Criatura 1", orc)
	fmt.Println("NPC 1", jormungandr) // Não pode utilizar CalculoDano()

	// Objetos diferentes com mesmo método implementado
	fmt.Println("Forca de Ataque Mi", mi.CalcForcaAtaque())
	fmt.Println("Forca de Ataque Orc", orc.CalcForcaAtaque())

	// Utilizando Interfaces
	fmt.Println("Dano Mi", CalculoDano(&mi, &orc))
	fmt.Println("Dano Orc", CalculoDano(&orc, &mi))

	fmt.Println("Iniciando Rinha de Go")
	RinhaDeGo(&mi, &orc)

}
