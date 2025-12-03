package repository



var ListaDichos []string = []string{"como dice el dicho"}


// Esto crea un slice vacío de enteros

func AgregarDicho (dicho string) {
	ListaDichos = append(ListaDichos,dicho)
}


func ObtenerDicho (indice int) string{

	dicho := ListaDichos[indice]

	return  dicho
}


func ListarDichos ()string{

	cantidad := len(ListaDichos)
	var dichos string;

	for i := 0; i < cantidad; i++ {
		dichos+= ListaDichos[i]+"\n"
	}
	return dichos
}


func TamanioLista () int  {
	return  len(ListaDichos)
}