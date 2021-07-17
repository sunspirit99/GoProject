package main

import ("fmt")
import ("strings")
// import ("time")
import ("math/rand")

type doggy []dog

type dog struct{
	name string
	weight int
}

type cat struct{
	name string
	weight int
}

type DogSay interface{
	say(n int) string

	growup(n int) 
	sell(doggy []dog, n int)
}


type Animals interface{

	DogSay
}

// type cat struct{
// 	name string
// 	weight float64
// }

// type []cat Cats{}

// type []dog Dogs{}

func (d *dog) say (n int) string {
	
	s := []string {}
	for i := 0; i < n; i++ {
		s = append(s, "go")
	}
	return strings.Join(s, " ")

}

func (d *dog) growup (n int) {
	d.weight = d.weight + n
}

func (d *dog) sell(doggy []dog, i int) {

	doggy = append(doggy[:i], doggy[i+1:]...)
}



func genDog (n int) []zdog {
	var d = []dog {}
	var tinyDog = []string {"Corgy", "chihuahua", "poodle", }
	var mediumDog = []string {"Alaska", "husky", "PhuQuoc"}
	var giantDog = []string { "Bull", "Golden", "Zephyr"}
	for i:= 0 ; i < n; i++ {
		rw := rand.Intn(60 - 10) + 10 
		if rw < 60 && rw > 40 {
			j := rand.Intn(len(giantDog))
			dg := dog{ giantDog[j], rw }
			d = append(d, dg)
		} else if rw <= 40 && rw >= 20 {
			j := rand.Intn(len(mediumDog))
			dg := dog{ mediumDog[j], rw }
			d = append(d, dg)
		} else{
			j := rand.Intn(len(tinyDog))
			dg := dog{ tinyDog[j], rw }
			d = append(d, dg)
		}
		
	}
	return d
}

func (d dog) String() string{
	return fmt.Sprintf("[Name : %v], [Weight : %v kg]" , d.name, d.weight)
}

func FattestDog(doggy []dog) int{
	a := 0
	max := doggy[0].weight
	for i,_ := range doggy {
		if(doggy[i].weight > max){
			max = doggy[i].weight
		}
	}

	for i,_ := range doggy {
		if (doggy[i].weight == max) {
			a = i
			break
		}
	}

	return a

}

func main(){
	var a DogSay
	dog := genDog(10)

	// s = d

	for _,d := range dog {
		fmt.Println(d)
	}
	fmt.Println(len(dog))

	fmt.Println("One year later :")

	for _,d := range dog {
		a = &d
		a.growup(rand.Intn(10 - 2) + 2)
		fmt.Println(d)
		// time.Sleep(500 * time.Millisecond)
		
	}
	id := FattestDog(dog)
	
	dog[id].sell(dog, id)
	
	fmt.Println("Sold the fattest dog !!!")
	fmt.Println(len(dog))

	for _,d := range dog {
		fmt.Println(d)
	}

	fmt.Println(dog[FattestDog(dog)].weight)


	// fmt.Println(a.say(5))
	// fmt.Println(a.kill())
	// fmt.Println(d.name)
	// fmt.Println(s.say())
}