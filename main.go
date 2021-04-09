package main


import "gonum.org/v1/gonum/mat"

import "fmt"


func main() {
  fmt.Println("Introduzca n y m")
	var n int
	var m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)

	//creamos la matriz
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}

	//la llenamos
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Println(i, ":", j)
			var inp int
			fmt.Scanln(&inp)
			a[i][j] = inp
		}
	}
	fmt.Println("matriz original")
	fmt.Println(a)

	//rango
	//se va a tomar que todos los vectores son linealmente independiente
	//es decir que siempre se va a poder realizar el ejercicio

	//transpuesta

	at := make([][]int, m)
	for i := 0; i < m; i++ {
		at[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			at[j][i] = a[i][j]
		}
	}

	fmt.Println("Matriz transpuesta")
	fmt.Println(at)

	//multiplicar la trans por la original

	am := make([][]int, n)
	for i := 0; i < n; i++ {
		am[i] = make([]int, n)
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(at[0]); j++ {
			for k := 0; k < len(a[0]); k++ {
				am[i][j] += a[i][k] * at[k][j]
			}
		}
	}

	fmt.Println("multiplicacion de las matrices")
	fmt.Println(am)

  //sacamos la inversa de a*at
  //aqui usamos una libreria llamada gonum
  //lo que hacemos es castear la matriz a un tipo newDense
  var cont int = 0
  v := make([]float64,len(am)*len(am))
  for i := 0; i < len(am); i++ {
		for j := 0; j < len(am); j++ {
      v[cont] = float64(am[i][j])
      cont += 1
	  }
	}

  atd := mat.NewDense(len(am),len(am),v)

  var inv mat.Dense
  inversa := inv.Inverse(atd)

  if inversa != nil {
    fmt.Println("no es posible")
  }
  
  fmt.Println("inversa de la A*AT ")
  fa := mat.Formatted(&inv, mat.Prefix("       "), mat.Squeeze())
	fmt.Printf("inv = %.2g\n\n", fa)

  //multiplicar inversa por transpuesta
  //at to newDense
  
  var cont2 int = 0
  v1 := make([]float64,len(at)*len(at[0]))
  for i := 0; i < len(at); i++ {
		for j := 0; j < len(at[0]); j++ {
      v1[cont2] = float64(at[i][j])
      cont2 += 1
	  }
	}
  atnd := mat.NewDense(len(at),len(at[0]),v1)

  // matPrint(atnd)

  // multiplicamos inv*at
  fmt.Println("PseudoInversa de la Matrix A: ")
  var pseudoinversa mat.Dense
	pseudoinversa.Mul(atnd, &inv)
	fi := mat.Formatted(&pseudoinversa, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("psinv = %v\n\n", fi)
}

func matPrint(X mat.Matrix) {
 fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
 fmt.Printf("%v\n", fa)
}